package main

import (
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
		m.addTicker(stockReq.Name, crypto)

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
	m.addTicker(stockReq.Ticker, stock)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stock)
	if err != nil {
		logger.Errorf("Unable to encode ticker: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addTicker(ticker string, stock *Ticker) {
	tickerCount.Inc()
	stock.Ticker = strings.ToUpper(stock.Ticker)
	m.WatchingTicker[strings.ToUpper(ticker)] = stock
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
