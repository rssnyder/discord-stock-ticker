package utils

import "fmt"

// GetFloorPrice based on marketplace and name
func GetFloorPrice(marketplace, name string) (string, error) {
	var result string

	switch marketplace {
	case "solsea":
		solsea, err := GetSolseaData(name)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("%f SOL", solsea.Floorprice)
	case "solanart":
		solanart, err := GetSolanartData(name)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("%f SOL", solanart.Pagination.Floorpricefilters)
	default:
		opensea, err := GetOpenSeaData(name)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("%f ETH", opensea.Stats.FloorPrice)
	}

	return result, nil
}
