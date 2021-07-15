package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	m.addHolders(holders)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(holders)
	if err != nil {
		logger.Errorf("Unable to encode holders: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addHolders(holders *Holders) {
	holdersCount.Inc()
	m.WatchingHolders[fmt.Sprintf("%s-%s", holders.Network, holders.Address)] = holders
}

// DeleteHolders addds a new holders or crypto to the list of what to watch
func (m *Manager) DeleteHolders(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a holders")

	vars := mux.Vars(r)
	id := strings.ToUpper(vars["id"])

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
