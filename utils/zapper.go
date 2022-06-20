package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ZapperURL = "https://api.zapper.fi/v2/gas-prices?network=%s&eip1559=%t"
)

type ZapperEth1559Data struct {
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

type ZapperData struct {
	Eip1559  bool  `json:"eip1559"`
	Standard int64 `json:"standard"`
	Fast     int64 `json:"fast"`
	Instant  int64 `json:"instant"`
}

func GetZapperData(network string, eip1559 bool) (ZapperData, error) {

	var prices ZapperData

	reqUrl := fmt.Sprintf(ZapperURL, network, eip1559)

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

func GetZapperEth1559() (ZapperEth1559Data, error) {

	var prices ZapperEth1559Data

	reqUrl := fmt.Sprintf(ZapperURL, "ethereum", true)

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
