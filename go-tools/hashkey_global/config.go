package main

import "fmt"

// TODO
var APIEndPoint = "https://api.qa1.bmuxdc.com"

var API_QUOTE_KLINE string
var API_QUOTE_TICKER_24HR string
var API_QUOTE_TRADES string
var API_QUOTE_DEPTH_MERGED string
var API_QUOTE_BOOK_TICKER string
var API_QUOTE_TICKER_PRICE string
var API_QUOTE_DEPTH string

var API_SPOT_ORDER string
var API_SPOT_OPENORDERS string
var API_SPOT_CANCEL_ORDER_BY_IDS string

var API_ACCOUNT string
var API_ACCOUNT_TRADE_LIST string
var API_ACCOUNT_BALANCE_FLOW string

func init() {
	API_QUOTE_KLINE = fmt.Sprintf("%s/quote/v1/klines", APIEndPoint)
	API_QUOTE_DEPTH = fmt.Sprintf("%s/quote/v1/depth", APIEndPoint)
	API_QUOTE_TICKER_24HR = fmt.Sprintf("%s/quote/v1/ticker/24hr", APIEndPoint)
	API_QUOTE_TRADES = fmt.Sprintf("%s/quote/v1/trades", APIEndPoint)
	API_QUOTE_DEPTH_MERGED = fmt.Sprintf("%s/quote/v1/depth/merged", APIEndPoint)
	API_QUOTE_BOOK_TICKER = fmt.Sprintf("%s/quote/v1/ticker/bookTicker", APIEndPoint)
	API_QUOTE_TICKER_PRICE = fmt.Sprintf("%s/quote/v1/ticker/price", APIEndPoint)

	API_SPOT_ORDER = fmt.Sprintf("%s/api/v1/spot/order", APIEndPoint)
	API_SPOT_OPENORDERS = fmt.Sprintf("%s/api/v1/spot/openOrders", APIEndPoint)
	API_SPOT_CANCEL_ORDER_BY_IDS = fmt.Sprintf("%s/api/v1/spot/cancelOrderByIds", APIEndPoint)

	API_ACCOUNT = fmt.Sprintf("%s/api/v1/account", APIEndPoint)
	API_ACCOUNT_TRADE_LIST = fmt.Sprintf("%s/api/v1/account/trades", APIEndPoint)
	API_ACCOUNT_BALANCE_FLOW = fmt.Sprintf("%s/api/v1/account/balanceFlow", APIEndPoint)
}
