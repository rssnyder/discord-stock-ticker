package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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
		var importedToken Token

		err = rows.Scan(&importedToken.ClientID, &importedToken.Token, &importedToken.Name, &importedToken.Nickname, &importedToken.Color, &importedToken.Activity, &importedToken.Network, &importedToken.Contract, &importedToken.Decorator, &importedToken.Decimals, &importedToken.Source, &importedToken.Frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		// activate bot
		go importedToken.watchTokenPrice()
		m.WatchToken(&importedToken)
		logger.Infof("Loaded token from db: %s", importedToken.label())
	}
	rows.Close()

	// check all entries have id
	for _, token := range m.WatchingToken {
		if token.ClientID == "" {
			id, err := getIDToken(token.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE tokens SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, token.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", token.label())
				token.ClientID = id
			}
		}
	}
}

// AddToken adds a new Token or crypto to the list of what to watch
func (m *Manager) AddToken(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a token")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading body: %v", err)
		return
	}

	// unmarshal into struct
	var tokenReq Token
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

	// make sure token is valid
	if tokenReq.ClientID == "" {
		id, err := getIDToken(tokenReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenReq.ClientID = id
	}

	// ensure frequency is set
	if tokenReq.Frequency <= 0 {
		tokenReq.Frequency = 60
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
	if _, ok := m.WatchingToken[tokenReq.label()]; ok {
		logger.Error("Error: token already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	go tokenReq.watchTokenPrice()
	m.WatchToken(&tokenReq)

	if *db != "" {
		m.StoreToken(&tokenReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tokenReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added token: %s-%s\n", tokenReq.Network, tokenReq.Contract)
}

func (m *Manager) WatchToken(token *Token) {
	tokenCount.Inc()
	id := token.label()
	m.WatchingToken[id] = token
}

// StoreTicker puts a token into the db
func (m *Manager) StoreToken(token *Token) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO tokens(clientId, token, name, nickname, color, activity, network, contract, decorator, decimals, source, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store token in db %s: %s", token.label(), err)
		return
	}

	res, err := stmt.Exec(token.ClientID, token.Token, token.Name, token.Nickname, token.Color, token.Activity, token.Network, token.Contract, token.Decorator, token.Decimals, token.Source, token.Frequency)
	if err != nil {
		logger.Warningf("Unable to store token in db %s: %s", token.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store token in db %s: %s", token.label(), err)
		return
	}
}

// DeleteToken addds a new token or crypto to the list of what to watch
func (m *Manager) DeleteToken(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a token")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingToken[id]; !ok {
		logger.Error("Error: no token found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: token not found")
		return
	}

	// send shutdown sign
	m.WatchingToken[id].Close <- 1
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

	logger.Infof("Deleted token %s", id)
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
		logger.Error("Error: no token found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: token not found")
		return
	}

	// send shutdown sign
	m.WatchingToken[id].Close <- 1

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingToken[id].Frequency) * 2 * time.Second)

	// start the token again
	go m.WatchingToken[id].watchTokenPrice()

	logger.Infof("Restarted token %s", id)
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
