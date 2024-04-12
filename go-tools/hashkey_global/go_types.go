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

type CreateSpotOrderRequest struct {
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

type CreateSpotOrderResponse struct {
	AccountId     string            `json:"accountId" url:"accountId"`
	Symbol        string            `json:"symbol" url:"symbol"`
	SymbolName    string            `json:"symbolName" url:"symbolName"`
	ClientOrderId string            `json:"clientOrderId" url:"clientOrderId"`
	OrderId       string            `json:"orderId" url:"orderId"`
	TransactTime  string            `json:"transactTime" url:"transactTime"`
	Price         string            `json:"price" url:"price"`
	OrigQty       string            `json:"origQty" url:"origQty"`
	ExecutedQty   string            `json:"executedQty" url:"executedQty"`
	Status        string            `json:"status" url:"status"`
	TimeInForce   string            `json:"timeInForce" url:"timeInForce"`
	Type          SpotOrderTypeEnum `json:"type" url:"type"`
	Side          SpotOrderSideEnum `json:"side" url:"side"`
	ReqAmount     string            `json:"reqAmount" url:"reqAmount"`
	Concentration string            `json:"concentration" url:"concentration"`
}

// cancel order
type CancelOrderRequest struct {
	OrderId       int64  `url:"orderId"`
	ClientOrderId string `url:"clientOrderId"`
	Timestamp     int64  `url:"timestamp"`
}

type CancelOrderResponse struct {
	AccountId     string            `json:"accountId"`     // 账户编号
	Symbol        string            `json:"symbol"`        // 交易对
	ClientOrderId string            `json:"clientOrderId"` // 客户定义的订单ID，如果请求中未发送，则会自动生成
	OrderId       string            `json:"orderId"`       // 系统生成的订单ID
	TransactTime  string            `json:"transactTime"`  // 交易的毫秒时间戳
	Price         string            `json:"price"`         // 价格
	OrigQty       string            `json:"origQty"`       // 数量
	ExecutedQty   string            `json:"executedQty"`   // 已交易的数量
	Status        string            `json:"status"`        // 订单状态
	TimeInForce   string            `json:"timeInForce"`   // 订单有效期
	Type          SpotOrderTypeEnum `json:"type"`          // 订单类型
	Side          SpotOrderSideEnum `json:"side"`          // 买卖方向
}

// query current open id
type QueryOpenOrdersRequest struct {
	FromOrderId int64  `url:"fromOrderId"`
	Symbol      string `url:"symbol"`
	Limit       int    `url:"limit"`
	Timestamp   int64  `url:"timestamp"`
}

type QueryOpenOrdersResponse []*Order

type Order struct {
	AccountId          string `json:"accountId"`          // 账户编号
	ExchangeId         string `json:"exchangeId"`         // 交易所编号
	Symbol             string `json:"symbol"`             // 交易对
	SymbolName         string `json:"symbolName"`         // 交易对名称
	ClientOrderId      string `json:"clientOrderId"`      // 客户定义的订单ID
	OrderId            string `json:"orderId"`            // 系统生成的订单ID
	Price              string `json:"price"`              // 价格
	OrigQty            string `json:"origQty"`            // 原始数量
	ExecutedQty        string `json:"executedQty"`        // 已执行数量
	CumulativeQuoteQty string `json:"cumulativeQuoteQty"` // 累计报价数量
	AvgPrice           string `json:"avgPrice"`           // 平均成交价格
	Status             string `json:"status"`             // 订单状态
	TimeInForce        string `json:"timeInForce"`        // 订单有效期
	Type               string `json:"type"`               // 订单类型
	Side               string `json:"side"`               // 买卖方向
	StopPrice          string `json:"stopPrice"`          // 止损价格
	IcebergQty         string `json:"icebergQty"`         // 冰山订单数量
	Time               string `json:"time"`               // 当前时间戳
	UpdateTime         string `json:"updateTime"`         // 更新时间戳
	IsWorking          bool   `json:"isWorking"`          // 是否在工作
	ReqAmount          string `json:"reqAmount"`          // 请求的现金金额
}

// cancel all open orders
type CancelAllOpenOrdersRequest struct {
	Symbol    string            `url:"symbol"`
	Side      SpotOrderSideEnum `url:"side"`
	Timestamp int64             `url:"timestamp"`
}

type CancelAllOpenOrdersReponse struct {
	Success bool `json:"success"`
}

// cancel multi-orders
type CancelMultiOrdersRequest struct {
	IDS       string `url:"ids"`
	Timestamp int64  `url:"timestamp"`
}

type CancelMultiOrdersResult struct {
	OrderId string `json:"orderId"`
	Code    string `json:"code"`
}

type CancelMultiOrdersResponse struct {
	Code   string                    `json:"code"`
	Result []CancelMultiOrdersResult `json:"result"`
}
