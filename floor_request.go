package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ImportFloor pulls in bots from the provided db
func (m *Manager) ImportFloor() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, nickname, marketplace, name, frequency FROM floors")
	if err != nil {
		logger.Warningf("Unable to query tokens in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedFloor Floor

		err = rows.Scan(&importedFloor.ClientID, &importedFloor.Token, &importedFloor.Nickname, &importedFloor.Marketplace, &importedFloor.Name, &importedFloor.Frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		go importedFloor.watchFloorPrice()
		m.StoreFloor(&importedFloor, false)
		logger.Infof("Loaded floor from db: %s", importedFloor.Marketplace)
	}
	rows.Close()
}

// AddTicker adds a new Ticker or crypto to the list of what to watch
func (m *Manager) AddFloor(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a floor")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var floorReq Floor
	if err := json.Unmarshal(body, &floorReq); err != nil {
		logger.Errorf("Unmarshalling: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if floorReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure frequency is set
	if floorReq.Frequency <= 0 {
		floorReq.Frequency = 60
	}

	// ensure marketplace is set
	if floorReq.Marketplace == "" {
		logger.Error("Marketplace required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure name is set
	if floorReq.Name == "" {
		logger.Error("Name required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if already existing
	if _, ok := m.WatchingFloor[floorReq.label()]; ok {
		logger.Error("Marketplace already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	go floorReq.watchFloorPrice()
	m.StoreFloor(&floorReq, true)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(floorReq)
	if err != nil {
		logger.Errorf("Unable to encode floor: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added floor: %s\n", floorReq.Marketplace)
}

func (m *Manager) StoreFloor(floor *Floor, update bool) {
	floorCount.Inc()
	id := floor.label()
	m.WatchingFloor[id] = floor

	var noDB *sql.DB
	if (m.DB == noDB) || !update {
		return
	}

	// query
	stmt, err := m.DB.Prepare("SELECT id FROM floors WHERE marketplace = ? AND name = ? LIMIT 1")
	if err != nil {
		logger.Warningf("Unable to query floor in db %s: %s", id, err)
		return
	}

	rows, err := stmt.Query(floor.Marketplace, floor.Name)
	if err != nil {
		logger.Warningf("Unable to query floor in db %s: %s", id, err)
		return
	}

	var existingId int

	for rows.Next() {
		err = rows.Scan(&existingId)
		if err != nil {
			logger.Warningf("Unable to query floor in db %s: %s", id, err)
			return
		}
	}
	rows.Close()

	if existingId != 0 {

		// update entry in db
		stmt, err := m.DB.Prepare("update floors set clientId = ?, token = ?, nickname = ?, marketplace = ?, name = ?, frequency = ? WHERE id = ?")
		if err != nil {
			logger.Warningf("Unable to update floor in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(floor.ClientID, floor.Token, floor.Nickname, floor.Marketplace, floor.Name, floor.Frequency, existingId)
		if err != nil {
			logger.Warningf("Unable to update floor in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to update floor in db %s: %s", id, err)
			return
		}

		logger.Infof("Updated floor in db %s", id)
	} else {

		// store new entry in db
		stmt, err := m.DB.Prepare("INSERT INTO floors(clientId, token, nickname, marketplace, name, frequency) values(?,?,?,?,?,?)")
		if err != nil {
			logger.Warningf("Unable to store floor in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(floor.ClientID, floor.Token, floor.Nickname, floor.Marketplace, floor.Name, floor.Frequency)
		if err != nil {
			logger.Warningf("Unable to store floor in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to store floor in db %s: %s", id, err)
			return
		}
	}
}

// DeleteFloor addds a new floor or crypto to the list of what to watch
func (m *Manager) DeleteFloor(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a floor")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingFloor[id]; !ok {
		logger.Errorf("No floor found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// send shutdown sign
	m.WatchingFloor[id].Close <- 1
	floorCount.Dec()

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		stmt, err := m.DB.Prepare("DELETE FROM floors WHERE marketplace = ?")
		if err != nil {
			logger.Warningf("Unable to query holder in db %s: %s", id, err)
			return
		}

		_, err = stmt.Exec(m.WatchingFloor[id].Marketplace)
		if err != nil {
			logger.Warningf("Unable to query holder in db %s: %s", id, err)
			return
		}
	}

	// remove from cache
	delete(m.WatchingFloor, id)

	logger.Infof("Deleted floor %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartFloor stops and starts a floor
func (m *Manager) RestartFloor(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a floor")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingFloor[id]; !ok {
		logger.Errorf("No floor found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// send shutdown sign
	m.WatchingFloor[id].Close <- 1

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingFloor[id].Frequency) * 2 * time.Second)

	// start the ticker again
	go m.WatchingFloor[id].watchFloorPrice()

	logger.Infof("Restarted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetFloor returns a list of what the manager is watching
func (m *Manager) GetFloor(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingFloor); err != nil {
		logger.Errorf("Serving request: %s", err)
	}
}
