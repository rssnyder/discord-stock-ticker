package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	magicedenEscrowURL = "https://api-mainnet.magiceden.io/rpc/getCollectionEscrowStats/%s"
)

type MagicedenEscrow struct {
	Results struct {
		Symbol                   string `json:"symbol"`
		Enabledattributesfilters bool   `json:"enabledAttributesFilters"`
		Availableattributes      []struct {
			Count     int `json:"count"`
			Floor     int `json:"floor"`
			Attribute struct {
				TraitType string `json:"trait_type"`
				Value     string `json:"value"`
			} `json:"attribute"`
		} `json:"availableAttributes"`
		Floorprice       float64 `json:"floorPrice"`
		Listedcount      float64 `json:"listedCount"`
		Listedtotalvalue float64 `json:"listedTotalValue"`
		Avgprice24Hr     float64 `json:"avgPrice24hr"`
		Volume24Hr       float64 `json:"volume24hr"`
		Volumeall        float64 `json:"volumeAll"`
	} `json:"results"`
}

func GetMagicedenData(collection string) (MagicedenEscrow, error) {
	var result MagicedenEscrow

	reqUrl := fmt.Sprintf(magicedenEscrowURL, collection)

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

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(results, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
