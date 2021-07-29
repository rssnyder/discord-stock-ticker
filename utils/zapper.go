package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GasURL = "http://api.zapper.fi/v1/gas-price?network=%s&api_key=%s"
	apiKey = "96e0cc51-a62e-42ca-acee-910ea7d2a241"
)

type GasPrices struct {
	Standard int `json:"standard"`
	Fast     int `json:"fast"`
	Instant  int `json:"instant"`
}

func GetGasPrices(network string) (GasPrices, error) {

	var prices GasPrices

	reqUrl := fmt.Sprintf(GasURL, network, apiKey)

	req, err := http.NewRequest("GET", reqUrl, nil)
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
		fmt.Printf(resp.Status)
		return prices, err
	}

	return prices, nil
}
