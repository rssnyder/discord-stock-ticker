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

// ImportMarketCap pulls in bots from the provided db
func (m *Manager) ImportMarketCap() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, ticker, name, nickname, color, activity, decorator, decimals, currency, currencySymbol, frequency FROM marketcaps")
	if err != nil {
		logger.Warningf("Unable to query marketcaps in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedMarketCap MarketCap

		err = rows.Scan(&importedMarketCap.ClientID, &importedMarketCap.Token, &importedMarketCap.Ticker, &importedMarketCap.Name, &importedMarketCap.Nickname, &importedMarketCap.Color, &importedMarketCap.Activity, &importedMarketCap.Decorator, &importedMarketCap.Decimals, &importedMarketCap.Currency, &importedMarketCap.CurrencySymbol, &importedMarketCap.Frequency)
		if err != nil {
			logger.Errorf("Unable to load marketcaps from db: %s", err)
			continue
		}

		// activate bot
		go importedMarketCap.watchMarketCap()
		m.WatchMarketCap(&importedMarketCap)
		logger.Infof("Loaded marketcap from db: %s", importedMarketCap.label())
	}
	rows.Close()

	// check all entries have id
	for _, marketcap := range m.WatchingMarketCap {
		if marketcap.ClientID == "" {
			id, err := getIDToken(marketcap.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE marketcaps SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, marketcap.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", marketcap.label())
				marketcap.ClientID = id
			}
		}
	}
}

// AddMarketCap adds a new MarketCap or crypto to the list of what to watch
func (m *Manager) AddMarketCap(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a MarketCap")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var marketCapReq MarketCap
	if err := json.Unmarshal(body, &marketCapReq); err != nil {
		logger.Errorf("Unmarshalling: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if marketCapReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make sure token is valid
	if marketCapReq.ClientID == "" {
		id, err := getIDToken(marketCapReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		marketCapReq.ClientID = id
	}

	// ensure frequency is set
	if marketCapReq.Frequency <= 0 {
		marketCapReq.Frequency = 60
	}

	// ensure currency is set
	if marketCapReq.Currency == "" {
		marketCapReq.Currency = "usd"
	}

	// ensure name is set
	if marketCapReq.Name == "" {
		logger.Error("Name required for crypto")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure currency is set
	if marketCapReq.CurrencySymbol == "" {
		marketCapReq.CurrencySymbol = "$"
	}

	// check if already existing
	if _, ok := m.WatchingMarketCap[marketCapReq.label()]; ok {
		logger.Error("MarketCap already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	go marketCapReq.watchMarketCap()
	m.WatchMarketCap(&marketCapReq)

	if *db != "" {
		m.StoreMarketcap(&marketCapReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(marketCapReq)
	if err != nil {
		logger.Errorf("Unable to encode marketcaps: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added marketcap: %s\n", marketCapReq.Name)
}

func (m *Manager) WatchMarketCap(marketcap *MarketCap) {
	marketcapCount.Inc()
	id := marketcap.label()
	m.WatchingMarketCap[id] = marketcap
}

// StoreTicker puts a marketcap into the db
func (m *Manager) StoreMarketcap(marketcap *MarketCap) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO marketcaps(clientId, token, ticker, name, nickname, color, activity, decorator, decimals, currency, currencySymbol, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store marketcap in db %s: %s", marketcap.label(), err)
		return
	}

	res, err := stmt.Exec(marketcap.ClientID, marketcap.Token, marketcap.Ticker, marketcap.Name, marketcap.Nickname, marketcap.Color, marketcap.Activity, marketcap.Decorator, marketcap.Decimals, marketcap.Currency, marketcap.CurrencySymbol, marketcap.Frequency)
	if err != nil {
		logger.Warningf("Unable to store marketcap in db %s: %s", marketcap.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store marketcap in db %s: %s", marketcap.label(), err)
		return
	}
}

// DeleteMarketCap addds a new MarketCap or crypto to the list of what to watch
func (m *Manager) DeleteMarketCap(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a marketcap")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingMarketCap[id]; !ok {
		logger.Errorf("No marketcap found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingMarketCap[id].Close <- 1
	marketcapCount.Dec()

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		stmt, err := m.DB.Prepare("DELETE FROM marketcaps WHERE name = ?")
		if err != nil {
			logger.Warningf("Unable to query marketcap in db %s: %s", id, err)
			return
		}

		_, err = stmt.Exec(m.WatchingMarketCap[id].Name)
		if err != nil {
			logger.Warningf("Unable to query marketcap in db %s: %s", id, err)
			return
		}
	}

	// remove from cache
	delete(m.WatchingMarketCap, id)

	logger.Infof("Deleted marketcap %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartMarketCap stops and starts a marketcap
func (m *Manager) RestartMarketCap(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a marketcap")

	vars := mux.Vars(r)
	id := strings.ToUpper(vars["id"])

	if _, ok := m.WatchingMarketCap[id]; !ok {
		logger.Errorf("No marketcap found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingMarketCap[id].Close <- 1

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingMarketCap[id].Frequency) * 2 * time.Second)

	// start the marketcap again
	go m.WatchingMarketCap[id].watchMarketCap()

	logger.Infof("Restarted marketcap %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetMarketCaps returns a list of what the manager is watching
func (m *Manager) GetMarketCaps(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingMarketCap); err != nil {
		logger.Errorf("Serving request: %s", err)
	}
}
