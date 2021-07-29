package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	OneInchURL = "https://api.1inch.exchange/v3.0/137/quote?fromTokenAddress=%s&toTokenAddress=%s&amount=10000000000000000000"
)

// The following is the API response 1inch gives
type ExchangeData struct {
	Fromtoken struct {
		Address  string `json:"address"`
		Decimals int    `json:"decimals"`
		Symbol   string `json:"symbol"`
		Name     string `json:"name"`
		Logouri  string `json:"logoURI"`
		Iscustom bool   `json:"isCustom"`
	} `json:"fromToken"`
	Totoken struct {
		Symbol   string `json:"symbol"`
		Name     string `json:"name"`
		Decimals int    `json:"decimals"`
		Address  string `json:"address"`
		Logouri  string `json:"logoURI"`
	} `json:"toToken"`
	Totokenamount   string `json:"toTokenAmount"`
	Fromtokenamount string `json:"fromTokenAmount"`
	Protocols       [][][]struct {
		Name             string `json:"name"`
		Part             int    `json:"part"`
		Fromtokenaddress string `json:"fromTokenAddress"`
		Totokenaddress   string `json:"toTokenAddress"`
	} `json:"protocols"`
	Estimatedgas int `json:"estimatedGas"`
}

// GetMaticPrice retrieves the price of a given ticker using the 1inch API
func GetMaticPrice(contract, currency string) (ExchangeData, error) {
	var price ExchangeData

	reqURL := fmt.Sprintf(OneInchURL, contract, currency)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return price, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return price, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return price, err
	}
	err = json.Unmarshal(results, &price)
	if err != nil {
		return price, err
	}
	return price, nil
}
