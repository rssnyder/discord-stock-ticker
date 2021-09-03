package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// TokenRequest represents the json coming in from the request
type TokenRequest struct {
	Network   string `json:"network"`
	Contract  string `json:"contract"`
	Token     string `json:"discord_bot_token"`
	Name      string `json:"name"`
	Nickname  bool   `json:"set_nickname"`
	Frequency int    `json:"frequency" default:"60"`
	Color     bool   `json:"set_color"`
	Decorator string `json:"decorator" default:"-"`
	Activity  string `json:"activity"`
	Decimals  int    `json:"decimals"`
	Source    string `json:"source"`
}

// AddToken adds a new Token or crypto to the list of what to watch
func (m *Manager) AddToken(w http.ResponseWriter, r *http.Request) {
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
	var tokenReq TokenRequest
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

	// ensure network is set, default to eth
	if tokenReq.Network == "" {
		tokenReq.Network = "ethereum"
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
	if _, ok := m.WatchingToken[strings.ToUpper(tokenReq.Contract)]; ok {
		logger.Error("Error: ticker already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	token := NewToken(tokenReq.Network, tokenReq.Contract, tokenReq.Token, tokenReq.Name, tokenReq.Nickname, tokenReq.Frequency, tokenReq.Decimals, tokenReq.Activity, tokenReq.Color, tokenReq.Decorator, tokenReq.Source)
	m.addToken(token)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addToken(token *Token) {
	tokenCount.Inc()
	id := fmt.Sprintf("%s-%s", token.Network, token.Contract)
	m.WatchingToken[id] = token

	var noDB *sql.DB
	if m.DB == noDB {
		return
	}

	// query
	stmt, err := m.DB.Prepare("SELECT id FROM tickers WHERE tickerType = 'token' AND network = ? AND contract = ? LIMIT 1")
	if err != nil {
		logger.Warningf("Unable to query token in db %s: %s", id, err)
		return
	}

	rows, err := stmt.Query(token.Network, token.Contract)
	if err != nil {
		logger.Warningf("Unable to query token in db %s: %s", id, err)
		return
	}

	var existingId int

	for rows.Next() {
		err = rows.Scan(&existingId)
		if err != nil {
			logger.Warningf("Unable to query token in db %s: %s", id, err)
			return
		}
	}
	rows.Close()

	if existingId != 0 {

		// update entry in db
		stmt, err := m.DB.Prepare("update tickers set token = ?, name = ?, nickname = ?, color = ?, activity = ?, network = ?, contract = ?, decorator = ?, decimals = ?, source = ?, frequency = ? WHERE id = ?")
		if err != nil {
			logger.Warningf("Unable to update token in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(token.token, token.Name, token.Nickname, token.Color, token.Activity, token.Network, token.Contract, token.Decorator, token.Decimals, token.Source, token.Frequency, existingId)
		if err != nil {
			logger.Warningf("Unable to update token in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to update token in db %s: %s", id, err)
			return
		}

		logger.Infof("Updated token in db %s", id)
	} else {

		// store new entry in db
		stmt, err := m.DB.Prepare("INSERT INTO tickers(tickerType, token, name, nickname, color, activity, network, contract, decorator, decimals, source, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			logger.Warningf("Unable to store token in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec("token", token.token, token.Name, token.Nickname, token.Color, token.Activity, token.Network, token.Contract, token.Decorator, token.Decimals, token.Source, token.Frequency)
		if err != nil {
			logger.Warningf("Unable to store token in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to store token in db %s: %s", id, err)
			return
		}
	}
}

// DeleteToken addds a new token or crypto to the list of what to watch
func (m *Manager) DeleteToken(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a ticker")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingToken[id]; !ok {
		logger.Error("Error: no ticker found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: ticker not found")
		return
	}
	// send shutdown sign
	m.WatchingToken[id].Shutdown()
	tokenCount.Dec()

	// remove from cache
	delete(m.WatchingToken, id)

	logger.Infof("Deleted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetToken returns a list of what the manager is watching
func (m *Manager) GetToken(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingToken); err != nil {
		logger.Errorf("Error serving request: %v", err)
		fmt.Fprintf(w, "Error: %v", err)
	}
}
