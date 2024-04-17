package main

import "time"

type StatusResp struct {
	ErrorMessage string `json:"error_message"`
}

type Currency struct {
	Symbol      string           `json:"symbol"`
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Amount      float64          `json:"amount"`
	LastUpdated time.Time        `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

type PriceConversionResp struct {
	Data   Currency   `json:"data"`
	Status StatusResp `json:"status"`
}
