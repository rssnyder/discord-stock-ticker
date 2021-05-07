package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	YahooURL = "https://query1.finance.yahoo.com/v10/finance/quoteSummary/%s?modules=price"
)

// The following is the API response yahoo gives
type PriceResults struct {
	QuoteSummary Results `json:"quoteSummary"`
	Error        string  `json:"error"`
}

type Results struct {
	Results []Result `json:"result"`
}

type Result struct {
	Price Pricing `json:"price"`
}

type Pricing struct {
	MaxAge                     int    `json:"maxAge"`
	PreMarketChangePercent     Change `json:"preMarketChangePercent"`
	PreMarketChange            Change `json:"preMarketChange"`
	PreMarketPrice             Change `json:"preMarketPrice"`
	PreMarketSource            string `json:"preMarketSource"`
	PostMarketChangePercent    Change `json:"postMarketChangePercent"`
	PostMarketChange           Change `json:"postMarketChange"`
	PostMarketTime             int    `json:"postMarketTime"`
	PostMarketPrice            Change `json:"postMarketPrice"`
	PostMarketSource           string `json:"postMarketSource"`
	RegularMarketChangePercent Change `json:"regularMarketChangePercent"`
	RegularMarketChange        Change `json:"regularMarketChange"`
	RegularMarketTime          int    `json:"regularMarketTime"`
	RegularMarketPrice         Change `json:"regularMarketPrice"`
	RegularMarketSource        string `json:"regularMarketSource"`
	RegularMarketDayHigh       Change `json:"regularMarketDayHigh"`
	RegularMarketDayLow        Change `json:"regularMarketDayLow"`
	RegularMarketVolume        Change `json:"regularMarketVolume"`
	AverageDailyVolume10Day    Change `json:"averageDailyVolume10Day"`
	AverageDailyVolume3Month   Change `json:"averageDailyVolume3Month"`
	RegularMarketPreviousClose Change `json:"regularMarketPreviousClose"`
	RegularMarketOpen          Change `json:"regularMarketOpen"`
	StrikePrice                Change `json:"strikePrice"`
	OpenInterest               Change `json:"openInterest"`
	Exchange                   string `json:"exchange"`
	ExchangeName               string `json:"exchangeName"`
	QuoteType                  string `json:"quoteType"`
	QuoteSourceName            string `json:"quoteSourceName"`
	ExchangeDataDelayedBy      int    `json:"exchangeDataDelayedBy"`
	MarketState                string `json:"marketState"`
	Symbol                     string `json:"symbol"`
	ShortName                  string `json:"shortName"`
	LongName                   string `json:"longName"`
	Currency                   string `json:"currency"`
	CurrencySymbol             string `json:"currencySymbol"`
	PriceHint                  Change `json:"priceHint"`
	Volume24Hr                 Change `json:"volume24Hr"`
	VolumeAllCurrencies        Change `json:"volumeAllCurrencies"`
	CirculatingSupply          Change `json:"circulatingSupply"`
	MarketCap                  Change `json:"marketCap"`
}

type Change struct {
	Raw     float64 `json:"raw"`
	Fmt     string  `json:"fmt"`
	LongFmt string  `json:"longFmt,omitempty"`
}

// GetStockPrice retrieves the price of a given ticker using the yahoo API
func GetStockPrice(ticker string) (PriceResults, error) {
	var price PriceResults
	reqURL := fmt.Sprintf(YahooURL, ticker)
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

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return price, err
	}
	err = json.Unmarshal(results, &price)
	if err != nil {
		return price, err
	}
	return price, nil
}
