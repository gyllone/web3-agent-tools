package main

import (
	"encoding/json"
	"testing"
)

func TestRequestKline(t *testing.T) {
	resp, err := getQuoteKline(&QuoteKlineRequest{
		Symbol:   "ETHUSDT",
		Interval: QuoteKlineInterval_1d,
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetQuoteTicker24hr(t *testing.T) {
	resp, err := getQuoteTicker24hr(&QuoteTicker24HRequest{
		Symbol: "BTCUSDT",
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetQuoteTrades(t *testing.T) {
	resp, err := getQuoteTrades(&QuoteTradesRequest{
		Symbol: "BTCUSDT",
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetQuoteMergedDepth(t *testing.T) {
	resp, err := getQuoteMergedDepth(&QuoteMergedDepthRequest{
		Symbol: "BTCUSDT",
		Scale:  0,
		Limit:  10,
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetQuoteBookTicker(t *testing.T) {
	resp, err := getQuoteBookTicker(&QuoteBookTickerRequest{
		// Symbol: "BTCUSDT",
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetQuoteTickerPrice(t *testing.T) {
	resp, err := getQuoteTickerPrice(&QuoteTickerPriceRequest{
		Symbol: "BTCUSDT",
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetQuoteDepth(t *testing.T) {
	resp, err := getQuoteDepth(&QuoteDepthRequest{
		Symbol: "BTCUSDT",
		Limit:  5,
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}
