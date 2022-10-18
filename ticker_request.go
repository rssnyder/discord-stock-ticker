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

// ImportTicker pulls in bots from the provided db
func (m *Manager) ImportTicker() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, ticker, name, nickname, color, crypto, activity, decorator, decimals, currency, currencySymbol, pair, pairFlip, multiplier, twelveDataKey, frequency FROM tickers")
	if err != nil {
		logger.Warningf("Unable to query tokens in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedTicker Ticker

		err = rows.Scan(&importedTicker.ClientID, &importedTicker.Token, &importedTicker.Ticker, &importedTicker.Name, &importedTicker.Nickname, &importedTicker.Color, &importedTicker.Crypto, &importedTicker.Activity, &importedTicker.Decorator, &importedTicker.Decimals, &importedTicker.Currency, &importedTicker.CurrencySymbol, &importedTicker.Pair, &importedTicker.PairFlip, &importedTicker.Multiplier, &importedTicker.TwelveDataKey, &importedTicker.Frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		// catch corrections
		if importedTicker.Multiplier == 0 {
			importedTicker.Multiplier = 1
		}

		// activate bot
		if importedTicker.Crypto {
			go importedTicker.watchCryptoPrice()
		} else {
			go importedTicker.watchStockPrice()
		}
		m.WatchTicker(&importedTicker)
		logger.Infof("Loaded ticker from db: %s", importedTicker.label())
	}
	rows.Close()

	// check all entries have id
	for _, ticker := range m.WatchingTicker {
		if ticker.ClientID == "" {
			id, err := getIDToken(ticker.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE tickers SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, ticker.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", ticker.label())
				ticker.ClientID = id
			}
		}
	}
}

// AddTicker takes in a new ticker from the API
func (m *Manager) AddTicker(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a ticker")

	// read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var stockReq Ticker
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

	// make sure token is valid
	if stockReq.ClientID == "" {
		id, err := getIDToken(stockReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		stockReq.ClientID = id
	}

	// ensure frequency is set
	if stockReq.Frequency <= 0 {
		stockReq.Frequency = 60
	}

	// ensure currency is set
	if stockReq.Currency == "" {
		stockReq.Currency = "usd"
	}

	// ensure multiplier is set
	if stockReq.Multiplier == 0 {
		stockReq.Multiplier = 1
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
		if _, ok := m.WatchingTicker[stockReq.label()]; ok {
			logger.Error("Ticker already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		go stockReq.watchCryptoPrice()
		m.WatchTicker(&stockReq)
	} else {
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
		if _, ok := m.WatchingTicker[stockReq.label()]; ok {
			logger.Error("Ticker already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		go stockReq.watchStockPrice()
		m.WatchTicker(&stockReq)
	}

	if *db != "" {
		m.StoreTicker(&stockReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stockReq)
	if err != nil {
		logger.Errorf("Unable to encode ticker: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added ticker: %s\n", stockReq.Ticker)
}

// WatchTicker keeps track of running
func (m *Manager) WatchTicker(ticker *Ticker) {
	tickerCount.Inc()
	id := ticker.label()
	m.WatchingTicker[id] = ticker
}

// StoreTicker puts a ticker into the db
func (m *Manager) StoreTicker(ticker *Ticker) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO tickers(clientId, token, ticker, name, nickname, color, crypto, activity, decorator, decimals, currency, currencySymbol, pair, pairFlip, multiplier, twelveDataKey, frequency) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store ticker in db %s: %s", ticker.label(), err)
		return
	}

	res, err := stmt.Exec(ticker.ClientID, ticker.Token, ticker.Ticker, ticker.Name, ticker.Nickname, ticker.Color, ticker.Crypto, ticker.Activity, ticker.Decorator, ticker.Decimals, ticker.Currency, ticker.CurrencySymbol, ticker.Pair, ticker.PairFlip, ticker.Multiplier, ticker.TwelveDataKey, ticker.Frequency)
	if err != nil {
		logger.Warningf("Unable to store ticker in db %s: %s", ticker.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store ticker in db %s: %s", ticker.label(), err)
		return
	}
}

// DeleteTicker stops and removes a ticker
func (m *Manager) DeleteTicker(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a ticker")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingTicker[id]; !ok {
		logger.Errorf("No ticker found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// send shutdown sign
	m.WatchingTicker[id].Close <- 1
	tickerCount.Dec()

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		if m.WatchingTicker[id].Crypto {
			stmt, err := m.DB.Prepare("DELETE FROM tickers WHERE name = ?")
			if err != nil {
				logger.Warningf("Unable to query ticker in db %s: %s", id, err)
				return
			}

			_, err = stmt.Exec(m.WatchingTicker[id].Name)
			if err != nil {
				logger.Warningf("Unable to query ticker in db %s: %s", id, err)
				return
			}
		} else {
			stmt, err := m.DB.Prepare("DELETE FROM tickers WHERE ticker = ?")
			if err != nil {
				logger.Warningf("Unable to query ticker in db %s: %s", id, err)
				return
			}

			_, err = stmt.Exec(m.WatchingTicker[id].Ticker)
			if err != nil {
				logger.Warningf("Unable to query ticker in db %s: %s", id, err)
				return
			}
		}
	}

	// remove from cache
	delete(m.WatchingTicker, id)

	logger.Infof("Deleted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartTicker stops and starts a ticker
func (m *Manager) RestartTicker(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a ticker")

	vars := mux.Vars(r)
	id := strings.ToUpper(vars["id"])

	if _, ok := m.WatchingTicker[id]; !ok {
		logger.Errorf("No ticker found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.WatchingTicker[id].Close <- 1

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingTicker[id].Frequency) * 2 * time.Second)

	// start the ticker again
	if m.WatchingTicker[id].Crypto {
		go m.WatchingTicker[id].watchCryptoPrice()
	} else {
		go m.WatchingTicker[id].watchStockPrice()
	}

	logger.Infof("Restarted ticker %s", id)
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
