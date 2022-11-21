package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// ImportValueLocked pulls in bots from the provided db
func (m *Manager) ImportValueLocked() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, ticker, name, nickname, activity, decorator, decimals, currency, currencySymbol, source, frequency FROM valuelocks")
	if err != nil {
		logger.Warningf("Unable to query valuelocks in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedValueLocked ValueLocked

		err = rows.Scan(&importedValueLocked.ClientID, &importedValueLocked.Token, &importedValueLocked.Ticker, &importedValueLocked.Name, &importedValueLocked.Nickname, &importedValueLocked.Activity, &importedValueLocked.Decorator, &importedValueLocked.Decimals, &importedValueLocked.Currency, &importedValueLocked.CurrencySymbol, &importedValueLocked.Source, &importedValueLocked.Frequency)
		if err != nil {
			logger.Errorf("Unable to load valuelocks from db: %s", err)
			continue
		}

		// activate bot
		go importedValueLocked.watchValueLocked()
		m.WatchValueLocked(&importedValueLocked)
		logger.Infof("Loaded valuelocked from db: %s", importedValueLocked.label())
	}
	rows.Close()

	// check all entries have id
	for _, valuelocked := range m.WatchingValueLocked {
		if valuelocked.ClientID == "" {
			id, err := getIDToken(valuelocked.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE valuelocks SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, valuelocked.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", valuelocked.label())
				valuelocked.ClientID = id
			}
		}
	}
}

// AddValueLocked adds a new ValueLocked or crypto to the list of what to watch
func (m *Manager) AddValueLocked(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a ValueLocked")

	// read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var valueLockedReq ValueLocked
	if err := json.Unmarshal(body, &valueLockedReq); err != nil {
		logger.Errorf("Unmarshalling: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if valueLockedReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make sure token is valid
	if valueLockedReq.ClientID == "" {
		id, err := getIDToken(valueLockedReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		valueLockedReq.ClientID = id
	}

	// ensure frequency is set
	if valueLockedReq.Frequency <= 0 {
		valueLockedReq.Frequency = 60
	}

	// ensure currency is set
	if valueLockedReq.Currency == "" {
		valueLockedReq.Currency = "usd"
	}

	// ensure name is set
	if valueLockedReq.Name == "" {
		logger.Error("Name required for crypto")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure currency is set
	if valueLockedReq.CurrencySymbol == "" {
		valueLockedReq.CurrencySymbol = "$"
	}

	// check if already existing
	if _, ok := m.WatchingValueLocked[valueLockedReq.label()]; ok {
		logger.Error("ValueLocked already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	go valueLockedReq.watchValueLocked()
	m.WatchValueLocked(&valueLockedReq)

	if *db != "" {
		m.StoreValueLocked(&valueLockedReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(valueLockedReq)
	if err != nil {
		logger.Errorf("Unable to encode valuelocks: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added valuelocked: %s\n", valueLockedReq.Name)
}

func (m *Manager) WatchValueLocked(valuelocked *ValueLocked) {
	valuelockedCount.Inc()
	id := valuelocked.label()
	m.WatchingValueLocked[id] = valuelocked
}

// StoreTicker puts a valuelocked into the db
func (m *Manager) StoreValueLocked(valuelocked *ValueLocked) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO valuelocks(clientId, token, ticker, name, nickname, activity, decorator, decimals, currency, currencySymbol, source, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store valuelocked in db %s: %s", valuelocked.label(), err)
		return
	}

	res, err := stmt.Exec(valuelocked.ClientID, valuelocked.Token, valuelocked.Ticker, valuelocked.Name, valuelocked.Nickname, valuelocked.Activity, valuelocked.Decorator, valuelocked.Decimals, valuelocked.Currency, valuelocked.CurrencySymbol, valuelocked.Source, valuelocked.Frequency)
	if err != nil {
		logger.Warningf("Unable to store valuelocked in db %s: %s", valuelocked.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store valuelocked in db %s: %s", valuelocked.label(), err)
		return
	}
}

// DeleteValueLocked addds a new ValueLocked or crypto to the list of what to watch
func (m *Manager) DeleteValueLocked(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a valuelocked")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingValueLocked[id]; !ok {
		logger.Errorf("No valuelocked found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingValueLocked[id].Shutdown()
	valuelockedCount.Dec()

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		stmt, err := m.DB.Prepare("DELETE FROM valuelocks WHERE name = ?")
		if err != nil {
			logger.Warningf("Unable to query valuelocked in db %s: %s", id, err)
			return
		}

		_, err = stmt.Exec(m.WatchingValueLocked[id].Name)
		if err != nil {
			logger.Warningf("Unable to query valuelocked in db %s: %s", id, err)
			return
		}
	}

	// remove from cache
	delete(m.WatchingValueLocked, id)

	logger.Infof("Deleted valuelocked %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartValueLocked stops and starts a valuelocked
func (m *Manager) RestartValueLocked(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a valuelocked")

	vars := mux.Vars(r)
	id := strings.ToUpper(vars["id"])

	if _, ok := m.WatchingValueLocked[id]; !ok {
		logger.Errorf("No valuelocked found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingValueLocked[id].Shutdown()

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingValueLocked[id].Frequency) * 2 * time.Second)

	// start the valuelocked again
	go m.WatchingValueLocked[id].watchValueLocked()

	logger.Infof("Restarted valuelocked %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetValueLockeds returns a list of what the manager is watching
func (m *Manager) GetValueLockeds(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingValueLocked); err != nil {
		logger.Errorf("Serving request: %s", err)
	}
}
