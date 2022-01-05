package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ZapperURL = "http://api.zapper.fi/v1/gas-price?network=%s&api_key=%s"
	apiKey    = "96e0cc51-a62e-42ca-acee-910ea7d2a241"
)

func GetZapperData(network string) (GasData, error) {

	var prices GasData

	reqUrl := fmt.Sprintf(ZapperURL, network, apiKey)

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
		return prices, err
	}

	return prices, nil
}
