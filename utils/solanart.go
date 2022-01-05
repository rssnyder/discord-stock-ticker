package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	SolanartURL = "https://qzlsklfacc.medianetwork.cloud/get_nft?collection=%s&page=0&limit=1&order=&fits=any&trait=&search=&min=0&max=0&listed=true&ownedby=&attrib_count=&bid=all"
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
		Maxpricefilters   int     `json:"maxPriceFilters"`
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
