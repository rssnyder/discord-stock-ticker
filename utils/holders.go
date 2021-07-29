package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	holdersUrl = "https://eth-token-holders.cloud.rileysnyder.org/%s/%s"
)

func GetHolders(chain, contract string) string {
	var holders string

	reqURL := fmt.Sprintf(holdersUrl, chain, contract)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return holders
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return holders
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return holders
	}

	holders = string(results)

	return holders
}
