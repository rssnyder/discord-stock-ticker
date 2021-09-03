package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// TickerRequest represents the json coming in from the request
type TickerRequest struct {
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
}

// AddTicker adds a new Ticker or crypto to the list of what to watch
func (m *Manager) AddTicker(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a ticker")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var stockReq TickerRequest
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

	// add stock or crypto ticker
	if stockReq.Crypto {

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
		if _, ok := m.WatchingTicker[strings.ToUpper(stockReq.Name)]; ok {
			logger.Error("Ticker already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		crypto := NewCrypto(stockReq.Ticker, stockReq.Token, stockReq.Name, stockReq.Nickname, stockReq.Color, stockReq.Decorator, stockReq.Frequency, stockReq.Currency, stockReq.Pair, stockReq.PairFlip, stockReq.Activity, stockReq.Decimals, stockReq.CurrencySymbol, m.Cache, m.Context)
		m.addTicker(true, crypto)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(crypto)
		if err != nil {
			logger.Errorf("Unable to encode ticker: %s", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	// ensure ticker is set
	if stockReq.Ticker == "" {
		logger.Error("Ticker required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ensure name is set
	if stockReq.Name == "" {
		stockReq.Name = stockReq.Ticker
	}

	// check if already existing
	if _, ok := m.WatchingTicker[strings.ToUpper(stockReq.Ticker)]; ok {
		logger.Error("Ticker already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	stock := NewStock(stockReq.Ticker, stockReq.Token, stockReq.Name, stockReq.Nickname, stockReq.Color, stockReq.Decorator, stockReq.Frequency, stockReq.Currency, stockReq.Activity, stockReq.Decimals, stockReq.TwelveDataKey)
	m.addTicker(false, stock)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stock)
	if err != nil {
		logger.Errorf("Unable to encode ticker: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addTicker(crypto bool, stock *Ticker) {
	tickerCount.Inc()
	var id string
	if crypto {
		id = strings.ToUpper(stock.Name)
	} else {
		id = strings.ToUpper(stock.Ticker)
	}
	m.WatchingTicker[id] = stock

	var noDB *sql.DB
	if m.DB == noDB {
		return
	}

	// query
	var existingId int
	if crypto {
		stmt, err := m.DB.Prepare("SELECT id FROM tickers WHERE tickerType = 'ticker' AND name = ? LIMIT 1")
		if err != nil {
			logger.Warningf("Unable to query ticker in db %s: %s", id, err)
			return
		}

		rows, err := stmt.Query(stock.Name)
		if err != nil {
			logger.Warningf("Unable to query ticker in db %s: %s", id, err)
			return
		}

		for rows.Next() {
			err = rows.Scan(&existingId)
			if err != nil {
				logger.Warningf("Unable to query ticker in db %s: %s", id, err)
				return
			}
		}
		rows.Close()
	} else {
		stmt, err := m.DB.Prepare("SELECT id FROM tickers WHERE tickerType = 'ticker' AND ticker = ?")
		if err != nil {
			logger.Warningf("Unable to query ticker in db %s: %s", id, err)
			return
		}

		rows, err := stmt.Query(stock.Ticker)
		if err != nil {
			logger.Warningf("Unable to query ticker in db %s: %s", id, err)
			return
		}

		for rows.Next() {
			err = rows.Scan(&existingId)
			if err != nil {
				logger.Warningf("Unable to query ticker in db %s: %s", id, err)
				return
			}
		}
		rows.Close()
	}

	if existingId != 0 {

		// update entry in db
		stmt, err := m.DB.Prepare("update tickers set token = ?, ticker = ?, name = ?, nickname = ?, color = ?, crypto = ?, activity = ?, decorator = ?, decimals = ?, currency = ?, currencySymbol = ?, pair = ?, pairFlip = ?, twelveDataKey = ?, frequency = ? WHERE id = ?")
		if err != nil {
			logger.Warningf("Unable to update ticker in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(stock.token, stock.Ticker, stock.Name, stock.Nickname, stock.Color, crypto, stock.Activity, stock.Decorator, stock.Decimals, stock.Currency, stock.CurrencySymbol, stock.Pair, stock.PairFlip, stock.TwelveDataKey, stock.Frequency, existingId)
		if err != nil {
			logger.Warningf("Unable to update ticker in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to update ticker in db %s: %s", id, err)
			return
		}

		logger.Infof("Updated ticker in db %s", id)
	} else {

		// store new entry in db
		stmt, err := m.DB.Prepare("INSERT INTO tickers(tickerType, token, ticker, name, nickname, color, crypto, activity, decorator, decimals, currency, currencySymbol, pair, pairFlip, twelveDataKey, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			logger.Warningf("Unable to store ticker in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec("ticker", stock.token, stock.Ticker, stock.Name, stock.Nickname, stock.Color, crypto, stock.Activity, stock.Decorator, stock.Decimals, stock.Currency, stock.CurrencySymbol, stock.Pair, stock.PairFlip, stock.TwelveDataKey, stock.Frequency)
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

// DeleteTicker addds a new Ticker or crypto to the list of what to watch
func (m *Manager) DeleteTicker(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a ticker")

	vars := mux.Vars(r)
	id := strings.ToUpper(vars["id"])

	if _, ok := m.WatchingTicker[id]; !ok {
		logger.Errorf("No ticker found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingTicker[id].Shutdown()
	tickerCount.Dec()

	// remove from cache
	delete(m.WatchingTicker, id)

	logger.Infof("Deleted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetTickers returns a list of what the manager is watching
func (m *Manager) GetTickers(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingTicker); err != nil {
		logger.Errorf("Serving request: %s", err)
	}
}
