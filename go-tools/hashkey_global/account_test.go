package main

import (
	"testing"
	"time"
)

func TestGetAccountInfo(t *testing.T) {
	resp, err := getAccountInfo(&GetAccountInfoRequest{
		Timestamp: time.Now().UnixMilli(),
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestGetAccountTradeList(t *testing.T) {
	resp, err := getAccountTradeList(&GetAccountTradeListRequest{
		Timestamp: time.Now().UnixMilli(),
		// EndTime:   1713148313000,
		// StartTime: 1712802713000,
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}

func TestGetAccountBalanceFlow(t *testing.T) {
	resp, err := getAccountBalanceFlow(&GetAccountBalanceFlowRequest{
		Timestamp: time.Now().UnixMilli(),
		EndTime:   1713182770000,
		StartTime: 1712802713000,
		FlowType:  BalanceFlowType_Trade,
		Limit:     10,
	}, &testAuth)
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}
