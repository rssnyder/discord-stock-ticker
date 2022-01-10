package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	SolanaURL = "https://api.dexlab.space/v1/trades/%s/24h"
)

type DexLabPrice struct {
	Success bool `json:"success"`
	Data    []struct {
		Orderid       string    `json:"orderId"`
		Price         string    `json:"price"`
		Size          string    `json:"size"`
		Market        string    `json:"market"`
		Side          string    `json:"side"`
		Time          time.Time `json:"time"`
		Feecost       string    `json:"feeCost"`
		Marketaddress string    `json:"marketAddress"`
		Createdat     time.Time `json:"createdAt"`
	} `json:"data"`
}

func GetDexLabPrice(address string) (string, error) {
	var price DexLabPrice
	var result string

	reqUrl := fmt.Sprintf(SolanaURL, address)

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

	if len(price.Data) == 0 {
		return result, err
	}

	return price.Data[0].Price, nil
}
