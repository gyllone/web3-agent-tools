package main

func getDepositAddress(req *GetDepositAddressRequest, auth *HashKeyApiAuth) (*GetDepositAddressResponse, error) {
	resp := GetDepositAddressResponse{}
	return &resp, requestWithSignature(req, getDeserializeJsonFunc(&resp), API_ACCOUNT_DEPOSIT_ADDRESS, "GET", auth)
}

func getDepositOrders(req *GetDepositOrdersRequest, auth *HashKeyApiAuth) (*GetDepositOrdersResponse, error) {
	resp := &GetDepositOrdersResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(&resp), API_ACCOUNT_DEPOSIT_ORDERS, "GET", auth)
}

// not implemented
func authDepositAddress() {

}

func getWithdrawOrders(req *GetWithdrawOrdersRequest, auth *HashKeyApiAuth) (*GetWithdrawOrdersResponse, error) {
	resp := GetWithdrawOrdersResponse{}
	return &resp, requestWithSignature(req, getDeserializeJsonFunc(&resp), API_ACCOUNT_WITHDRAW_ORDERS, "GET", auth)

}

func withdraw(req *WithdrawRequest, auth *HashKeyApiAuth) (*WithdrawResponse, error) {
	resp := WithdrawResponse{}
	return &resp, requestWithSignature(req, getDeserializeJsonFunc(&resp), API_ACCOUNT_WITHDRAW, "POST", auth)
}
