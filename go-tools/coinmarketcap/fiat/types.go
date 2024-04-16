package main

type StatusResp struct {
	ErrorMessage string `json:"error_message"`
}

type Fiat struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Symbol string `json:"symbol"`
}

type FiatResp struct {
	Data   []Fiat     `json:"data"`
	Status StatusResp `json:"status"`
}
