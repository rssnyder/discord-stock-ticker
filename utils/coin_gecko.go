package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	GeckoURL = "https://api.coingecko.com/api/v3/coins/%s"
)

type CurrentPrice struct {
	USD float64 `json:"usd"`
	BTC float64 `json:"btc"`
}

type MarketData struct {
	CurrentPrice               CurrentPrice `json:"current_price"`
	PriceChange                float64      `json:"price_change_24h"`
	PriceChangePercent         float64      `json:"price_change_percentage_24h"`
	PriceChangeCurrency        CurrentPrice `json:"price_change_24h_in_currency"`
	PriceChangePercentCurrency CurrentPrice `json:"price_change_percentage_24h_in_currency"`
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
func GetCryptoPriceCache(client *redis.Client, ctx context.Context, ticker string) (GeckoPriceResults, error) {
	var currentPrice CurrentPrice
	var marketData MarketData
	var geckoPriceResults GeckoPriceResults
	var symbol string
	var name string

	// CurrentPrice
	var priceFloat float64
	var priceBTCFloat float64

	price, err := client.Get(ctx, fmt.Sprintf("%s#CurrentPrice", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceFloat, err = strconv.ParseFloat(price, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	priceBTC, err := client.Get(ctx, fmt.Sprintf("%s#CurrentPriceBTC", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceBTCFloat, err = strconv.ParseFloat(priceBTC, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	currentPrice = CurrentPrice{priceFloat, priceBTCFloat}

	// PriceChangeCurrency
	var priceChangeCurrencyFloat float64
	var priceChangeCurrencyBTCFloat float64

	priceChangeCurrency, err := client.Get(ctx, fmt.Sprintf("%s#PriceChangeCurrencyUSD", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangeCurrencyFloat, err = strconv.ParseFloat(priceChangeCurrency, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	priceChangeCurrencyBTC, err := client.Get(ctx, fmt.Sprintf("%s#PriceChangeCurrencyBTC", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangeCurrencyBTCFloat, err = strconv.ParseFloat(priceChangeCurrencyBTC, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	priceChangeCurrencyFinal := CurrentPrice{priceChangeCurrencyFloat, priceChangeCurrencyBTCFloat}

	// PriceChangePercentCurrency
	var priceChangePercentCurrencyFloat float64
	var priceChangePercentCurrencyBTCFloat float64

	priceChangePercentCurrency, err := client.Get(ctx, fmt.Sprintf("%s#PriceChangeCurrencyUSD", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangePercentCurrencyFloat, err = strconv.ParseFloat(priceChangePercentCurrency, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	priceChangePercentCurrencyBTC, err := client.Get(ctx, fmt.Sprintf("%s#PriceChangeCurrencyBTC", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangePercentCurrencyBTCFloat, err = strconv.ParseFloat(priceChangePercentCurrencyBTC, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	priceChangePercentCurrencyFinal := CurrentPrice{priceChangePercentCurrencyFloat, priceChangePercentCurrencyBTCFloat}

	// price change
	priceChange, err := client.Get(ctx, fmt.Sprintf("%s#PriceChange24H", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangeFloat, err := strconv.ParseFloat(priceChange, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
		marketData = MarketData{currentPrice, priceChangeFloat, 0.00, priceChangeCurrencyFinal, priceChangePercentCurrencyFinal}
	}

	// price change percent
	priceChangePercent, err := client.Get(ctx, fmt.Sprintf("%s#PriceChangePercentage24H", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangePercentFloat, err := strconv.ParseFloat(priceChangePercent, 32)
		if err != nil {
			priceChangePercentFloat = marketData.PriceChangePercent
		}
		marketData.PriceChangePercent = priceChangePercentFloat
	}

	// symbol
	symbol, err = client.Get(ctx, fmt.Sprintf("%s#Symbol", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	}

	// name
	name, err = client.Get(ctx, fmt.Sprintf("%s#Name", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
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
