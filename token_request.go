package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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
	ClientID  string `json:"client_id"`
}

// ImportToken pulls in bots from the provided db
func (m *Manager) ImportToken() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, name, nickname, color, activity, network, contract, decorator, decimals, source, frequency FROM tokens")
	if err != nil {
		logger.Warningf("Unable to query tokens in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var clientID, token, name, activity, network, contract, decorator, source string
		var nickname, color bool
		var decimals, frequency int
		err = rows.Scan(&clientID, &token, &name, &nickname, &color, &activity, &network, &contract, &decorator, &decimals, &source, &frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		// activate bot
		t := NewToken(clientID, network, contract, token, name, nickname, frequency, decimals, activity, color, decorator, source, lastUpdate)
		m.addToken(t, false)
		logger.Infof("Loaded token from db: %s-%s", network, contract)
	}
	rows.Close()
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

	token := NewToken(tokenReq.ClientID, tokenReq.Network, tokenReq.Contract, tokenReq.Token, tokenReq.Name, tokenReq.Nickname, tokenReq.Frequency, tokenReq.Decimals, tokenReq.Activity, tokenReq.Color, tokenReq.Decorator, tokenReq.Source, lastUpdate)
	m.addToken(token, true)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added token: %s-%s\n", token.Network, token.Contract)
}

func (m *Manager) addToken(token *Token, update bool) {
	tokenCount.Inc()
	id := fmt.Sprintf("%s-%s", token.Network, token.Contract)
	m.WatchingToken[id] = token

	var noDB *sql.DB
	if (m.DB == noDB) || !update {
		return
	}

	// query
	stmt, err := m.DB.Prepare("SELECT id FROM tokens WHERE network = ? AND contract = ? LIMIT 1")
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
		stmt, err := m.DB.Prepare("update tokens set clientId = ?, token = ?, name = ?, nickname = ?, color = ?, activity = ?, network = ?, contract = ?, decorator = ?, decimals = ?, source = ?, frequency = ? WHERE id = ?")
		if err != nil {
			logger.Warningf("Unable to update token in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(token.ClientID, token.token, token.Name, token.Nickname, token.Color, token.Activity, token.Network, token.Contract, token.Decorator, token.Decimals, token.Source, token.Frequency, existingId)
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
		stmt, err := m.DB.Prepare("INSERT INTO tokens(clientId, token, name, nickname, color, activity, network, contract, decorator, decimals, source, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			logger.Warningf("Unable to store token in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(token.ClientID, token.token, token.Name, token.Nickname, token.Color, token.Activity, token.Network, token.Contract, token.Decorator, token.Decimals, token.Source, token.Frequency)
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

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		stmt, err := m.DB.Prepare("DELETE FROM tokens WHERE network = ? AND contract = ?")
		if err != nil {
			logger.Warningf("Unable to query token in db %s: %s", id, err)
			return
		}

		_, err = stmt.Exec(m.WatchingToken[id].Network, m.WatchingToken[id].Contract)
		if err != nil {
			logger.Warningf("Unable to query token in db %s: %s", id, err)
			return
		}
	}

	// remove from cache
	delete(m.WatchingToken, id)

	logger.Infof("Deleted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartToken stops and starts a Token
func (m *Manager) RestartToken(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a token")

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

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingToken[id].Frequency) * 2 * time.Second)

	// start the ticker again
	m.WatchingToken[id].Start()

	logger.Infof("Restarted ticker %s", id)
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
