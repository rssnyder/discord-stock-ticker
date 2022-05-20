package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	SolseaURL = "https://solsea.io/collection/%s"
)

type SolseaCollection struct {
	Floorprice float64
}

func GetSolseaData(collection string) (SolseaCollection, error) {
	var result SolseaCollection
	var re = regexp.MustCompile(`<span>Floor</span>.*</span>`)

	reqUrl := fmt.Sprintf(SolseaURL, collection)

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

	var matches = re.FindStringSubmatch(string(results))
	if len(matches) == 0 {
		return result, errors.New("can't parse page")
	}

	var first = strings.Split(matches[0], "</span>")[1]
	var second = strings.Split(first, ">")[1]	

	f, err := strconv.ParseFloat(second, 64)
	if err != nil {
		return result, err
	}

	result.Floorprice = f

	return result, nil
}
