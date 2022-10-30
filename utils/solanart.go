package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	SolanartURL = "https://api-v2.solanart.io/nfts?collection_id=%s&page=0&limit=1&attributes=[]&query=&listed=all&order_by=price&direction=ASC"
)

type SolanartCollection struct {
	Pagination struct {
		Currentpage       int     `json:"currentPage"`
		Perpage           int     `json:"perPage"`
		Nextpage          int     `json:"nextPage"`
		Maxpages          int     `json:"maxPages"`
		Maxitems          int     `json:"maxItems"`
		Owners            int     `json:"Owners"`
		Floorpricefilters float64 `json:"floorPriceFilters"`
		Maxpricefilters   float64 `json:"maxPriceFilters"`
	} `json:"pagination"`
}

func GetSolanartData(collection string) (SolanartCollection, error) {
	var result SolanartCollection

	reqUrl := fmt.Sprintf(SolanartURL, collection)

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
