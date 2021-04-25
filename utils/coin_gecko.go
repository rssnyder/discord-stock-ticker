package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	GeckoURL = "https://api.coingecko.com/api/v3/coins/%s"
)

type CurrentPrice struct {
	USD float64 `json:"usd"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
	PriceChange  float64      `json:"price_change_24h"`
}

// The following is the API response gecko gives
type GeckoPriceResults struct {
	ID         string     `json:"id"`
	Symbol     string     `json:"symbol"`
	Name       string     `json:"name"`
	MarketData MarketData `json:"market_data"`
}

// GetCryptoPrice retrieves the price of a given ticker using the coin gecko API
func GetCryptoPrice(ticker string) (GeckoPriceResults, error) {
	var price GeckoPriceResults

	reqURL := fmt.Sprintf(GeckoURL, ticker)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return price, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return price, err
	}

	if resp.StatusCode == 429 {
		fmt.Println("Being rate limited by coingecko")

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)
		time.Sleep(time.Duration(n) * time.Second)
		secondAttempt, err := GetCryptoPrice(ticker)
		if err != nil {
			return price, err
		} else {
			return secondAttempt, nil
		}
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return price, err
	}
	err = json.Unmarshal(results, &price)
	if err != nil {
		fmt.Printf(resp.Status)
		return price, err
	}

	return price, nil
}

// GetCryptoPriceCache attempt to use cache to get info
func GetCryptoPriceCache(client *redis.Client, ticker string) (GeckoPriceResults, error) {
	var price GeckoPriceResults

	var currentPrice CurrentPrice
	var marketData MarketData
	var geckoPriceResults GeckoPriceResults
	var symbol string
	var name string

	// coin price
	price, err := rdb.Get(ctx, Sprintf("%s#CurrentPrice", ticker)).Result()
	if err == redis.Nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	} else if err != nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	} else {
		currentPrice = currentPrice{price}
	}

	// price change
	priceChange, err := rdb.Get(ctx, Sprintf("%s#PriceChange24H", ticker)).Result()
	if err == redis.Nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	} else if err != nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	} else {
		marketData = MarketData{currentPrice, priceChange}
	}

	// symbol
	symbol, err = rdb.Get(ctx, Sprintf("%s#Symbol", ticker)).Result()
	if err == redis.Nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	} else if err != nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	}

	// name
	name, err = rdb.Get(ctx, Sprintf("%s#Name", ticker)).Result()
	if err == redis.Nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	} else if err != nil {
		price, err = GetCryptoPrice(ticker)
		return price, err
	}

	geckoPriceResults = GeckoPriceResults{
		ticker,
		symbol,
		name,
		marketData,
	}

	fmt.Println("cache hit")
	return geckoPriceResults, nil
}
