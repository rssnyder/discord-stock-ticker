package main

import (
	"context"
	"flag"
	"os"
	"sync"

	env "github.com/caitlinelfring/go-env-default"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var (
	logger       = log.New()
	address      *string
	redisAddress *string
	cache        *bool
	rdb          *redis.Client
	ctx          context.Context
)

func init() {
	// initialize logging
	logLevel := flag.Int("logLevel", 0, "defines the log level. 0=production builds. 1=dev builds.")
	address = flag.String("address", "localhost:8080", "address:port to bind http server to.")
	redisAddress = flag.String("redisAddress", "localhost:6379", "address:port for redis server.")
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

	if *cache {
		rdb = redis.NewClient(&redis.Options{
			Addr:     *redisAddress,
			Password: "",
			DB:       0,
		})
		ctx = context.Background()
	}

	wg.Add(1)
	m := NewManager(*address, rdb, ctx)

	// check for inital bots
	if os.Getenv("DISCORD_BOT_TOKEN") != "" {
		s := addInitialStock()
		m.addStock(s.Ticker, s)
	}

	// wait forever
	wg.Wait()
}

func addInitialStock() *Stock {
	var stock *Stock

	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		logger.Fatal("Discord bot token is not set! Shutting down.")
	}

	ticker := os.Getenv("TICKER")

	// now get settings for it
	nickname := env.GetBoolDefault("SET_NICKNAME", false)
	color := env.GetBoolDefault("SET_COLOR", false)
	percentage := env.GetBoolDefault("PERCENTAGE", false)
	arrows := env.GetBoolDefault("ARROWS", false)
	decorator := env.GetDefault("DECORATOR", "-")
	frequency := env.GetIntDefault("FREQUENCY", 60)
	currency := env.GetDefault("CURRENCY", "usd")

	var stockName string
	if name, ok := os.LookupEnv("STOCK_NAME"); ok {
		stockName = name
	} else {
		stockName = ticker
	}

	switch os.Getenv("CRYPTO_NAME") {
	case "":
		// if it's not a crypto, it's a stock
		stock = NewStock(ticker, token, stockName, nickname, color, percentage, arrows, decorator, frequency, currency)
	default:
		stock = NewCrypto(ticker, token, os.Getenv("CRYPTO_NAME"), nickname, color, percentage, arrows, decorator, frequency, currency, rdb, ctx)
	}
	return stock
}
