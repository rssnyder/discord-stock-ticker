package main

import (
	"context"
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Manager holds a list of the crypto and stocks we are watching
type Manager struct {
	WatchingTicker  map[string]*Ticker
	WatchingBoard   map[string]*Board
	WatchingGas     map[string]*Gas
	WatchingToken   map[string]*Token
	WatchingHolders map[string]*Holders
	DB              *sql.DB
	Cache           *redis.Client
	Context         context.Context
	sync.RWMutex
}

// NewManager stores all the information about the current stocks being watched and
func NewManager(address string, dbFile string, count prometheus.Gauge, cache *redis.Client, context context.Context) *Manager {
	m := &Manager{
		WatchingTicker:  make(map[string]*Ticker),
		WatchingBoard:   make(map[string]*Board),
		WatchingGas:     make(map[string]*Gas),
		WatchingToken:   make(map[string]*Token),
		WatchingHolders: make(map[string]*Holders),
		DB:              dbInit(dbFile),
		Cache:           cache,
		Context:         context,
	}

	// Create a router to accept requests
	r := mux.NewRouter()

	// Ticker
	r.HandleFunc("/ticker", m.AddTicker).Methods("POST")
	r.HandleFunc("/ticker/{id}", m.DeleteTicker).Methods("DELETE")
	r.HandleFunc("/ticker", m.GetTickers).Methods("GET")

	// Board
	r.HandleFunc("/tickerboard", m.AddBoard).Methods("POST")
	r.HandleFunc("/tickerboard/{id}", m.DeleteBoard).Methods("DELETE")
	r.HandleFunc("/tickerboard", m.GetBoards).Methods("GET")

	// Gas
	r.HandleFunc("/gas", m.AddGas).Methods("POST")
	r.HandleFunc("/gas/{id}", m.DeleteGas).Methods("DELETE")
	r.HandleFunc("/gas", m.GetGas).Methods("GET")

	// Token
	r.HandleFunc("/token", m.AddToken).Methods("POST")
	r.HandleFunc("/token/{id}", m.DeleteToken).Methods("DELETE")
	r.HandleFunc("/token", m.GetToken).Methods("GET")

	// Holders
	r.HandleFunc("/holders", m.AddHolders).Methods("POST")
	r.HandleFunc("/holders/{id}", m.DeleteHolders).Methods("DELETE")
	r.HandleFunc("/holders", m.GetHolders).Methods("GET")

	// Metrics
	prometheus.MustRegister(tickerCount)
	prometheus.MustRegister(boardCount)
	prometheus.MustRegister(gasCount)
	prometheus.MustRegister(tokenCount)
	prometheus.MustRegister(holdersCount)
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
			logger.Fatal(err)
		}
	}()

	return m
}

func dbInit(fileName string) *sql.DB {
	var db *sql.DB

	if fileName == "" {
		logger.Warning("Will not be storing state.")
		return db
	}

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		logger.Errorf("Unable to open db file: %s\n", err)
		logger.Warning("Will not be storing state.")
		return db
	}

	bootstrap := `CREATE TABLE IF NOT EXISTS tickers (
		id integer primary key autoincrement,
		tickerType string,
		clientId string,
		token string,
		frequency integer,
		ticker string,
		name string,
		nickname bool,
		color bool,
		crypto bool,
		activity string,
		network string,
		contract string,
		decorator string,
		decimals integer,
		header string,
		source string,
		address string,
		currency string,
		currencySymbol string,
		pair string,
		pairFlip bool,
		twelveDataKey string,
		items string
	);`

	_, err = db.Exec(bootstrap)
	if err != nil {
		logger.Errorf("Unable to bootstrap db file: %s\n", err)
		logger.Warning("Will not be storing state.")
		var dbNull *sql.DB
		return dbNull
	}

	logger.Infof("Will be storing state in %s\n", fileName)

	return db
}
