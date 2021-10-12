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

var (
	tickerCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "ticker_count",
			Help: "Number of tickers.",
		},
	)
	marketcapCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "marketcap_count",
			Help: "Number of marketcaps.",
		},
	)
	boardCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "board_count",
			Help: "Number of board.",
		},
	)
	gasCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "gas_count",
			Help: "Number of gas.",
		},
	)
	tokenCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "token_count",
			Help: "Number of tokens.",
		},
	)
	holdersCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "holders_count",
			Help: "Number of holders.",
		},
	)
	lastUpdate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "time_of_last_update",
			Help: "Number of seconds since the ticker last updated.",
		},
		[]string{
			"ticker",
			"type",
			"guild",
		},
	)
)

// Manager holds a list of the crypto and stocks we are watching
type Manager struct {
	WatchingTicker    map[string]*Ticker
	WatchingMarketCap map[string]*MarketCap
	WatchingBoard     map[string]*Board
	WatchingGas       map[string]*Gas
	WatchingToken     map[string]*Token
	WatchingHolders   map[string]*Holders
	DB                *sql.DB
	Cache             *redis.Client
	Context           context.Context
	sync.RWMutex
}

// NewManager stores all the information about the current stocks being watched and
func NewManager(address string, dbFile string, count prometheus.Gauge, cache *redis.Client, context context.Context) *Manager {
	m := &Manager{
		WatchingTicker:    make(map[string]*Ticker),
		WatchingMarketCap: make(map[string]*MarketCap),
		WatchingBoard:     make(map[string]*Board),
		WatchingGas:       make(map[string]*Gas),
		WatchingToken:     make(map[string]*Token),
		WatchingHolders:   make(map[string]*Holders),
		DB:                dbInit(dbFile),
		Cache:             cache,
		Context:           context,
	}

	// Create a router to accept requests
	r := mux.NewRouter()

	// Ticker
	r.HandleFunc("/ticker", m.AddTicker).Methods("POST")
	r.HandleFunc("/ticker/{id}", m.DeleteTicker).Methods("DELETE")
	r.HandleFunc("/ticker/{id}", m.RestartTicker).Methods("PATCH")
	r.HandleFunc("/ticker", m.GetTickers).Methods("GET")

	// MarketCap
	r.HandleFunc("/marketcap", m.AddMarketCap).Methods("POST")
	r.HandleFunc("/marketcap/{id}", m.DeleteMarketCap).Methods("DELETE")
	r.HandleFunc("/marketcap/{id}", m.RestartMarketCap).Methods("PATCH")
	r.HandleFunc("/marketcap", m.GetMarketCaps).Methods("GET")

	// Board
	r.HandleFunc("/tickerboard", m.AddBoard).Methods("POST")
	r.HandleFunc("/tickerboard/{id}", m.DeleteBoard).Methods("DELETE")
	r.HandleFunc("/tickerboard/{id}", m.RestartBoard).Methods("PATCH")
	r.HandleFunc("/tickerboard", m.GetBoards).Methods("GET")

	// Gas
	r.HandleFunc("/gas", m.AddGas).Methods("POST")
	r.HandleFunc("/gas/{id}", m.DeleteGas).Methods("DELETE")
	r.HandleFunc("/gas/{id}", m.RestartGas).Methods("PATCH")
	r.HandleFunc("/gas", m.GetGas).Methods("GET")

	// Token
	r.HandleFunc("/token", m.AddToken).Methods("POST")
	r.HandleFunc("/token/{id}", m.DeleteToken).Methods("DELETE")
	r.HandleFunc("/token/{id}", m.RestartToken).Methods("PATCH")
	r.HandleFunc("/token", m.GetToken).Methods("GET")

	// Holders
	r.HandleFunc("/holders", m.AddHolders).Methods("POST")
	r.HandleFunc("/holders/{id}", m.DeleteHolders).Methods("DELETE")
	r.HandleFunc("/holders/{id}", m.RestartHolders).Methods("PATCH")
	r.HandleFunc("/holders", m.GetHolders).Methods("GET")

	// Metrics
	p := prometheus.NewRegistry()
	p.MustRegister(tickerCount)
	p.MustRegister(marketcapCount)
	p.MustRegister(boardCount)
	p.MustRegister(gasCount)
	p.MustRegister(tokenCount)
	p.MustRegister(holdersCount)
	p.MustRegister(lastUpdate)
	handler := promhttp.HandlerFor(p, promhttp.HandlerOpts{})
	r.Handle("/metrics", handler)

	// Pull in existing bots
	var noDB *sql.DB
	if m.DB != noDB {
		m.ImportToken()
		m.ImportMarketCap()
		m.ImportTicker()
		m.ImportHolder()
		m.ImportGas()
		m.ImportBoard()
	}

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
		clientId string,
		token string,
		frequency integer,
		ticker string,
		name string,
		nickname bool,
		color bool,
		crypto bool,
		activity string,
		decorator string,
		decimals integer,
		currency string,
		currencySymbol string,
		pair string,
		pairFlip bool,
		twelveDataKey string
	);
	CREATE TABLE IF NOT EXISTS marketcaps (
		id integer primary key autoincrement,
		clientId string,
		token string,
		frequency integer,
		ticker string,
		name string,
		nickname bool,
		color bool,
		activity string,
		decorator string,
		decimals integer,
		currency string,
		currencySymbol string
	);
	CREATE TABLE IF NOT EXISTS tokens (
		id integer primary key autoincrement,
		clientId string,
		token string,
		frequency integer,
		name string,
		nickname bool,
		color bool,
		activity string,
		network string,
		contract string,
		decorator string,
		decimals integer,
		source string
	);
	CREATE TABLE IF NOT EXISTS holders (
		id integer primary key autoincrement,
		clientId string,
		token string,
		frequency integer,
		nickname bool,
		activity string,
		network string,
		address string
	);
	CREATE TABLE IF NOT EXISTS gases (
		id integer primary key autoincrement,
		clientId string,
		token string,
		frequency integer,
		nickname bool,
		network string
	);
	CREATE TABLE IF NOT EXISTS boards (
		id integer primary key autoincrement,
		clientId string,
		token string,
		frequency integer,
		name string,
		nickname bool,
		color bool,
		crypto bool,
		header string,
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
