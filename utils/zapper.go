package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ZapperURL = "https://api.zapper.fi/v2/gas-prices?network=%s&eip1559=true"
)

type ZapperData struct {
	Eip1559  bool `json:"eip1559"`
	Standard struct {
		BaseFeePerGas        int64 `json:"baseFeePerGas"`
		MaxPriorityFeePerGas int64 `json:"maxPriorityFeePerGas"`
		MaxFeePerGas         int64 `json:"maxFeePerGas"`
	} `json:"standard"`
	Fast struct {
		BaseFeePerGas        int64 `json:"baseFeePerGas"`
		MaxPriorityFeePerGas int64 `json:"maxPriorityFeePerGas"`
		MaxFeePerGas         int64 `json:"maxFeePerGas"`
	} `json:"fast"`
	Instant struct {
		BaseFeePerGas        int64 `json:"baseFeePerGas"`
		MaxPriorityFeePerGas int64 `json:"maxPriorityFeePerGas"`
		MaxFeePerGas         int64 `json:"maxFeePerGas"`
	} `json:"instant"`
}

func GetZapperData(network string) (ZapperData, error) {

	var prices ZapperData

	reqUrl := fmt.Sprintf(ZapperURL, network)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return prices, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Basic OTZlMGNjNTEtYTYyZS00MmNhLWFjZWUtOTEwZWE3ZDJhMjQxOg==")
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
