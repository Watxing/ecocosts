package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"io/ioutil"
)

// this is pretty trash >_<
type Stock struct {
	QuoteSummary struct {
		Results []struct {
			Quote struct {
				CurrentPrice float64 `json:"currentPrice"`
			} `json:"financialData"`
		} `json:"result"`
		Error struct {
			Code string `json:"code"`
			Desc string `json:"description"`
		} `json:"error"`
	} `json:"quoteSummary"`
}

func getStockQuote(s string) error {
	client := http.Client{Timeout: 10 * time.Second}
	url := "https://query2.finance.yahoo.com/v10/finance/quoteSummary/"+s+"?formatted=false&modules=financialData"

	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	read, _ := ioutil.ReadAll(r.Body)
	test := new(Stock)

	err = json.Unmarshal(read, &test)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s, test)

	return nil
}

func main() {
	getStockQuote("INTC")
	getStockQuote("STOCK")
}
