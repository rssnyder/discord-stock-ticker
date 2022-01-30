package main

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "modernc.org/sqlite"
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
	circulatingCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "circulating_count",
			Help: "Number of circulatings.",
		},
	)
	valuelockedCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "valuelocked_count",
			Help: "Number of valuelocked.",
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
	floorCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "floor_count",
			Help: "Number of floor.",
		},
	)
	cacheHits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_hit",
			Help: "Number of times the cache had data",
		},
	)
	cacheMisses = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_miss",
			Help: "Number of times the cache lacked data",
		},
	)
	rateLimited = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ratelimited",
			Help: "Number of times we have been rate limited",
		},
	)
	updateError = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "update_error",
			Help: "Number of times we have failed to update",
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
	WatchingTicker      map[string]*Ticker
	WatchingMarketCap   map[string]*MarketCap
	WatchingCirculating map[string]*Circulating
	WatchingValueLocked map[string]*ValueLocked
	WatchingBoard       map[string]*Board
	WatchingGas         map[string]*Gas
	WatchingToken       map[string]*Token
	WatchingHolders     map[string]*Holders
	WatchingFloor       map[string]*Floor
	DB                  *sql.DB
	Cache               *redis.Client
	Context             context.Context
	sync.RWMutex
}

// NewManager stores all the information about the current stocks being watched and
func NewManager(address string, dbFile string, count prometheus.Gauge, cache *redis.Client, context context.Context) *Manager {
	m := &Manager{
		WatchingTicker:      make(map[string]*Ticker),
		WatchingMarketCap:   make(map[string]*MarketCap),
		WatchingCirculating: make(map[string]*Circulating),
		WatchingValueLocked: make(map[string]*ValueLocked),
		WatchingBoard:       make(map[string]*Board),
		WatchingGas:         make(map[string]*Gas),
		WatchingToken:       make(map[string]*Token),
		WatchingHolders:     make(map[string]*Holders),
		WatchingFloor:       make(map[string]*Floor),
		DB:                  dbInit(dbFile),
		Cache:               cache,
		Context:             context,
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

	// Circulating
	r.HandleFunc("/circulating", m.AddCirculating).Methods("POST")
	r.HandleFunc("/circulating/{id}", m.DeleteCirculating).Methods("DELETE")
	r.HandleFunc("/circulating/{id}", m.RestartCirculating).Methods("PATCH")
	r.HandleFunc("/circulating", m.GetCirculatings).Methods("GET")

	// Value Locked
	r.HandleFunc("/valuelocked", m.AddValueLocked).Methods("POST")
	r.HandleFunc("/valuelocked/{id}", m.DeleteValueLocked).Methods("DELETE")
	r.HandleFunc("/valuelocked/{id}", m.RestartValueLocked).Methods("PATCH")
	r.HandleFunc("/valuelocked", m.GetValueLockeds).Methods("GET")

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

	// Floor
	r.HandleFunc("/floor", m.AddFloor).Methods("POST")
	r.HandleFunc("/floor/{id}", m.DeleteFloor).Methods("DELETE")
	r.HandleFunc("/floor/{id}", m.RestartFloor).Methods("PATCH")
	r.HandleFunc("/floor", m.GetFloor).Methods("GET")

	// Metrics
	p := prometheus.NewRegistry()
	p.MustRegister(tickerCount)
	p.MustRegister(marketcapCount)
	p.MustRegister(circulatingCount)
	p.MustRegister(valuelockedCount)
	p.MustRegister(boardCount)
	p.MustRegister(gasCount)
	p.MustRegister(tokenCount)
	p.MustRegister(holdersCount)
	p.MustRegister(floorCount)
	p.MustRegister(lastUpdate)
	p.MustRegister(cacheHits)
	p.MustRegister(cacheMisses)
	p.MustRegister(rateLimited)
	p.MustRegister(updateError)
	handler := promhttp.HandlerFor(p, promhttp.HandlerOpts{})
	r.Handle("/metrics", handler)

	// Pull in existing bots
	var noDB *sql.DB
	if m.DB != noDB {
		m.ImportToken()
		m.ImportMarketCap()
		m.ImportCirculating()
		m.ImportValueLocked()
		m.ImportTicker()
		m.ImportHolder()
		m.ImportGas()
		m.ImportBoard()
		m.ImportFloor()
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

	db, err := sql.Open("sqlite", fileName)
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
		multiplier integer,
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
	CREATE TABLE IF NOT EXISTS circulatings (
		id integer primary key autoincrement,
		clientId string,
		token string,
		frequency integer,
		ticker string,
		name string,
		nickname bool,
		activity string,
		decimals integer,
		currencySymbol string
	);
	CREATE TABLE IF NOT EXISTS valuelocks (
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
		currencySymbol string,
		source string
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
	);
	CREATE TABLE IF NOT EXISTS floors (
		id integer primary key autoincrement,
		clientId string,
		token string,
		frequency integer,
		nickname bool,
		marketplace string,
		name string
	);`

	_, err = db.Exec(bootstrap)
	if err != nil {
		logger.Errorf("Unable to bootstrap db file: %s\n", err)
		logger.Warning("Will not be storing state.")
		var dbNull *sql.DB
		return dbNull
	}

	// v3.8.0 - add multiplier
	_, err = db.Exec("alter table tickers add column multiplier default 1;")
	if err == nil {
		logger.Warnln("Added new column to tickers: multiplier (1)")
	} else if err.Error() == "SQL logic error: duplicate column name: multiplier (1)" {
		logger.Debug("New column already exists in tickers: multiplier (1)")
	} else if err != nil {
		logger.Errorln(err)
		logger.Warning("Will not be storing state.")
		var dbNull *sql.DB
		return dbNull
	}

	logger.Infof("Will be storing state in %s\n", fileName)

	return db
}

// getID retrive an id for a bot
func getIDToken(token string) (string, error) {
	var id string

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Errorf("Creating Discord session: %s", err)
		return id, err
	}

	botUser, err := dg.User("@me")
	if err != nil {
		logger.Errorf("Getting bot user: %s", err)
		return id, err
	}

	return botUser.ID, nil
}

// setName will update a bots name
func setName(session *discordgo.Session, name string) {

	user, err := session.User("@me")
	if err != nil {
		logger.Errorf("Getting bot user: %s", err)
		return
	}

	if user.Username == name {
		logger.Debugf("Username already matches: %s", name)
		return
	}

	_, err = session.UserUpdate("", "", name, "", "")
	if err != nil {
		logger.Errorf("Updating bot username: %s", err)
		return
	}

	logger.Debugf("%s changed to %s", user.Username, name)
}

// setRole changes color roles based on change
func setRole(session *discordgo.Session, id, guild string, increase bool) error {
	var redRole string
	var greeenRole string

	// get the roles for color changing
	roles, err := session.GuildRoles(guild)
	if err != nil {
		return err
	}

	// find role ids
	for _, r := range roles {
		if r.Name == "tickers-red" {
			redRole = r.ID
		} else if r.Name == "tickers-green" {
			greeenRole = r.ID
		}
	}

	// make sure roles exist
	if len(redRole) == 0 || len(greeenRole) == 0 {
		return errors.New("unable to find roles for color changes")
	}

	// assign role based on change
	if increase {
		err = session.GuildMemberRoleRemove(guild, id, redRole)
		if err != nil {
			return err
		}
		err = session.GuildMemberRoleAdd(guild, id, greeenRole)
		if err != nil {
			return err
		}
	} else {
		err = session.GuildMemberRoleRemove(guild, id, greeenRole)
		if err != nil {
			return err
		}
		err = session.GuildMemberRoleAdd(guild, id, redRole)
		if err != nil {
			return err
		}
	}

	return nil
}
