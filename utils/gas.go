package utils

import (
	"strconv"
)

type GasData struct {
	Standard int64 `json:"standard"`
	Fast     int64 `json:"fast"`
	Instant  int64 `json:"instant"`
}

// get gas prices based on network
func GetGasPrices(network string, apiToken string) (GasData, error) {
	switch network {
	case "ethereum":
		result, err := GetEtherscanGasData(apiToken)
		if err != nil {
			return GasData{}, err
		}
		safe, err := strconv.ParseInt(result.Result.SafeGasPrice, 10, 64)
		if err != nil {
			safe = 0
		}
		propose, err := strconv.ParseInt(result.Result.ProposeGasPrice, 10, 64)
		if err != nil {
			propose = 0
		}
		fast, err := strconv.ParseInt(result.Result.FastGasPrice, 10, 64)
		if err != nil {
			fast = 0
		}
		return GasData{
			Standard: safe,
			Fast:     propose,
			Instant:  fast,
		}, err
	default:
		result, err := GetZapperData(network, true, apiToken)
		return GasData{
			Standard: result.Standard,
			Fast:     result.Fast,
			Instant:  result.Instant,
		}, err
	}
}
