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
		m.StoreMarketCap(true, &importedMarketCap, false)
		logger.Infof("Loaded marketcap from db: %s", importedMarketCap.Name)
	}
	rows.Close()
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
	m.StoreMarketCap(true, &marketCapReq, true)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(marketCapReq)
	if err != nil {
		logger.Errorf("Unable to encode marketcaps: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added marketcap: %s\n", marketCapReq.Name)
}

func (m *Manager) StoreMarketCap(crypto bool, marketcap *MarketCap, update bool) {
	tickerCount.Inc()
	id := marketcap.label()
	m.WatchingMarketCap[id] = marketcap

	var noDB *sql.DB
	if (m.DB == noDB) || !update {
		return
	}

	// query
	var existingId int
	stmt, err := m.DB.Prepare("SELECT id FROM marketcaps WHERE name = ? LIMIT 1")
	if err != nil {
		logger.Warningf("Unable to query marketcap in db %s: %s", id, err)
		return
	}

	rows, err := stmt.Query(marketcap.Name)
	if err != nil {
		logger.Warningf("Unable to query marketcap in db %s: %s", id, err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&existingId)
		if err != nil {
			logger.Warningf("Unable to query marketcap in db %s: %s", id, err)
			return
		}
	}
	rows.Close()

	if existingId != 0 {

		// update entry in db
		stmt, err := m.DB.Prepare("update marketcaps set clientId = ?, token = ?, ticker = ?, name = ?, nickname = ?, color = ?, activity = ?, decorator = ?, decimals = ?, currency = ?, currencySymbol = ?, frequency = ? WHERE id = ?")
		if err != nil {
			logger.Warningf("Unable to update ticker in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(marketcap.ClientID, marketcap.Token, marketcap.Ticker, marketcap.Name, marketcap.Nickname, marketcap.Color, marketcap.Activity, marketcap.Decorator, marketcap.Decimals, marketcap.Currency, marketcap.CurrencySymbol, marketcap.Frequency, existingId)
		if err != nil {
			logger.Warningf("Unable to update marketcap in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to update marketcap in db %s: %s", id, err)
			return
		}

		logger.Infof("Updated marketcap in db %s", id)
	} else {

		// store new entry in db
		stmt, err := m.DB.Prepare("INSERT INTO marketcaps(clientId, token, ticker, name, nickname, color, activity, decorator, decimals, currency, currencySymbol, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			logger.Warningf("Unable to store ticker in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(marketcap.ClientID, marketcap.Token, marketcap.Ticker, marketcap.Name, marketcap.Nickname, marketcap.Color, marketcap.Activity, marketcap.Decorator, marketcap.Decimals, marketcap.Currency, marketcap.CurrencySymbol, marketcap.Frequency)
		if err != nil {
			logger.Warningf("Unable to store ticker in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to store ticker in db %s: %s", id, err)
			return
		}
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
	tickerCount.Dec()

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

	// start the ticker again
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
