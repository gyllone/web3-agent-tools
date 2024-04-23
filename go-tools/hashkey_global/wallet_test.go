package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestGetDepositAddress(t *testing.T) {
	resp, err := getDepositAddress(&GetDepositAddressRequest{
		Timestamp: time.Now().UnixMilli(),
		Coin:      "BTC",
		ChainType: "ETH",
	}, &testAuth)
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetDepositOrders(t *testing.T) {
	resp, err := getDepositOrders(&GetDepositOrdersRequest{
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestGetWithdrawOrders(t *testing.T) {
	resp, err := getWithdrawOrders(&GetWithdrawOrdersRequest{
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}

func TestWithdraw(t *testing.T) {
	resp, err := withdraw(&WithdrawRequest{
		Coin:      "USDT",
		ChainType: "ETH",
		Quantity:  1,
		Address:   "0xe44375739b761402242b844d38129d46c72167cf",
	}, &testAuth)
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}
