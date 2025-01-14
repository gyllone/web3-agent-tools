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
	StartTime *int64         `url:"startTime"`
	EndTime   *int64         `url:"endTime"`
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

type QuoteTicker24HRequest struct {
	Symbol string `url:"symbol"`
}

type QuoteTicker24HR struct {
	TimeStamp             int64  `json:"t"`
	Symbol                string `json:"s"`
	LatestTradedPrice     string `json:"c"`
	HighestPrice          string `json:"h"`
	LowestPrice           string `json:"l"`
	OpeningPrice          string `json:"o"`
	HighestBidPrice       string `json:"b"`
	HighestSellingPrice   string `json:"a"`
	BaseAssetTradeVolume  string `json:"v"`
	QuoteAssetTradeVolume string `json:"qv"`
}

type QuoteTicker24HResponse []QuoteTicker24HR

// -- quote trades
type QuoteTradesRequest struct {
	Symbol string `url:"symbol"`
	Limit  int    `url:"limit"`
}

type QuoteTrades struct {
	TradedTimestamp int64 `json:"t"`
	TradedPrice     int64 `json:"p"`
	Volume          int64 `json:"q"`
	IfBuyerMaker    bool  `json:"ibm"`
}

type QuoteTradesResponse []QuoteTrades

type QuoteMergedDepthRequest struct {
	Symbol string `url:"symbol"`
	Limit  int    `url:"limit"`
	Scale  int    `url:"scale"`
}

type OrderBookLayers []string

type QuoteMergedDepthResponse struct {
	Timestamp     int64             `json:"timestamp"`
	SellingLayers []OrderBookLayers `json:"a"`
	BuyingLayers  []OrderBookLayers `json:"b"`
}

type QuoteBookTickerRequest struct {
	Symbol string `url:"symbol"`
}

type QuoteBookTicker struct {
	Symbol         string `json:"s"`
	TopBidPrice    string `json:"b"`
	TopBidQuantity string `json:"bq"`
	TopAskPrice    string `json:"a"`
	TopAskQuantity string `json:"aq"`
	Timestamp      int64  `json:"t"`
}

type QuoteBookTickerResponse []QuoteBookTicker

type QuoteTickerPriceRequest struct {
	Symbol string `json:"symbol"`
}

type QuoteTickerPrice struct {
	Symbol            string `json:"s"`
	LatestTradedPrice string `json:"p"`
}

type QuoteTickerPriceResponse []QuoteTickerPrice

type QuoteDepthRequest struct {
	Symbol string `url:"symbol"`
	Limit  int    `url:"limit"`
}

type QuoteDepth []string

type QuoteDepthResponse struct {
	Timestamp   int64        `json:"t"`
	BuyingList  []QuoteDepth `json:"b"`
	SellingList []QuoteDepth `json:"a"`
}

type ErrorMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

/*
--- auth ------
*/
type HashKeyApiAuth struct {
	AccountId string
	Secret    string
	ApiKey    string
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
	Quantity *string           `json:"quantity" url:"quantity,omitempty"`
	Amount   *string           `json:"amount" url:"amount,omitempty"`
	Price    *string           `json:"price" url:"price,omitempty"`
	// NewClientOrderId string `json:"newClientOrderId" url:"newClientOrderId"`
	// TimeInForce      string `json:"timeInForce" url:"timeInForce"`
	Timestamp int64 `json:"timestamp" url:"timestamp"`
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

// create multi spot orders
type CreateMultiSpotOrderRequest struct {
	Timestamp int64                 `url:"timestamp"`
	Orders    []*BatchSpotOrderItem `url:"-"`
}

type SpotOrderResult struct {
	Code      string `json:"code"`
	SpotOrder *Order `json:"order"`
}

// use for multi order
type BatchSpotOrderItem struct {
	Symbol        string            `json:"symbol" url:"symbol"`
	Side          SpotOrderSideEnum `json:"side" url:"side"`
	Type          SpotOrderTypeEnum `json:"type" url:"type"`
	Quantity      *string           `json:"quantity,omitempty" url:"quantity,omitempty"`
	Amount        *string           `json:"amount,omitempty" url:"amount,omitempty"`
	Price         *string           `json:"price,omitempty" url:"price,omitempty"`
	ClientOrderId string            `json:"newClientOrderId,omitempty"`
}

type CreateMultiSpotOrderResponse struct {
	Code    int                `json:"code"`
	Results []*SpotOrderResult `json:"result"`
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

type QueryOrderRequest struct {
	OrderId   string `url:"orderId"`
	Timestamp int64  `url:"timestamp"`
}

type QueryOrderResponse struct {
	AccountId          string `json:"accountId"`
	ExchangeId         string `json:"exchangeId"`
	Symbol             string `json:"symbol"`
	SymbolName         string `json:"symbolName"`
	ClientOrderId      string `json:"clientOrderId"`
	OrderId            string `json:"orderId"`
	Price              string `json:"price"` // Assuming decimal values are represented as strings
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
	AvgPrice           string `json:"avgPrice"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	StopPrice          string `json:"stopPrice"`
	IcebergQty         string `json:"icebergQty"`
	Time               string `json:"time"`
	UpdateTime         string `json:"updateTime"`
	IsWorking          bool   `json:"isWorking"`
	ReqAmount          string `json:"reqAmount"`
	FeeCoin            string `json:"feeCoin"`
	FeeAmount          string `json:"feeAmount"`
	SumFeeAmount       string `json:"sumFeeAmount"`
}

type QueryAllTradedOrdersRequest struct {
	Timestamp   int64  `url:"timestamp"`
	FromOrderId int64  `url:"fromOrderId"`
	Symbol      string `url:"symbol"`
	StartTime   int64  `url:"startTime"`
	EndTime     int64  `url:"endTime"`
	Limit       int    `url:"limit"`
}

type QueryAllTradedOrders struct {
	ExchangeId         string `json:"exchangeId"`
	Symbol             string `json:"symbol"`
	SymbolName         string `json:"symbolName"`
	ClientOrderId      string `json:"clientOrderId"`
	OrderId            string `json:"orderId"`
	Price              string `json:"price"`
	OrigQty            string `json:"origQty"`
	ExecutedQty        string `json:"executedQty"`
	CumulativeQuoteQty string `json:"cumulativeQuoteQty"`
	AvgPrice           string `json:"avgPrice"`
	Status             string `json:"status"`
	TimeInForce        string `json:"timeInForce"`
	Type               string `json:"type"`
	Side               string `json:"side"`
	StopPrice          string `json:"stopPrice"`
	IcebergQty         string `json:"icebergQty"`
	Time               string `json:"time"`
	UpdateTime         string `json:"updateTime"`
	IsWorking          bool   `json:"isWorking"`
	ReqAmount          string `json:"reqAmount"` // DECIMAL type is represented as float64 in Go
}

type QueryAllTradedOrdersResponse []QueryAllTradedOrders

/*
----  account api -------
*/
type GetAccountInfoRequest struct {
	AccountId string `url:"accountId"`
	Timestamp int64  `url:"timestamp"`
}

type AccountBalance struct {
	Asset     string `json:"asset"`
	AssetId   string `json:"assetId"`
	AssetName string `json:"assetName"`
	Total     string `json:"total"`
	Free      string `json:"free"`
	Locked    string `json:"locked"`
}

type GetAccountInfoResponse struct {
	UserId   string           `json:"userId"`
	Balances []AccountBalance `json:"balances"`
}

// get account trade list
type GetAccountTradeListRequest struct {
	Timestamp     int64   `url:"timestamp"`
	Symbol        *string `url:"symbol,omitempty"`
	StartTime     *int64  `url:"startTime,omitempty"`
	EndTime       *int64  `url:"endTime,omitempty"`
	ClientOrderId *string `url:"clientOrderId,omitempty"`
	FromId        *int64  `url:"fromId,omitempty"`
	Told          *int64  `url:"endId,omitempty"`
	Limit         *int    `url:"limit,omitempty"`
}

type AccountTradeFeeInfo struct {
	FeeCoin     string `json:"feeCoin"`
	FeeCoinID   string `json:"feeCoinId"`
	FeeCoinName string `json:"feeCoinName"`
	Fee         string `json:"fee"`
}

type AccountTrade struct {
	ID              string              `json:"id"`
	ClientOrderID   string              `json:"clientOrderId"`
	TicketID        string              `json:"ticketId"`
	Symbol          string              `json:"symbol"`
	SymbolName      string              `json:"symbolName"`
	OrderID         string              `json:"orderId"`
	MatchOrderID    string              `json:"matchOrderId,omitempty"` // 使用omitempty忽略空值或零值
	Price           string              `json:"price"`
	Qty             string              `json:"qty"`
	Commission      string              `json:"commission"`
	CommissionAsset string              `json:"commissionAsset"`
	Time            string              `json:"time"`
	IsBuyer         bool                `json:"isBuyer"`
	IsMaker         bool                `json:"isMaker"`
	Fee             AccountTradeFeeInfo `json:"fee"`
	FeeCoinID       string              `json:"feeCoinId"`
	FeeAmount       string              `json:"feeAmount"`
	MakerRebate     string              `json:"makerRebate"`
}

// todo
type GetAccountTradeListResponse []AccountTrade

type BalanceFlowType int

const (
	BalanceFlowType_Trade                 = 1
	BalanceFlowType_Fee                   = 2
	BalanceFlowType_User_Account_Transfer = 51
	BalanceFlowType_Custody_Deposit       = 900
	BalanceFlowType_Custody_Withdraw      = 904
)

// get balance flow
type GetAccountBalanceFlowRequest struct {
	Timestamp int64           `url:"timestamp"`
	StartTime int64           `url:"startTime"`
	EndTime   int64           `url:"endTime"`
	Limit     int             `url:"limit"`
	FlowType  BalanceFlowType `url:"flowType"`
}

type AccountBalanceFlow struct {
	Id            string          `json:"id"`
	AccountId     string          `json:"accountId"`
	Coin          string          `json:"coin"`
	CoinId        string          `json:"coinId"`
	CoinName      string          `json:"coinName"`
	FlowTypeValue BalanceFlowType `json:"flowTypeValue"`
	FlowType      string          `json:"flowType"`
	Change        string          `json:"change"`
	Total         string          `json:"total"`
	Created       string          `json:"created"`
}

// todo
type GetAccountBalanceFlowResponse []AccountBalanceFlow

/*
--- get deposit address -----------
*/
type GetDepositAddressRequest struct {
	Timestamp int64  `url:"timestamp"`
	Coin      string `url:"coin"`
	ChainType string `url:"chainType"`
}

type GetDepositAddressResponse struct {
	CanDeposit              bool   `json:"canDeposit"`
	Address                 string `json:"address"`
	AddressExt              string `json:"addressExt"`
	MinQuantity             string `json:"minQuantity"`
	NeedAddressTag          bool   `json:"needAddressTag"`
	RequiredConfirmTimes    int    `json:"requiredConfirmTimes"`
	CanWithdrawConfirmTimes int    `json:"canWithdrawConfirmTimes"`
	CoinType                string `json:"coinType"`
}

type GetDepositOrdersRequest struct {
	Timestamp int64  `url:"timestamp"`
	Coin      string `url:"coin"`
	StartTime int64  `url:"startTime"`
	EndTime   int64  `url:"endTime"`
	FromId    int64  `url:"fromId"`
	Limit     int    `url:"limit"`
}

type DepositOrder struct {
	Time       string `json:"time"`
	Coin       string `json:"coin"`
	CoinName   string `json:"coinName"`
	Address    string `json:"address"`
	Quantity   string `json:"quantity"`
	Status     int    `json:"status"`
	StatusCode int    `json:"statusCode"`
	TxId       string `json:"txId"`
}

type GetDepositOrdersResponse []DepositOrder

type GetWithdrawOrdersRequest struct {
	Timestamp int64   `url:"timestamp"`
	Coin      *string `url:"coin,omitempty"`
	StartTime *int64  `url:"startTime,omitempty"`
	EndTime   *int64  `url:"endTime,omitempty"`
	Remark    *string `url:"remark,omitempty"`
	Limit     *int    `url:"limit,omitempty"`
}

type WithdrawOrder struct {
	Time           string `json:"time"`           // Timestamp
	ID             string `json:"id"`             // ID
	Coin           string `json:"coin"`           // Coin
	CoinID         string `json:"coinId"`         // Coin ID
	CoinName       string `json:"coinName"`       // Coin Name
	Address        string `json:"address"`        // Withdrawal Address
	Quantity       string `json:"quantity"`       // Withdrawal amount entered by the user
	ArriveQuantity string `json:"arriveQuantity"` // Net amount received
	AddressUrl     string `json:"addressUrl"`     // Withdrawal address URL
	FeeCoinID      string `json:"feeCoinId"`      // Fee Currency ID
	FeeCoinName    string `json:"feeCoinName"`    // Fee Currency Name
	Fee            string `json:"fee"`            // Handling fee
	Remark         string `json:"remark"`         // Remark (To be released)
}

type GetWithdrawOrdersResponse []WithdrawOrder

type WithdrawRequest struct {
	Timestamp     int64   `url:"timestamp"`
	Coin          string  `url:"coin"`
	ClientOrderId string  `url:"clientOrderId"`
	Address       string  `url:"address"`
	AddressExt    string  `url:"addressExt"`
	Quantity      float64 `url:"quantity"`
	ChainType     string  `url:"chainType"`
}

type WithdrawResponse struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
	OrderId string `json:"orderId"`
}
