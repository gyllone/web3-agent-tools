package main

// get Account balance
func getAccountInfo(req *GetAccountInfoRequest, auth *HashKeyApiAuth) (*GetAccountInfoResponse, error) {
	resp := &GetAccountInfoResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(&resp), API_ACCOUNT, "GET", auth)
}

// get account trade list
func getAccountTradeList(req *GetAccountTradeListRequest, auth *HashKeyApiAuth) (GetAccountTradeListResponse, error) {
	resp := GetAccountTradeListResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(&resp), API_ACCOUNT_TRADE_LIST, "GET", auth)
}

// get account balance flow
// TODO: not work
func getAccountBalanceFlow(req *GetAccountBalanceFlowRequest, auth *HashKeyApiAuth) (GetAccountBalanceFlowResponse, error) {
	var resp GetAccountBalanceFlowResponse
	return resp, requestWithSignature(req, getDeserializeJsonFunc(&resp), API_ACCOUNT_BALANCE_FLOW, "GET", auth)
}
