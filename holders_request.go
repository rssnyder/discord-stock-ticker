package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// HoldersRequest represents the json coming in from the request
type HoldersRequest struct {
	Network   string `json:"network"`
	Address   string `json:"address"`
	Activity  string `json:"activity"`
	Token     string `json:"discord_bot_token"`
	Nickname  bool   `json:"set_nickname"`
	Frequency int    `json:"frequency" default:"60"`
}

// ImportHolder pulls in bots from the provided db
func (m *Manager) ImportHolder() {

	// query
	rows, err := m.DB.Query("SELECT token, nickname, activity, network, address, frequency FROM holders")
	if err != nil {
		logger.Warningf("Unable to query tokens in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var token, activity, network, address string
		var nickname bool
		var frequency int
		err = rows.Scan(&token, &nickname, &activity, &network, &address, &frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		h := NewHolders(network, address, activity, token, nickname, frequency)
		m.addHolders(h, false)
		logger.Infof("Loaded holder from db: %s-%s", network, address)
	}
	rows.Close()
}

// AddTicker adds a new Ticker or crypto to the list of what to watch
func (m *Manager) AddHolders(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a holders")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var holdersReq HoldersRequest
	if err := json.Unmarshal(body, &holdersReq); err != nil {
		logger.Errorf("Unmarshalling: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if holdersReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure network is set
	if holdersReq.Network == "" {
		logger.Error("Network required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure address is set
	if holdersReq.Address == "" {
		logger.Error("Address required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if already existing
	if _, ok := m.WatchingHolders[fmt.Sprintf("%s-%s", holdersReq.Network, holdersReq.Address)]; ok {
		logger.Error("Network already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	holders := NewHolders(holdersReq.Network, holdersReq.Address, holdersReq.Activity, holdersReq.Token, holdersReq.Nickname, holdersReq.Frequency)
	m.addHolders(holders, true)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(holders)
	if err != nil {
		logger.Errorf("Unable to encode holders: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addHolders(holders *Holders, update bool) {
	holdersCount.Inc()
	id := fmt.Sprintf("%s-%s", holders.Network, holders.Address)
	m.WatchingHolders[id] = holders

	var noDB *sql.DB
	if (m.DB == noDB) || !update {
		return
	}

	// query
	stmt, err := m.DB.Prepare("SELECT id FROM holders WHERE network = ? AND address = ? LIMIT 1")
	if err != nil {
		logger.Warningf("Unable to query holders in db %s: %s", id, err)
		return
	}

	rows, err := stmt.Query(holders.Network, holders.Address)
	if err != nil {
		logger.Warningf("Unable to query holders in db %s: %s", id, err)
		return
	}

	var existingId int

	for rows.Next() {
		err = rows.Scan(&existingId)
		if err != nil {
			logger.Warningf("Unable to query holders in db %s: %s", id, err)
			return
		}
	}
	rows.Close()

	if existingId != 0 {

		// update entry in db
		stmt, err := m.DB.Prepare("update holders set token = ?, nickname = ?, activity = ?, network = ?, address = ?, frequency = ? WHERE id = ?")
		if err != nil {
			logger.Warningf("Unable to update holders in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(holders.token, holders.Nickname, holders.Activity, holders.Network, holders.Address, holders.Frequency, existingId)
		if err != nil {
			logger.Warningf("Unable to update holders in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to update holders in db %s: %s", id, err)
			return
		}

		logger.Infof("Updated holders in db %s", id)
	} else {

		// store new entry in db
		stmt, err := m.DB.Prepare("INSERT INTO holders(token, nickname, activity, network, address, frequency) values(?,?,?,?,?,?)")
		if err != nil {
			logger.Warningf("Unable to store holders in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(holders.token, holders.Nickname, holders.Activity, holders.Network, holders.Address, holders.Frequency)
		if err != nil {
			logger.Warningf("Unable to store holders in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to store holders in db %s: %s", id, err)
			return
		}
	}
}

// DeleteHolders addds a new holders or crypto to the list of what to watch
func (m *Manager) DeleteHolders(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a holders")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingHolders[id]; !ok {
		logger.Errorf("No holders found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingHolders[id].Shutdown()
	holdersCount.Dec()

	// remove from cache
	delete(m.WatchingHolders, id)

	logger.Infof("Deleted holders %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetHolders returns a list of what the manager is watching
func (m *Manager) GetHolders(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingHolders); err != nil {
		logger.Errorf("Serving request: %s", err)
	}
}
