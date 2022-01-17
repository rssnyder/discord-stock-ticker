package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// ImportCirculating pulls in bots from the provided db
func (m *Manager) ImportCirculating() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, ticker, name, nickname, activity, decimals, currencySymbol, frequency FROM circulatings")
	if err != nil {
		logger.Warningf("Unable to query circulatings in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedCirculating Circulating

		err = rows.Scan(&importedCirculating.ClientID, &importedCirculating.Token, &importedCirculating.Ticker, &importedCirculating.Name, &importedCirculating.Nickname, &importedCirculating.Activity, &importedCirculating.Decimals, &importedCirculating.CurrencySymbol, &importedCirculating.Frequency)
		if err != nil {
			logger.Errorf("Unable to load circulatings from db: %s", err)
			continue
		}

		// activate bot
		go importedCirculating.watchCirculating()
		m.WatchCirculating(&importedCirculating)
		logger.Infof("Loaded circulating from db: %s", importedCirculating.label())
	}
	rows.Close()

	// check all entries have id
	for _, circulating := range m.WatchingCirculating {
		if circulating.ClientID == "" {
			id, err := getIDToken(circulating.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE circulatings SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, circulating.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", circulating.label())
				circulating.ClientID = id
			}
		}
	}
}

// AddCirculating adds a new Circulating or crypto to the list of what to watch
func (m *Manager) AddCirculating(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a Circulating")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var circulatingReq Circulating
	if err := json.Unmarshal(body, &circulatingReq); err != nil {
		logger.Errorf("Unmarshalling: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if circulatingReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make sure token is valid
	if circulatingReq.ClientID == "" {
		id, err := getIDToken(circulatingReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		circulatingReq.ClientID = id
	}

	// ensure frequency is set
	if circulatingReq.Frequency <= 0 {
		circulatingReq.Frequency = 60
	}

	// ensure name is set
	if circulatingReq.Name == "" {
		logger.Error("Name required for crypto")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if already existing
	if _, ok := m.WatchingCirculating[circulatingReq.label()]; ok {
		logger.Error("Circulating already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	go circulatingReq.watchCirculating()
	m.WatchCirculating(&circulatingReq)

	if *db != "" {
		m.StoreCirculating(&circulatingReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(circulatingReq)
	if err != nil {
		logger.Errorf("Unable to encode circulatings: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added circulating: %s\n", circulatingReq.Name)
}

func (m *Manager) WatchCirculating(circulating *Circulating) {
	circulatingCount.Inc()
	id := circulating.label()
	m.WatchingCirculating[id] = circulating
}

// StoreTicker puts a circulating into the db
func (m *Manager) StoreCirculating(circulating *Circulating) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO circulatings(clientId, token, ticker, name, nickname, activity, decimals, currencySymbol, frequency) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store circulating in db %s: %s", circulating.label(), err)
		return
	}

	res, err := stmt.Exec(circulating.ClientID, circulating.Token, circulating.Ticker, circulating.Name, circulating.Nickname, circulating.Activity, circulating.Decimals, circulating.CurrencySymbol, circulating.Frequency)
	if err != nil {
		logger.Warningf("Unable to store circulating in db %s: %s", circulating.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store circulating in db %s: %s", circulating.label(), err)
		return
	}
}

// DeleteCirculating addds a new Circulating or crypto to the list of what to watch
func (m *Manager) DeleteCirculating(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a circulating")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingCirculating[id]; !ok {
		logger.Errorf("No circulating found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingCirculating[id].Close <- 1
	circulatingCount.Dec()

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		stmt, err := m.DB.Prepare("DELETE FROM circulatings WHERE name = ?")
		if err != nil {
			logger.Warningf("Unable to query circulating in db %s: %s", id, err)
			return
		}

		_, err = stmt.Exec(m.WatchingCirculating[id].Name)
		if err != nil {
			logger.Warningf("Unable to query circulating in db %s: %s", id, err)
			return
		}
	}

	// remove from cache
	delete(m.WatchingCirculating, id)

	logger.Infof("Deleted circulating %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartCirculating stops and starts a circulating
func (m *Manager) RestartCirculating(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a circulating")

	vars := mux.Vars(r)
	id := strings.ToUpper(vars["id"])

	if _, ok := m.WatchingCirculating[id]; !ok {
		logger.Errorf("No circulating found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingCirculating[id].Close <- 1

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingCirculating[id].Frequency) * 2 * time.Second)

	// start the circulating again
	go m.WatchingCirculating[id].watchCirculating()

	logger.Infof("Restarted circulating %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetCirculatings returns a list of what the manager is watching
func (m *Manager) GetCirculatings(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingCirculating); err != nil {
		logger.Errorf("Serving request: %s", err)
	}
}
