package main

import (
	"flag"
	"os"
	"sync"

	env "github.com/caitlinelfring/go-env-default"
	log "github.com/sirupsen/logrus"
)

var logger = log.New()
var currency = "usd"

func init() {
	// initialize logging
	logLevel := flag.Int("logLevel", 0, "defines the log level. 0=production builds. 1=dev builds.")
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
	wg.Add(1)
	m := NewManager()

	// TODO: remove this and opt for api only
	s := addInitialStock()
	m.addStock(s.Ticker, s)

	connectToRedis()

	// wait forever
	wg.Wait()
}

// TODO: To be removed
func addInitialStock() *Stock {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		// logger.Fatal("Discord bot token is not set! Shutting down.")
	}

	ticker := os.Getenv("TICKER")
	if ticker == "" {
		// logger.Fatal("Ticker is not set!")
	}

	// now get settings for it
	nickname := os.Getenv("SET_NICKNAME")
	color := os.Getenv("SET_COLOR")
	flashChange := os.Getenv("FLASH_CHANGE")
	frequency := env.GetIntDefault("FREQUENCY", 60)
	var stock *Stock
	switch os.Getenv("CRYPTO_NAME") {
	case "":
		// if it's not a crypto, it's a stock
		stock = NewStock(ticker, token, os.Getenv("STOCK_NAME"), nickname, color, flashChange, frequency)
	default:
		stock = NewCrypto(ticker, token, os.Getenv("CRYPTO_NAME"), nickname, color, flashChange, frequency)
	}
	return stock
}

func connectToRedis() {
	redisServer := os.Getenv("REDIS_URL")
	if redisServer == "" {
		logger.Info("No redis server specified.")
		return
	}
	// TODO: connect to redis
	/*

	       # Use redis to store stats
	       r = Redis(host=redis_server, port=6379, db=0)

	       try:
	           for server in servers:
	               r.incr(server)
	       except exceptions.ConnectionError:
	           logging.info('No redis server found, not storing stats')

	   logging.info('servers: ' + str(servers))
	*/
}
