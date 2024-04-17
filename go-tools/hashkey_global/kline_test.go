package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestRequestKline(t *testing.T) {
	resp, err := getQuoteKline(&QuoteKlineRequest{
		Symbol:   "ETHUSDT",
		Interval: QuoteKlineInterval_1d,
		Limit:    getPtr(1),
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestParseKlineTime(t *testing.T) {
	ts := []string{
		"2023-04-05T17:45:30+08:00",
		"2024-04-02T12:04:05",
		"2024-04-02",
	}
	for _, tStr := range ts {
		value, err := time.Parse(time.RFC3339, tStr)
		if err != nil {
			t.Log(err)
			continue
		}
		t.Logf("%+v %s %d\n", value, value.Format(time.RFC3339), value.UnixMilli())
	}
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
