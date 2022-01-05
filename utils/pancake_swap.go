package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TokenURL = "https://api.pancakeswap.info/api/v2/tokens/%s"
)

type TokenPrice struct {
	UpdatedAt int64 `json:"updated_at"`
	Data      struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Price    string `json:"price"`
		PriceBnb string `json:"price_BNB"`
	} `json:"data"`
}

func GetPancakeTokenPrice(contract string) (string, error) {
	var price TokenPrice
	var result string

	reqUrl := fmt.Sprintf(TokenURL, contract)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(results, &price)
	if err != nil {
		return result, err
	}

	return price.Data.PriceBnb, nil
}
