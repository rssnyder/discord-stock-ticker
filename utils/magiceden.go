package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	magicedenURL = "http://api-mainnet.magiceden.dev/v2/collections/%s/stats"
)

type MagicedenOpenSeaCollection struct {
	Symbol       string  `json:"symbol"`
	Floorprice   float64 `json:"floorPrice"`
	Listedcount  float64 `json:"listedCount"`
	Avgprice24Hr float64 `json:"avgPrice24hr"`
	Volumeall    float64 `json:"volumeAll"`
}

func GetMagicedenData(collection string) (MagicedenOpenSeaCollection, error) {
	var result MagicedenOpenSeaCollection

	reqUrl := fmt.Sprintf(magicedenURL, collection)

	req, err := http.NewRequest("GET", reqUrl, nil)
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

	err = json.Unmarshal(results, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
