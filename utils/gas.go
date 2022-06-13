package utils

type GasData struct {
	Standard int64 `json:"standard"`
	Fast     int64 `json:"fast"`
	Instant  int64 `json:"instant"`
}

// get gas prices based on network
func GetGasPrices(network string) (GasData, error) {
	result, err := GetZapperData(network)
	return GasData{
		Standard: result.Standard.BaseFeePerGas,
		Fast:     result.Fast.BaseFeePerGas,
		Instant:  result.Instant.BaseFeePerGas,
	}, err
}
