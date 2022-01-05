package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	EthGasWatchURL = "https://ethgas.watch/api/gas"
)

type EthGasWatchData struct {
	Slow struct {
		Gwei int     `json:"gwei"`
		Usd  float64 `json:"usd"`
	} `json:"slow"`
	Normal struct {
		Gwei int     `json:"gwei"`
		Usd  float64 `json:"usd"`
	} `json:"normal"`
	Fast struct {
		Gwei int     `json:"gwei"`
		Usd  float64 `json:"usd"`
	} `json:"fast"`
	Instant struct {
		Gwei int     `json:"gwei"`
		Usd  float64 `json:"usd"`
	} `json:"instant"`
	Ethprice    float64 `json:"ethPrice"`
	Lastupdated int64   `json:"lastUpdated"`
	Sources     []struct {
		Name       string `json:"name"`
		Source     string `json:"source"`
		Fast       int    `json:"fast"`
		Standard   int    `json:"standard"`
		Slow       int    `json:"slow"`
		Lastblock  int    `json:"lastBlock,omitempty"`
		Instant    int    `json:"instant,omitempty"`
		Lastupdate int64  `json:"lastUpdate,omitempty"`
	} `json:"sources"`
}

func GetEthGasWatchData() (EthGasWatchData, error) {

	var prices EthGasWatchData

	req, err := http.NewRequest("GET", EthGasWatchURL, nil)
	if err != nil {
		return prices, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return prices, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return prices, err
	}

	err = json.Unmarshal(results, &prices)
	if err != nil {
		return prices, err
	}

	return prices, nil
}
