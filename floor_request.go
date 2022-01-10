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
	m.Lock()
	defer m.Unlock()

	// query
	rows, err := m.DB.Query("SELECT id, clientID, token, nickname, marketplace, name, frequency FROM floors")
	if err != nil {
		logger.Warningf("Unable to query tokens in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedFloor Floor
		var importedID int

		err = rows.Scan(&importedID, &importedFloor.ClientID, &importedFloor.Token, &importedFloor.Nickname, &importedFloor.Marketplace, &importedFloor.Name, &importedFloor.Frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		go importedFloor.watchFloorPrice()
		m.WatchFloor(&importedFloor)
		logger.Infof("Loaded floor from db: %s", importedFloor.label())
	}
	rows.Close()

	// check all entries have id
	for _, floor := range m.WatchingFloor {
		if floor.ClientID == "" {
			id, err := getIDToken(floor.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE floors SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, floor.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", floor.label())
				floor.ClientID = id
			}
		}
	}
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

	// make sure token is valid
	if floorReq.ClientID == "" {
		id, err := getIDToken(floorReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		floorReq.ClientID = id
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
	m.WatchFloor(&floorReq)

	if *db != "" {
		m.StoreFloor(&floorReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(floorReq)
	if err != nil {
		logger.Errorf("Unable to encode floor: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added floor: %s\n", floorReq.Marketplace)
}

func (m *Manager) WatchFloor(floor *Floor) {
	floorCount.Inc()
	id := floor.label()
	m.WatchingFloor[id] = floor
}

// StoreTicker puts a floor into the db
func (m *Manager) StoreFloor(floor *Floor) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO floors(clientId, token, nickname, marketplace, name, frequency) values(?,?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store floor in db %s: %s", floor.label(), err)
		return
	}

	res, err := stmt.Exec(floor.ClientID, floor.Token, floor.Nickname, floor.Marketplace, floor.Name, floor.Frequency)
	if err != nil {
		logger.Warningf("Unable to store floor in db %s: %s", floor.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store floor in db %s: %s", floor.label(), err)
		return
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

	// start the floor again
	go m.WatchingFloor[id].watchFloorPrice()

	logger.Infof("Restarted floor %s", id)
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
