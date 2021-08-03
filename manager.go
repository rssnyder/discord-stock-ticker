package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
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
	Cache           *redis.Client
	Context         context.Context
	sync.RWMutex
}

// NewManager stores all the information about the current stocks being watched and
func NewManager(address string, count prometheus.Gauge, cache *redis.Client, context context.Context) *Manager {
	m := &Manager{
		WatchingTicker:  make(map[string]*Ticker),
		WatchingBoard:   make(map[string]*Board),
		WatchingGas:     make(map[string]*Gas),
		WatchingToken:   make(map[string]*Token),
		WatchingHolders: make(map[string]*Holders),
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
			log.Println(err)
		}
	}()

	return m
}
