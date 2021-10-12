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

// MarketCapRequest represents the json coming in from the request
type MarketCapRequest struct {
	Ticker         string `json:"ticker"`
	Token          string `json:"discord_bot_token"`
	Name           string `json:"name"`
	Nickname       bool   `json:"set_nickname"`
	Crypto         bool   `json:"crypto"`
	Color          bool   `json:"set_color"`
	Decorator      string `json:"decorator"`
	Frequency      int    `json:"frequency"`
	Currency       string `json:"currency"`
	CurrencySymbol string `json:"currency_symbol"`
	Pair           string `json:"pair"`
	PairFlip       bool   `json:"pair_flip"`
	Activity       string `json:"activity"`
	Decimals       int    `json:"decimals"`
	TwelveDataKey  string `json:"twelve_data_key"`
	ClientID       string `json:"client_id"`
}

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
		var clientID, token, ticker, name, activity, decorator, currency, currencySymbol string
		var nickname, color bool
		var decimals, frequency int

		err = rows.Scan(&clientID, &token, &ticker, &name, &nickname, &color, &activity, &decorator, &decimals, &currency, &currencySymbol, &frequency)
		if err != nil {
			logger.Errorf("Unable to load marketcaps from db: %s", err)
			continue
		}

		// activate bot
		t := NewMarketCap(clientID, ticker, token, name, nickname, color, decorator, frequency, currency, activity, decimals, currencySymbol, lastUpdate, m.Cache, m.Context)
		m.addMarketCap(true, t, false)
		logger.Infof("Loaded marketcap from db: %s", name)
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
	var stockReq MarketCapRequest
	if err := json.Unmarshal(body, &stockReq); err != nil {
		logger.Errorf("Unmarshalling: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if stockReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure currency is set
	if stockReq.Currency == "" {
		stockReq.Currency = "usd"
	}

	// ensure name is set
	if stockReq.Name == "" {
		logger.Error("Name required for crypto")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure currency is set
	if stockReq.CurrencySymbol == "" {
		stockReq.CurrencySymbol = "$"
	}

	// check if already existing
	if _, ok := m.WatchingMarketCap[strings.ToUpper(stockReq.Name)]; ok {
		logger.Error("MarketCap already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	crypto := NewMarketCap(stockReq.ClientID, stockReq.Ticker, stockReq.Token, stockReq.Name, stockReq.Nickname, stockReq.Color, stockReq.Decorator, stockReq.Frequency, stockReq.Currency, stockReq.Activity, stockReq.Decimals, stockReq.CurrencySymbol, lastUpdate, m.Cache, m.Context)
	m.addMarketCap(true, crypto, true)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(crypto)
	if err != nil {
		logger.Errorf("Unable to encode marketcaps: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added marketcap: %s\n", crypto.Name)
}

func (m *Manager) addMarketCap(crypto bool, marketcap *MarketCap, update bool) {
	tickerCount.Inc()
	id := strings.ToUpper(marketcap.Name)
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

		res, err := stmt.Exec(marketcap.ClientID, marketcap.token, marketcap.Ticker, marketcap.Name, marketcap.Nickname, marketcap.Color, marketcap.Activity, marketcap.Decorator, marketcap.Decimals, marketcap.Currency, marketcap.CurrencySymbol, marketcap.Frequency, existingId)
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

		res, err := stmt.Exec(marketcap.ClientID, marketcap.token, marketcap.Ticker, marketcap.Name, marketcap.Nickname, marketcap.Color, marketcap.Activity, marketcap.Decorator, marketcap.Decimals, marketcap.Currency, marketcap.CurrencySymbol, marketcap.Frequency)
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
	id := strings.ToUpper(vars["id"])

	if _, ok := m.WatchingMarketCap[id]; !ok {
		logger.Errorf("No marketcap found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingMarketCap[id].Shutdown()
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
	m.WatchingMarketCap[id].Shutdown()

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingMarketCap[id].Frequency) * 2 * time.Second)

	// start the ticker again
	m.WatchingMarketCap[id].Start()

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
