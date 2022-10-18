package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	OneInchURL = "https://api.1inch.exchange/v3.0/%s/quote?fromTokenAddress=%s&toTokenAddress=%s&amount=%s"
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

// GetTokenPrice retrieves the price of a given ticker using the 1inch API
func Get1inchTokenPrice(network, contract string) (string, error) {
	var price ExchangeData
	var networkId string
	var amount string
	var currency string
	var result string

	// Get network id for 1inch, default to eth
	switch network {
	case "ethereum":
		networkId = "1"
		amount = "10000000000000000000"
		currency = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	case "binance-smart-chain":
		networkId = "56"
		amount = "10000000"
		currency = "0x55d398326f99059ff775485246999027b3197955"
	case "polygon":
		networkId = "137"
		amount = "10000000000000000000"
		currency = "0x2791bca1f2de4661ed88a30c99a7a9449aa84174"
	default:
		networkId = "1"
		amount = "10000000000000000000"
		currency = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	}

	reqURL := fmt.Sprintf(OneInchURL, networkId, contract, currency, amount)

	req, err := http.NewRequest("GET", reqURL, nil)
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

	results, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(results, &price)
	if err != nil {
		return result, err
	}

	return price.Totokenamount, nil
}
