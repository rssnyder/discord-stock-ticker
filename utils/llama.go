package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	llamaTVLURL = "https://api.llama.fi/tvl/%s"
)

func GetLlamaTVL(slug string) (float64, error) {
	var result float64

	reqUrl := fmt.Sprintf(llamaTVLURL, slug)

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

	result, err = strconv.ParseFloat(string(results[:]), 64)
	if err != nil {
		return result, err
	}

	return result, nil
}
