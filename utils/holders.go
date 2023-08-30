package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ChainURL = "https://api.covalenthq.com/v1/%s/tokens/%s/token_holders_v2/"
)

type Holders struct {
	Data struct {
		UpdatedAt time.Time `json:"updated_at"`
		ChainID   int       `json:"chain_id"`
		ChainName string    `json:"chain_name"`
		Items     []struct {
			ContractDecimals     int      `json:"contract_decimals"`
			ContractName         string   `json:"contract_name"`
			ContractTickerSymbol string   `json:"contract_ticker_symbol"`
			ContractAddress      string   `json:"contract_address"`
			SupportsErc          []string `json:"supports_erc"`
			LogoURL              string   `json:"logo_url"`
			Address              string   `json:"address"`
			Balance              string   `json:"balance"`
			TotalSupply          string   `json:"total_supply"`
			BlockHeight          int      `json:"block_height"`
		} `json:"items"`
		Pagination struct {
			HasMore    bool `json:"has_more"`
			PageNumber int  `json:"page_number"`
			PageSize   int  `json:"page_size"`
			TotalCount int  `json:"total_count"`
		} `json:"pagination"`
	} `json:"data"`
	Error        bool        `json:"error"`
	ErrorMessage interface{} `json:"error_message"`
	ErrorCode    interface{} `json:"error_code"`
}

func GetHolders(chain, contract, apiKey string) (int, error) {
	var holders Holders
	var result int

	switch chain {
	case "ethereum":
		chain = "eth-mainnet"
	case "binance-smart-chain":
		chain = "bsc-mainnet"
	default:
		chain = "eth-mainnet"
	}

	reqUrl := fmt.Sprintf(ChainURL, chain, contract)

	req, err := http.NewRequest("GET", reqUrl, nil)
	req.SetBasicAuth(apiKey, "")
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

	err = json.Unmarshal(results, &holders)
	if err != nil {
		return result, err
	}

	if holders.Error {
		return result, errors.New(fmt.Sprintf("%v", holders.ErrorMessage))
	}

	return holders.Data.Pagination.TotalCount, err
}
