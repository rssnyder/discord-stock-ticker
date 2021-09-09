package main

import (
	"context"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
)

var (
	logger        = log.New()
	address       *string
	db            *string
	redisAddress  *string
	redisPassword *string
	redisDB       *int
	cache         *bool
	rdb           *redis.Client
	ctx           context.Context
)

func init() {
	// initialize logging
	logLevel := flag.Int("logLevel", 0, "defines the log level. 0=production builds. 1=dev builds.")
	address = flag.String("address", "localhost:8080", "address:port to bind http server to.")
	db = flag.String("db", "", "file to store tickers in")
	redisAddress = flag.String("redisAddress", "localhost:6379", "address:port for redis server.")
	redisPassword = flag.String("redisPassword", "", "redis password")
	redisDB = flag.Int("redisDB", 0, "redis db to use")
	cache = flag.Bool("cache", false, "enable cache for coingecko")
	flag.Parse()
	logger.Out = os.Stdout
	switch *logLevel {
	case 0:
		logger.SetLevel(log.InfoLevel)
	default:
		logger.SetLevel(log.DebugLevel)
	}
}

func main() {
	var wg sync.WaitGroup

	// Redis is used a an optional cache for coingecko data
	if *cache {
		rdb = redis.NewClient(&redis.Options{
			Addr:     *redisAddress,
			Password: *redisPassword,
			DB:       *redisDB,
		})
		ctx = context.Background()
	}

	// Create the bot manager
	wg.Add(1)
	NewManager(*address, *db, tickerCount, rdb, ctx)

	// wait forever
	wg.Wait()
}
