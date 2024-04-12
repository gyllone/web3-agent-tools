package main

import "fmt"

// TODO
var APIEndPoint = "https://api.qa1.bmuxdc.com"

var API_KLINE string
var API_SPOT_ORDER string

func init() {
	API_KLINE = fmt.Sprintf("%s/quote/v1/klines", APIEndPoint)
	API_SPOT_ORDER = fmt.Sprintf("%s/api/v1/spot/order", APIEndPoint)
}
