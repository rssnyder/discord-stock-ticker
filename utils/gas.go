package utils

type GasData struct {
	Standard int `json:"standard"`
	Fast     int `json:"fast"`
	Instant  int `json:"instant"`
}

// get gas prices based on network
func GetGasPrices(network string) (GasData, error) {
	switch network {
	case "polygon":
		return GetZapperData(network)
	case "binance-smart-chain":
		return GetZapperData(network)
	default:
		result, err := GetEthGasWatchData()
		return GasData{
			Standard: result.Normal.Gwei,
			Fast:     result.Fast.Gwei,
			Instant:  result.Instant.Gwei,
		}, err
	}
}
