package main

import (
	"strings"
	"testing"
	"time"
)

var testAuth = HashKeyApiAuth{
	Secret: "4FcFv4f3kIsfpVh2HcL7rES9iLN3nfTubeZvzjb97jXnWpndvvfFLgHIAqrNEJip",
	ApiKey: "AAyrJfY2VBLYCIBNa8k8kbppXAeBzF1i9Tmz2DawdEbWLzzBjGj0bRKFW0q0F3cd",
}

func TestCreateSpotOrder(t *testing.T) {
	resp, err := createSpotOrder(&CreateSpotOrderRequest{
		Symbol:    "BTCUSDT",
		Side:      "BUY",
		Type:      SpotOrderTypeEnum_LIMIT,
		Quantity:  0.0001,
		Price:     399,
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
