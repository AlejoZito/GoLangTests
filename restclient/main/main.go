package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main(){
	resp, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	defer resp.Body.Close()

	var currentPrice BitcoinCurrentPriceResponse
	decoder := json.NewDecoder(resp.Body)
	parseErr := decoder.Decode(&currentPrice)

	if parseErr == nil{
		fmt.Printf("Bitcoin price: %s %s", currentPrice.Bpi.USD.Rate, currentPrice.Bpi.USD.Code)
	}
}

type BitcoinCurrentPriceResponse struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		USD Currency `json:"USD"`
		GBP Currency `json:"GBP"`
		EUR Currency `json:"EUR"`
	} `json:"bpi"`
}

type Currency struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}