package main

import "testing"

func TestCreateSignature(t *testing.T) {
	msg := `symbol=ETHBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1538323200000`
	signature := create_signature(msg, "lH3ELTNiFxCQTmi9pPcWWikhsjO04Yoqw3euoHUuOLC3GYBW64ZqzQsiOEHXQS76")
	t.Log(signature)
}
