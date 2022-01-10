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
	buildVersion  = "development"
	address       *string
	db            *string
	frequency     *int
	redisAddress  *string
	redisPassword *string
	redisDB       *int
	cache         *bool
	managed       *bool
	version       *bool
	rdb           *redis.Client
	ctx           context.Context
)

func init() {
	logLevel := flag.Int("logLevel", 0, "defines the log level: 0=INFO 1=DEBUG")
	address = flag.String("address", "0.0.0.0:8080", "address:port to bind http server to")
	db = flag.String("db", "", "file to store tickers in")
	frequency = flag.Int("frequency", 0, "set frequency for all tickers")
	redisAddress = flag.String("redisAddress", "localhost:6379", "address:port for redis server")
	redisPassword = flag.String("redisPassword", "", "redis password")
	redisDB = flag.Int("redisDB", 0, "redis db to use")
	cache = flag.Bool("cache", false, "enable cache for coingecko")
	managed = flag.Bool("managed", false, "forcefully keep db and discord updated with bot values")
	version = flag.Bool("version", false, "print version")
	flag.Parse()

	// init logger
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

	if *version {
		logger.Infof("discord-stock-ticker@%s\n", buildVersion)
		return
	}

	logger.Infof("Running discord-stock-ticker version %s...", buildVersion)

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
