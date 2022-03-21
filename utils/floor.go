package utils

import "fmt"

// GetFloorPrice based on marketplace and name
func GetFloorPrice(marketplace, name string) (string, string, error) {
	var result string
	var activity string

	switch marketplace {
	case "magiceden":
		magiceden, err := GetMagicedenData(name)
		if err != nil {
			return result, activity, err
		}
		result = fmt.Sprintf("%f SOL", magiceden.Results.Floorprice/1000000000)
	case "solsea":
		solsea, err := GetSolseaData(name)
		if err != nil {
			return result, activity, err
		}
		result = fmt.Sprintf("%f SOL", solsea.Floorprice)
		activity = "Solsea: Floor"
	case "solanart":
		solanart, err := GetSolanartData(name)
		if err != nil {
			return result, activity, err
		}
		result = fmt.Sprintf("%f SOL", solanart.Pagination.Floorpricefilters)
		activity = "SolanArt: Floor"
	default:
		opensea, err := GetOpenSeaData(name)
		if err != nil {
			return result, activity, err
		}
		result = fmt.Sprintf("Ξ%f", opensea.Stats.FloorPrice)
		activity = fmt.Sprintf("%.0f | %.2fk", opensea.Stats.OneDaySales, opensea.Stats.TotalSupply/1000)
	}

	return result, activity, nil
}
