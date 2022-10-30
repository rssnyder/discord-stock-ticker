package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ZapperURL = "https://api.zapper.fi/v2/gas-prices?network=%s&eip1559=%t&api_key=%s"
)

type ZapperData struct {
	Eip1559  bool  `json:"eip1559"`
	Standard int64 `json:"standard"`
	Fast     int64 `json:"fast"`
	Instant  int64 `json:"instant"`
}

func GetZapperData(network string, eip1559 bool, apiKey string) (ZapperData, error) {

	var prices ZapperData

	reqUrl := fmt.Sprintf(ZapperURL, network, eip1559, apiKey)

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
