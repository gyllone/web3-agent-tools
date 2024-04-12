package main

import "fmt"

// TODO
var APIEndPoint = "https://api.qa1.bmuxdc.com"

var API_KLINE string
var API_SPOT_ORDER string
var API_SPOT_OPENORDERS string
var API_SPOT_CANCEL_ORDER_BY_IDS string

func init() {
	API_KLINE = fmt.Sprintf("%s/quote/v1/klines", APIEndPoint)
	API_SPOT_ORDER = fmt.Sprintf("%s/api/v1/spot/order", APIEndPoint)
	API_SPOT_OPENORDERS = fmt.Sprintf("%s/api/v1/spot/openOrders", APIEndPoint)
	API_SPOT_CANCEL_ORDER_BY_IDS = fmt.Sprintf("%s/api/v1/spot/cancelOrderByIds", APIEndPoint)
}
