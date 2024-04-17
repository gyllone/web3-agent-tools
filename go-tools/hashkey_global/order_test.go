package main

import (
	"strings"
	"testing"
	"time"
)

var testAuth = HashKeyApiAuth{
	Secret: "Hy0Y5SkAB5SK28PxniiTxhUsfMImXoh96YcnZcKMtEDznA3oogHyuI9ytLq1oyVd",
	ApiKey: "gnb0CeMO43AJgkF6OwIaES7bvF0SoJm59mEu2VxsSSjHJgyX3jwYIEcSm4jDYkvd",
}

func TestCreateMakrtBuySpotOrder(t *testing.T) {
	resp, err := createSpotOrder(&CreateSpotOrderRequest{
		Symbol:   "BTCUSDT",
		Side:     SpotOrderSideEnum_BUY,
		Type:     SpotOrderTypeEnum_MARKET,
		Quantity: getPtr("0.01"),
		//Price:    getPtr("399"),
		// Amount:    getPtr("10"),
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestCreateMakrtSellSpotOrder(t *testing.T) {
	resp, err := createSpotOrder(&CreateSpotOrderRequest{
		Symbol:   "BTCUSDT",
		Side:     SpotOrderSideEnum_SELL,
		Type:     SpotOrderTypeEnum_MARKET,
		Quantity: getPtr("0.01"),
		//Price:    getPtr("399"),
		// Amount:    getPtr("10"),
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestCreateLimitSellSpotOrder(t *testing.T) {
	resp, err := createSpotOrder(&CreateSpotOrderRequest{
		Symbol:   "BTCUSDT",
		Side:     SpotOrderSideEnum_SELL,
		Type:     SpotOrderTypeEnum_LIMIT,
		Quantity: getPtr("0.01"),
		Price:    getPtr("399"),
		// Amount:    getPtr("10"),
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestQuerySplotOrder(t *testing.T) {
	resp, err := querySpotOrder(&QueryOrderRequest{
		OrderId:   "1664553449910475776",
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestQueryAllTradedOrders(t *testing.T) {
	resp, err := queryAllTradedOrders(&QueryAllTradedOrdersRequest{
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestQueryOpenSpotOrders(t *testing.T) {
	resp, err := queryCurrentOpenSoptOrder(&QueryOpenOrdersRequest{
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestCancelSpotOrder(t *testing.T) {
	resp, err := cancelSpotOrder(&CancelOrderRequest{
		OrderId:   1662421806043108352,
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestCancelMultiOrder(t *testing.T) {
	resp, err := cancelMultiSpotOrders(&CancelMultiOrdersRequest{
		IDS:       strings.Join([]string{"1662421703458821120", "1662421590246202368"}, ","),
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestCancelAllSportOrder(t *testing.T) {
	resp, err := cancelAllSoptOrders(&CancelAllOpenOrdersRequest{
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}
