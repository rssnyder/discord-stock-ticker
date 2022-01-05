package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TwelveDataTS = "https://api.twelvedata.com/time_series?symbol=%s&interval=1%s&outputsize=2&apikey=%s"
)

type TimeSeries struct {
	Meta struct {
		Symbol           string `json:"symbol"`
		Interval         string `json:"interval"`
		Currency         string `json:"currency"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Exchange         string `json:"exchange"`
		Type             string `json:"type"`
	} `json:"meta"`
	Values []struct {
		Datetime string `json:"datetime"`
		Open     string `json:"open"`
		High     string `json:"high"`
		Low      string `json:"low"`
		Close    string `json:"close"`
		Volume   string `json:"volume"`
	} `json:"values"`
	Status string `json:"status"`
}

// TimeSeries retrieves the time series data for a stock
func GetTimeSeries(ticker string, interval string, apiKey string) (TimeSeries, error) {
	var price TimeSeries

	reqURL := fmt.Sprintf(TwelveDataTS, ticker, interval, apiKey)
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
