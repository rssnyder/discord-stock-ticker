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

	// Redis is used a an optional cache for coingecko data
	if *cache {
		rdb = redis.NewClient(&redis.Options{
			Addr:     *redisAddress,
			Password: "",
			DB:       0,
		})
		ctx = context.Background()
	}

	// Create the bot manager
	wg.Add(1)
	m := NewManager(*address, rdb, ctx)

	// Check for inital bots
	if os.Getenv("DISCORD_BOT_TOKEN") != "" {
		s := addInitialStock()
		m.addStock(s.Ticker, s)
	}

	// wait forever
	wg.Wait()
}

// addInitialStock looks for env vars to configure a bot on boot
func addInitialStock() *Stock {
	var stock *Stock

	// Discord token is the minimum value needed
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		logger.Fatal("Discord bot token is not set! Shutting down.")
	}

	// Get settings for bootstrapped bot
	ticker := os.Getenv("TICKER")
	nickname := env.GetBoolDefault("SET_NICKNAME", false)
	color := env.GetBoolDefault("SET_COLOR", false)
	decorator := env.GetDefault("DECORATOR", "")
	frequency := env.GetIntDefault("FREQUENCY", 60)
	currency := env.GetDefault("CURRENCY", "usd")
	bitcoin := env.GetBoolDefault("BITCOIN", false)
	activity := env.GetDefault("ACTIVITY", "")

	// Check for stock name options
	var stockName string
	if name, ok := os.LookupEnv("STOCK_NAME"); ok {
		stockName = name
	} else {
		stockName = ticker
	}

	// Check if the target ticker is a crypto
	switch os.Getenv("CRYPTO_NAME") {
	case "":
		// If it's not a crypto, it's a stock
		stock = NewStock(ticker, token, stockName, nickname, color, decorator, frequency, currency, activity)
	default:
		stock = NewCrypto(ticker, token, os.Getenv("CRYPTO_NAME"), nickname, color, decorator, frequency, currency, bitcoin, activity, rdb, ctx)
	}
	return stock
}
