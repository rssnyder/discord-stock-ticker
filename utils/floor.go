package utils

import (
	"fmt"
)

// GetFloorPrice based on marketplace and name
func GetFloorPrice(marketplace, name string) (float64, string, string, string, error) {
	var result float64
	var activity string
	var currency string
	var collectionStats string

	switch marketplace {
	case "magiceden":
		magiceden, err := GetMagicedenData(name)
		currency = "SOL" // ME API currently doesn't return the currency.

		if err != nil {
			return 0, activity, currency, "", err
		}

		result = magiceden.Floorprice / 1e9
		activity = "MagicEden Floor"
		collectionStats = fmt.Sprintf("Avg. Price %.2f %s (24h) | Vol. %s | List. %s", magiceden.Avgprice24Hr/1e9, currency, AmountConverter(magiceden.Volumeall/1e9), AmountConverter(magiceden.Listedcount))
	case "solsea":
		solsea, err := GetSolseaData(name)
		currency = "SOL" // Solsea currently only support solana collections.

		if err != nil {
			return 0, activity, currency, "", err
		}

		result = solsea.Floorprice
		activity = "Solsea Floor"
		collectionStats = "" // Solsea currently doesn't return the collection stats.
	case "solanart":
		solanart, err := GetSolanartData(name)
		currency = "SOL" // Solanart currently only support solana collections.

		if err != nil {
			return 0, activity, currency, "", err
		}

		result = solanart.Pagination.Floorpricefilters
		activity = "SolanArt Floor"
		collectionStats = "" // Solanart currently doesn't return the collection stats.
	default:
		opensea, err := GetOpenSeaData(name)
		currency = "ETH" // OpenSea API currently doesn't return the currency.

		if err != nil {
			return 0, activity, currency, "", err
		}

		result = opensea.Stats.FloorPrice
		activity = "OpenSea Floor"
		collectionStats = fmt.Sprintf("%.0f | %s", opensea.Stats.OneDaySales, AmountConverter(opensea.Stats.TotalSupply))
	}

	return result, activity, currency, collectionStats, nil
}
