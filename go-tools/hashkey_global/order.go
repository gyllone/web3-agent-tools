package main

// create a spot order
func createSpotOrder(req *CreateSpotOrderRequest, auth *HashKeyApiAuth) (*CreateSpotOrderResponse, error) {
	resp := &CreateSpotOrderResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "POST", auth)
}

// query a spot order by order id
func querySpotOrder(req *QueryOrderRequest, auth *HashKeyApiAuth) (*QueryOrderResponse, error) {
	resp := &QueryOrderResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "GET", auth)
}

// cancel a spot order by order id
func cancelSpotOrder(req *CancelOrderRequest, auth *HashKeyApiAuth) (*CancelOrderResponse, error) {
	resp := &CancelOrderResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "DELETE", auth)
}

// query all open spot orders
func queryCurrentOpenSoptOrder(req *QueryOpenOrdersRequest, auth *HashKeyApiAuth) (QueryOpenOrdersResponse, error) {
	var arrays QueryOpenOrdersResponse
	return arrays, requestWithSignature(req, getDeserializeJsonFunc(&arrays), API_SPOT_OPENORDERS, "GET", auth)
}

// cancel all spot orders
func cancelAllSoptOrders(req *CancelAllOpenOrdersRequest, auth *HashKeyApiAuth) (*CancelAllOpenOrdersReponse, error) {
	resp := &CancelAllOpenOrdersReponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_OPENORDERS, "DELETE", auth)
}

// cancel multi spot orders
func cancelMultiSpotOrders(req *CancelMultiOrdersRequest, auth *HashKeyApiAuth) (*CancelMultiOrdersResponse, error) {
	resp := &CancelMultiOrdersResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_CANCEL_ORDER_BY_IDS, "DELETE", auth)
}

// query traded orders
func queryAllTradedOrders(req *QueryAllTradedOrdersRequest, auth *HashKeyApiAuth) (*QueryAllTradedOrdersResponse, error) {
	resp := &QueryAllTradedOrdersResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_TRADED_ORDERS, "GET", auth)
}
