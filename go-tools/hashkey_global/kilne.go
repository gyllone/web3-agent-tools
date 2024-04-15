package main

func parseKline(objects [][]interface{}) QuoteKlineResponse {

	res := make(QuoteKlineResponse, 0, len(objects))
	for _, array := range objects {
		point := &QuoteKlinePoint{
			OpenTime:                 int64(array[0].(float64)),
			OpenPrice:                array[1].(string),
			HighPrice:                array[2].(string),
			LowPrice:                 array[3].(string),
			ClosePrice:               array[4].(string),
			Volume:                   array[5].(string),
			CloseTime:                int64(array[6].(float64)),
			QuoteAssetVolume:         array[7].(string),
			NumberOfTrades:           int(array[8].(float64)),
			TakerBuyBaseAssetVolume:  array[9].(string),
			TakerBuyQuoteAssetVolume: array[10].(string),
		}
		res = append(res, point)
	}
	return res
}

func getQuoteKline(req *QuoteKlineRequest) (QuoteKlineResponse, error) {
	var objects [][]interface{}
	if err := requestWithoutSignature(req, getDeserializeJsonFunc(&objects), API_QUOTE_KLINE, "GET"); err != nil {
		return nil, err
	}
	return parseKline(objects), nil
}

func getQuoteDepth(req *QuoteDepthRequest) (*QuoteDepthResponse, error) {
	resp := QuoteDepthResponse{}
	return &resp, requestWithoutSignature(req, getDeserializeJsonFunc(&resp), API_QUOTE_DEPTH, "GET")
}

func getQuoteTicker24hr(req *QuoteTicker24HRequest) (*QuoteTicker24HResponse, error) {
	resp := QuoteTicker24HResponse{}
	return &resp, requestWithoutSignature(req, getDeserializeJsonFunc(&resp), API_QUOTE_TICKER_24HR, "GET")
}

func getQuoteTrades(req *QuoteTradesRequest) (*QuoteTradesResponse, error) {
	resp := QuoteTradesResponse{}
	return &resp, requestWithoutSignature(req, getDeserializeJsonFunc(&resp), API_QUOTE_TRADES, "GET")
}

func getQuoteMergedDepth(req *QuoteMergedDepthRequest) (*QuoteMergedDepthResponse, error) {
	resp := QuoteMergedDepthResponse{}
	return &resp, requestWithoutSignature(req, getDeserializeJsonFunc(&resp), API_QUOTE_DEPTH_MERGED, "GET")
}

func getQuoteBookTicker(req *QuoteBookTickerRequest) (*QuoteBookTickerResponse, error) {
	resp := QuoteBookTickerResponse{}
	return &resp, requestWithoutSignature(req, getDeserializeJsonFunc(&resp), API_QUOTE_BOOK_TICKER, "GET")
}

func getQuoteTickerPrice(req *QuoteTickerPriceRequest) (*QuoteTickerPriceResponse, error) {
	resp := QuoteTickerPriceResponse{}
	return &resp, requestWithoutSignature(req, getDeserializeJsonFunc(&resp), API_QUOTE_TICKER_PRICE, "GET")
}
