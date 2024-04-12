package main

type QuoteKlineENUM string

const (
	QuoteKlineInterval_1min   QuoteKlineENUM = "1m"
	QuoteKlineInterval_1h     QuoteKlineENUM = "1h"
	QuoteKlineInterval_4h     QuoteKlineENUM = "4h"
	QuoteKlineInterval_1d     QuoteKlineENUM = "1d"
	QuoteKlineInterval_1w     QuoteKlineENUM = "1w"
	QuoteKlineInterval_1Month QuoteKlineENUM = "1M"
)

type QuoteKlineRequest struct {
	Symbol    string         `url:"symbol"`
	Interval  QuoteKlineENUM `url:"interval"`
	Limit     *int           `url:"limit"`
	StartTime *int           `url:"startTime"`
	EndTime   *int           `url:"endTime"`
}

type QuoteKlinePoint struct {
	OpenTime                 int64
	OpenPrice                string
	HighPrice                string
	LowPrice                 string
	ClosePrice               string
	Volume                   string
	CloseTime                int64
	QuoteAssetVolume         string
	NumberOfTrades           int
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
}

type QuoteKlineResponse []*QuoteKlinePoint

type ErrorMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

/*
--- auth ------
*/
type HashKeyAuth struct {
	Secret string
	ApiKey string
}

/*
--- spot order api ------
*/
type SpotOrderSideEnum string

const (
	SpotOrderSideEnum_BUY  = "BUY"
	SpotOrderSideEnum_SELL = "SELL"
)

type SpotOrderTypeEnum string

const (
	SpotOrderTypeEnum_LIMIT       = "LIMIT"
	SpotOrderTypeEnum_MARKET      = "MARKET"
	SpotOrderTypeEnum_LIMIT_MAKER = "LIMIT_MAKER"
)

type SpotOrderRequest struct {
	Symbol   string            `json:"symbol" url:"symbol"`
	Side     SpotOrderSideEnum `json:"side" url:"side"`
	Type     SpotOrderTypeEnum `json:"type" url:"type"`
	Quantity float64           `json:"quantity" url:"quantity"`
	//Amount           float64           `json:"amount" url:"amount"`
	Price            float64 `json:"price" url:"price"`
	NewClientOrderId string  `json:"newClientOrderId" url:"newClientOrderId"`
	TimeInForce      string  `json:"timeInForce" url:"timeInForce"`
	Timestamp        int64   `json:"timestamp" url:"timestamp"`
}

type SpotOrderResponse struct {
	AccountId     string            `json:"accountId" url:"accountId"`
	Symbol        string            `json:"symbol" url:"symbol"`
	SymbolName    string            `json:"symbolName" url:"symbolName"`
	ClientOrderId string            `json:"clientOrderId" url:"clientOrderId"`
	OrderId       int64             `json:"orderId" url:"orderId"`
	TransactTime  int64             `json:"transactTime" url:"transactTime"`
	Price         float64           `json:"price" url:"price"`
	OrigQty       float64           `json:"origQty" url:"origQty"`
	ExecutedQty   float64           `json:"executedQty" url:"executedQty"`
	Status        string            `json:"status" url:"status"`
	TimeInForce   string            `json:"timeInForce" url:"timeInForce"`
	Type          SpotOrderTypeEnum `json:"type" url:"type"`
	Side          SpotOrderSideEnum `json:"side" url:"side"`
	ReqAmount     float64           `json:"reqAmount" url:"reqAmount"`
	Concentration string            `json:"concentration" url:"concentration"`
}
