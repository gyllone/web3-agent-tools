package main

import (
	"testing"
	"time"
)

func TestCreateSpotOrder(t *testing.T) {
	resp, err := createSpotOrder(&SpotOrderRequest{
		Symbol:    "BTCUSDT",
		Side:      "BUY",
		Type:      SpotOrderTypeEnum_LIMIT,
		Quantity:  0.0001,
		Price:     399,
		Timestamp: time.Now().UnixMilli(),
	}, &HashKeyAuth{
		Secret: "4FcFv4f3kIsfpVh2HcL7rES9iLN3nfTubeZvzjb97jXnWpndvvfFLgHIAqrNEJip",
		ApiKey: "AAyrJfY2VBLYCIBNa8k8kbppXAeBzF1i9Tmz2DawdEbWLzzBjGj0bRKFW0q0F3cd",
	})
	if err != nil {
		t.Log(err)
	}
	t.Logf("resp:%+v\n", resp)
}
