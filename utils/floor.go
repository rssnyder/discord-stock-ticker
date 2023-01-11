package utils

import (
	"fmt"
)

// GetFloorPrice based on marketplace and name
func GetFloorPrice(marketplace, name string) (float64, string, string, error) {
	var result float64
	var activity string
	var currency string

	switch marketplace {
	case "magiceden":
		magiceden, err := GetMagicedenData(name)
		currency = "SOL" // ME API currently doesn't return the currency.

		if err != nil {
			return 0, activity, currency, err
		}
		result = magiceden.Floorprice / 1000000000
		activity = "MagicEden: Floor"
	case "solsea":
		solsea, err := GetSolseaData(name)
		currency = "SOL" // Solsea currently only support solana collections.

		if err != nil {
			return 0, activity, currency, err
		}
		result = solsea.Floorprice
		activity = "Solsea: Floor"
	case "solanart":
		solanart, err := GetSolanartData(name)
		currency = "SOL" // Solanart currently only support solana collections.

		if err != nil {
			return 0, activity, currency, err
		}
		result = solanart.Pagination.Floorpricefilters
		activity = "SolanArt: Floor"
	default:
		opensea, err := GetOpenSeaData(name)
		currency = "ETH" // OpenSea API currently doesn't return the currency.

		if err != nil {
			return 0, activity, currency, err
		}
		result = opensea.Stats.FloorPrice
		activity = fmt.Sprintf("%.0f | %.2fk", opensea.Stats.OneDaySales, opensea.Stats.TotalSupply/1000)
	}

	return result, activity, currency, nil
}
