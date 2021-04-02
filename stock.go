package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type stock struct {
	client_id int
	symbol    string
	quantity  int
}

type quote struct {
	currPrice float64
	prevClose float64
}

func fetch(url string) ([]byte, error) {
	// http client with timeout
	client := http.Client{Timeout: 10 * time.Second}

	// fetch URL
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// convert to byte slice
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func getStockQuote(symbol string) (quote, error) {
	var q quote

	resp, err := fetch("https://query2.finance.yahoo.com/v10/finance/quoteSummary/" +
		symbol + "?formatted=false&modules=price")
	if err != nil {
		return q, err
	}

	// get current stock price
	re := regexp.MustCompile(`regularMarketPrice\":[0-9]*\.[0-9]+`)
	currPrice := string(re.Find(resp))
	currPrice = strings.TrimPrefix(currPrice, "regularMarketPrice\":")
	q.currPrice, err = strconv.ParseFloat(currPrice, 64)
	if err != nil {
		return q, err
	}

	// get previous close
	re = regexp.MustCompile(`regularMarketPreviousClose\":[0-9]*\.[0-9]+`)
	prevClose := string(re.Find(resp))
	prevClose = strings.TrimPrefix(prevClose, "regularMarketPreviousClose\":")
	q.prevClose, err = strconv.ParseFloat(prevClose, 64)
	if err != nil {
		return q, err
	}

	return q, nil
}
