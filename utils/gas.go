package utils

type GasData struct {
	Standard int64 `json:"standard"`
	Fast     int64 `json:"fast"`
	Instant  int64 `json:"instant"`
}

// get gas prices based on network
func GetGasPrices(network string) (GasData, error) {
	switch network {
	case "ethereum":
		result, err := GetZapperEth1559()
		return GasData{
			Standard: result.Standard.BaseFeePerGas,
			Fast:     result.Fast.BaseFeePerGas,
			Instant:  result.Instant.BaseFeePerGas,
		}, err
	default:
		result, err := GetZapperData(network, true)
		return GasData{
			Standard: result.Standard,
			Fast:     result.Fast,
			Instant:  result.Instant,
		}, err
	}
}
