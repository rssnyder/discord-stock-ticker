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
	CurrentPrice            CurrentPrice `json:"current_price"`
	MarketCap               CurrentPrice `json:"market_cap"`
	PriceChangePercent      float64      `json:"price_change_percentage_24h"`
	PriceChangeCurrency     CurrentPrice `json:"price_change_24h_in_currency"`
	MarketCapChangePercent  float64      `json:"market_cap_change_percentage_24h"`
	MarketCapChangeCurrency CurrentPrice `json:"market_cap_change_24h_in_currency"`
	TotalSupply             float64      `json:"total_supply"`
	CirculatingSupply       float64      `json:"circulating_supply"`
	MarketCapRank           int64        `json:"market_cap_rank"`
}


// The following is the API response gecko gives
type GeckoPriceResults struct {
	ID            string     `json:"id"`
	Symbol        string     `json:"symbol"`
	Name          string     `json:"name"`
	MarketData    MarketData `json:"market_data"`
	MarketCapRank int64    'json:"market_cap_rank"'
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
	var currentMarketCap CurrentPrice
	var marketData MarketData
	var marketCapRank int64
	var geckoPriceResults GeckoPriceResults
	var symbol string
	var name string
	var totalSupply string
	var circulatingSupply string

	// get CurrentPrice
	var priceFloat float64
	price, err := client.Get(ctx, fmt.Sprintf("%s#CurrentPrice", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
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
	
	// get marketCapRank
	var marketCapRankFloat int64
	marketCapRank, err := client.Get(ctx, fmt.sprintF("%s#marketCapRank", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.PrintIn("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		marketCapRankFloat, err = strconv.ParseFloat(marketCapRank, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}
		

	
	// get MarketCap
	var marketCapFloat float64
	marketCap, err := client.Get(ctx, fmt.Sprintf("%s#MarketCap", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		marketCapFloat, err = strconv.ParseFloat(marketCap, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	// get current btc price
	var btcFloat float64
	btc, err := client.Get(ctx, "bitcoin#CurrentPrice").Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss on btc")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		btcFloat, err = strconv.ParseFloat(btc, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	currentPrice = CurrentPrice{priceFloat, priceFloat / btcFloat}
	currentMarketCap = CurrentPrice{marketCapFloat, marketCapFloat / btcFloat}

	// price change
	var priceChangeFloat float64
	priceChange, err := client.Get(ctx, fmt.Sprintf("%s#PriceChange24H", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangeFloat, err = strconv.ParseFloat(priceChange, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	priceChangeCurrency := CurrentPrice{priceChangeFloat, priceChangeFloat / btcFloat}

	// marketCap change
	var marketCapChangeFloat float64
	marketCapChange, err := client.Get(ctx, fmt.Sprintf("%s#MarketCapChange24H", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		marketCapChangeFloat, err = strconv.ParseFloat(marketCapChange, 32)
		if err != nil {
			geckoPriceResults, err = GetCryptoPrice(ticker)
			return geckoPriceResults, err
		}
	}

	marketCapChangeCurrency := CurrentPrice{marketCapChangeFloat, marketCapChangeFloat / btcFloat}

	// price change percent
	var priceChangePercentFloat float64
	priceChangePercent, err := client.Get(ctx, fmt.Sprintf("%s#PriceChangePercentage24H", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		priceChangePercentFloat, err = strconv.ParseFloat(priceChangePercent, 32)
		if err != nil {
			priceChangePercentFloat = 0.00
		}
	}

	// marketCap change percent
	var marketCapChangePercentFloat float64
	marketCapChangePercent, err := client.Get(ctx, fmt.Sprintf("%s#MarketCapChangePercentage24H", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		marketCapChangePercentFloat, err = strconv.ParseFloat(marketCapChangePercent, 32)
		if err != nil {
			priceChangePercentFloat = 0.00
		}
	}

	// totalSupply
	var totalSupplyFloat float64
	totalSupply, err = client.Get(ctx, fmt.Sprintf("%s#TotalSupply", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		totalSupplyFloat, err = strconv.ParseFloat(totalSupply, 32)
		if err != nil {
			totalSupplyFloat = 0.00
		}
	}

	// name
	var circulatingSupplyFloat float64
	circulatingSupply, err = client.Get(ctx, fmt.Sprintf("%s#CirculatingSupply", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	} else {
		circulatingSupplyFloat, err = strconv.ParseFloat(circulatingSupply, 32)
		if err != nil {
			circulatingSupplyFloat = 0.00
		}
	}

	marketData = MarketData{
		currentPrice,
		currentMarketCap,
		priceChangePercentFloat,
		priceChangeCurrency,
		marketCapChangePercentFloat,
		marketCapChangeCurrency,
		totalSupplyFloat,
		circulatingSupplyFloat,
	}

	// symbol
	symbol, err = client.Get(ctx, fmt.Sprintf("%s#Symbol", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
		return geckoPriceResults, err
	} else if err != nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		return geckoPriceResults, err
	}

	// name
	name, err = client.Get(ctx, fmt.Sprintf("%s#Name", ticker)).Result()
	if err == redis.Nil {
		geckoPriceResults, err = GetCryptoPrice(ticker)
		fmt.Println("cache miss")
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
