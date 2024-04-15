package main

func createSpotOrder(req *CreateSpotOrderRequest, auth *HashKeyApiAuth) (*CreateSpotOrderResponse, error) {

	resp := &CreateSpotOrderResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "POST", auth)
}

func cancelSpotOrder(req *CancelOrderRequest, auth *HashKeyApiAuth) (*CancelOrderResponse, error) {
	resp := &CancelOrderResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "DELETE", auth)
}

func queryCurrentOpenSoptOrder(req *QueryOpenOrdersRequest, auth *HashKeyApiAuth) (QueryOpenOrdersResponse, error) {
	var arrays QueryOpenOrdersResponse
	return arrays, requestWithSignature(req, getDeserializeJsonFunc(&arrays), API_SPOT_OPENORDERS, "GET", auth)
}

func cancelAllSoptOrders(req *CancelAllOpenOrdersRequest, auth *HashKeyApiAuth) (*CancelAllOpenOrdersReponse, error) {
	resp := &CancelAllOpenOrdersReponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_OPENORDERS, "DELETE", auth)
}

func cancelMultiSpotOrders(req *CancelMultiOrdersRequest, auth *HashKeyApiAuth) (*CancelMultiOrdersResponse, error) {
	resp := &CancelMultiOrdersResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_CANCEL_ORDER_BY_IDS, "DELETE", auth)
}
