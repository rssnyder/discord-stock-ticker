package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	ChainURL = "https://%s/token/%s"
)

func GetHolders(chain, contract string) string {
	var holders string

	switch chain {
	case "ethereum":
		chain = "etherscan.io"
	case "binance-smart-chain":
		chain = "bscscan.com"
	default:
		chain = "etherscan.io"
	}

	reqURL := fmt.Sprintf(ChainURL, chain, contract)

	response, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("div").Each(func(index int, element *goquery.Selection) {
		exists := element.HasClass("mr-3")
		if exists {
			holders = strings.TrimSpace(element.Text())
		}
	})

	return holders
}
