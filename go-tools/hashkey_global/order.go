package main

func createSpotOrder(req *CreateSpotOrderRequest, auth *HashKeyAuth) (*CreateSpotOrderResponse, error) {

	resp := &CreateSpotOrderResponse{}
	return resp, request(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "POST", auth)
}

func cancelSpotOrder(req *CancelOrderRequest, auth *HashKeyAuth) (*CancelOrderResponse, error) {
	resp := &CancelOrderResponse{}
	return resp, request(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "DELETE", auth)
}

func queryCurrentSoptOrder(req *QueryOpenOrdersRequest, auth *HashKeyAuth) (QueryOpenOrdersResponse, error) {
	var arrays QueryOpenOrdersResponse
	return arrays, request(req, getDeserializeJsonFunc(&arrays), API_SPOT_OPENORDERS, "GET", auth)
}

func cancelAllSoptOrders(req *CancelAllOpenOrdersRequest, auth *HashKeyAuth) (*CancelAllOpenOrdersReponse, error) {
	resp := &CancelAllOpenOrdersReponse{}
	return resp, request(req, getDeserializeJsonFunc(resp), API_SPOT_OPENORDERS, "DELETE", auth)
}

func cancelMultiSpotOrders(req *CancelMultiOrdersRequest, auth *HashKeyAuth) (*CancelMultiOrdersResponse, error) {
	resp := &CancelMultiOrdersResponse{}
	return resp, request(req, getDeserializeJsonFunc(resp), API_SPOT_CANCEL_ORDER_BY_IDS, "DELETE", auth)
}
