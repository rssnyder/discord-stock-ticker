package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	EtherscanURL = "https://api.etherscan.io/api?module=gastracker&action=gasoracle&apikey=%s"
)

type EtherscanResult struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}

type EtherscanData struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Result  EtherscanResult `json:"result"`
}

func GetEtherscanGasData(apiToken string) (EtherscanData, error) {

	var prices EtherscanData

	reqUrl := fmt.Sprintf(EtherscanURL, apiToken)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return prices, err
	}
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return prices, err
	}

	results, err := io.ReadAll(resp.Body)
	if err != nil {
		return prices, err
	}

	err = json.Unmarshal(results, &prices)
	if err != nil {
		return prices, err
	}

	return prices, nil
}
