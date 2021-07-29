package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// MaticRequest represents the json coming in from the request
type MaticRequest struct {
	Contract  string `json:"contract"`
	Token     string `json:"discord_bot_token"`
	Name      string `json:"name"`
	Nickname  bool   `json:"set_nickname"`
	Frequency int    `json:"frequency" default:"60"`
	Currency  string `json:"currency" default:"0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"`
	Decimals  int    `json:"decimals"`
}

// AddMatic adds a new Matic or crypto to the list of what to watch
func (m *Manager) AddMatic(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a ticker")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading body: %v", err)
		return
	}

	// unmarshal into struct
	var tokenReq MaticRequest
	if err := json.Unmarshal(body, &tokenReq); err != nil {
		logger.Errorf("Error unmarshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error unmarshalling: %v", err)
		return
	}

	// ensure token is set
	if tokenReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: token required")
		return
	}

	// ensure currency is set
	if tokenReq.Currency == "" {
		tokenReq.Currency = "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"
	}

	// ensure freq is set
	if tokenReq.Frequency == 0 {
		tokenReq.Frequency = 60
	}

	// ensure name is set
	if tokenReq.Name == "" {
		logger.Error("Name required for token")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Name required")
		return
	}

	// check if already existing
	if _, ok := m.WatchingMatic[strings.ToUpper(tokenReq.Contract)]; ok {
		logger.Error("Error: ticker already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	token := NewMatic(tokenReq.Contract, tokenReq.Token, tokenReq.Name, tokenReq.Nickname, tokenReq.Frequency, tokenReq.Currency, tokenReq.Decimals)
	m.addMatic(token)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}

func (m *Manager) addMatic(token *Matic) {
	maticCount.Inc()
	m.WatchingMatic[token.Contract] = token
}

// DeleteMatic addds a new token or crypto to the list of what to watch
func (m *Manager) DeleteMatic(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a ticker")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingMatic[id]; !ok {
		logger.Error("Error: no ticker found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: ticker not found")
		return
	}
	// send shutdown sign
	m.WatchingMatic[id].Shutdown()
	maticCount.Dec()

	// remove from cache
	delete(m.WatchingMatic, id)

	logger.Infof("Deleted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetMatic returns a list of what the manager is watching
func (m *Manager) GetMatic(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingMatic); err != nil {
		logger.Errorf("Error serving request: %v", err)
		fmt.Fprintf(w, "Error: %v", err)
	}
}
