package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Manager holds a list of the crypto and stocks we are watching
type Manager struct {
	Watching map[string]*Stock
	Cache    *redis.Client
	Context  context.Context
	sync.RWMutex
}

// NewManager stores all the information about the current stocks being watched and
func NewManager(address string, count prometheus.Gauge, cache *redis.Client, context context.Context) *Manager {
	m := &Manager{
		Watching: make(map[string]*Stock),
		Cache:    cache,
		Context:  context,
	}

	// Create a router to accept requests
	r := mux.NewRouter()
	r.HandleFunc("/ticker", m.AddStock).Methods("POST")
	r.HandleFunc("/ticker/{id}", m.DeleteStock).Methods("DELETE")
	r.HandleFunc("/ticker", m.GetStocks).Methods("GET")

	// Metrics
	prometheus.MustRegister(tickerCount)
	r.Path("/metrics").Handler(promhttp.Handler())

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	logger.Infof("Starting api server on %s...", address)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	return m
}

// StockRequest represents the json coming in from the request
type StockRequest struct {
	Ticker    string `json:"ticker"`
	Token     string `json:"discord_bot_token"`
	Name      string `json:"name"`
	Nickname  bool   `json:"set_nickname"`
	Crypto    bool   `json:"crypto"`
	Color     bool   `json:"set_color"`
	Decorator string `json:"decorator" default:"-"`
	Frequency int    `json:"frequency" default:"60"`
	Currency  string `json:"currency" default:"usd"`
	Bitcoin   bool   `json:"bitcoin"`
	Activity  string `json:"activity"`
}

// AddStock adds a new stock or crypto to the list of what to watch
func (m *Manager) AddStock(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a ticker")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var stockReq StockRequest
	if err := json.Unmarshal(body, &stockReq); err != nil {
		logger.Errorf("Unmarshalling: %v", err)
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

		// check if already existing
		if _, ok := m.Watching[strings.ToUpper(stockReq.Name)]; ok {
			logger.Error("Ticker already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		crypto := NewCrypto(stockReq.Ticker, stockReq.Token, stockReq.Name, stockReq.Nickname, stockReq.Color, stockReq.Decorator, stockReq.Frequency, stockReq.Currency, stockReq.Bitcoin, stockReq.Activity, m.Cache, m.Context)
		m.addStock(stockReq.Name, crypto)
		tickerCount.Inc()

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(crypto)
		if err != nil {
			logger.Error("Unable to encode ticker: %v", err)
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
	if _, ok := m.Watching[strings.ToUpper(stockReq.Ticker)]; ok {
		logger.Error("Ticker already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	stock := NewStock(stockReq.Ticker, stockReq.Token, stockReq.Name, stockReq.Nickname, stockReq.Color, stockReq.Decorator, stockReq.Frequency, stockReq.Currency, stockReq.Activity)
	m.addStock(stockReq.Ticker, stock)
	tickerCount.Inc()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stock)
	if err != nil {
		logger.Error("Unable to encode ticker: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addStock(ticker string, stock *Stock) {
	stock.Ticker = strings.ToUpper(stock.Ticker)
	m.Watching[strings.ToUpper(ticker)] = stock
}

// DeleteStock addds a new stock or crypto to the list of what to watch
func (m *Manager) DeleteStock(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a ticker")

	vars := mux.Vars(r)
	id := strings.ToUpper(vars["id"])

	if _, ok := m.Watching[id]; !ok {
		logger.Error("No ticker found: %v", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// send shutdown sign
	m.Watching[id].Shutdown()
	tickerCount.Dec()

	// remove from cache
	delete(m.Watching, id)

	logger.Infof("Deleted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetStocks returns a list of what the manager is watching
func (m *Manager) GetStocks(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.Watching); err != nil {
		logger.Errorf("Serving request: %v", err)
	}
}
