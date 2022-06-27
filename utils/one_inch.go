package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	OneInchURL = "https://api.1inch.io/v4.0/%s/quote?fromTokenAddress=%s&toTokenAddress=%s&amount=%s"
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
		Symbol   string  `json:"symbol"`
		Name     string  `json:"name"`
		Decimals int     `json:"decimals"`
		Address  string  `json:"address"`
		Logouri  string  `json:"logoURI"`
		Tags     *string `json:"-"`
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
		currency = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC https://etherscan.io/address/0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
	case "optimism":
		networkId = "10"
		amount = "10000000000000000000"
		currency = "0x7F5c764cBc14f9669B88837ca1490cCa17c31607" // USDC https://optimistic.etherscan.io/token/0x7f5c764cbc14f9669b88837ca1490cca17c31607
	case "binance-smart-chain":
		networkId = "56"
		amount = "10000000"
		currency = "0x55d398326f99059fF775485246999027B3197955" // Binance-Peg BSC-USD (BSC-USD) https://bscscan.com/address/0x55d398326f99059ff775485246999027b3197955#code
	case "gnosis-chain":
		networkId = "100"
		amount = "10000000000000000000"
		currency = "0xDDAfbb505ad214D7b80b1f830fcCc89B60fb7A83" // USD//C from Ethereum https://blockscout.com/xdai/mainnet/token/0xDDAfbb505ad214D7b80b1f830fcCc89B60fb7A83/token-transfers
	case "polygon":
		networkId = "137"
		amount = "10000000000000000000"
		currency = "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174" // USD Coin (PoS) https://polygonscan.com/token/0x2791bca1f2de4661ed88a30c99a7a9449aa84174
	case "fantom":
		networkId = "250"
		amount = "10000000000000000000"
		currency = "0x04068DA6C83AFCFA0e13ba15A6696662335D5B75" // USDC https://ftmscan.com/token/0x04068da6c83afcfa0e13ba15a6696662335d5b75
	case "arbitrum":
		networkId = "42161"
		amount = "10000000000000000000"
		currency = "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8" // USD Coin (Arb1) https://arbiscan.io/token/0xff970a61a04b1ca14834a43f5de4533ebddb5cc8
	case "avalanche":
		networkId = "43114"
		amount = "10000000000000000000"
		currency = "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E" // USDC https://snowtrace.io/token/0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E
	default:
		networkId = "1"
		amount = "10000000000000000000"
		currency = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC https://etherscan.io/address/0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48
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

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(results, &price)
	if err != nil {
		return result, err
	}

	return price.Totokenamount, nil
}
