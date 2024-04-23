package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	query "github.com/google/go-querystring/query"
)

// create a spot order
func createSpotOrder(req *CreateSpotOrderRequest, auth *HashKeyApiAuth) (*CreateSpotOrderResponse, error) {
	resp := &CreateSpotOrderResponse{}
	return resp, requestWithSignature(req, getDeserializeJsonFunc(resp), API_SPOT_ORDER, "POST", auth)
}

func createMultiSpotOrder(req *CreateMultiSpotOrderRequest, auth *HashKeyApiAuth) (*CreateMultiSpotOrderResponse, error) {
	orderIdList := make([]string, 0, 0)
	dispatchedReqs := dispatchMultiSpotOrder(req)
	multiResp := make([]*CreateMultiSpotOrderResponse, len(dispatchedReqs), len(dispatchedReqs))
	mergedResponse := &CreateMultiSpotOrderResponse{
		Code:    0,
		Results: make([]*SpotOrderResult, 0, len(dispatchedReqs)),
	}

	var oneErr error = nil

	var wg sync.WaitGroup
	for i, _ := range dispatchedReqs {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			if resp, err := doCreateMultiSpotOrder(dispatchedReqs[index], auth); err != nil {
				oneErr = err
			} else {
				multiResp[index] = resp
			}
		}(i)
	}
	wg.Wait()

	fmt.Printf("multiResp %+v\n", multiResp)

	if oneErr != nil {
		for _, resp := range multiResp {
			if resp != nil {
				for _, order := range resp.Results {
					orderIdList = append(orderIdList, order.SpotOrder.OrderId)
				}
			}
		}
		// cancel others order
		cancelMultiSpotOrders(&CancelMultiOrdersRequest{
			Timestamp: time.Now().UnixMilli(),
			IDS:       strings.Join(orderIdList, ","),
		}, auth)

		return nil, oneErr
	}

	for _, resp := range multiResp {
		mergedResponse.Results = append(mergedResponse.Results, resp.Results...)
	}

	return mergedResponse, nil
}

func dispatchMultiSpotOrder(req *CreateMultiSpotOrderRequest) []*CreateMultiSpotOrderRequest {
	groupedReq := make(map[string][]*BatchSpotOrderItem)
	for _, ele := range req.Orders {
		if groupedReq[ele.Symbol] == nil {
			groupedReq[ele.Symbol] = make([]*BatchSpotOrderItem, 0, 1)
		}
		groupedReq[ele.Symbol] = append(groupedReq[ele.Symbol], ele)
	}

	resp := make([]*CreateMultiSpotOrderRequest, 0, len(groupedReq))
	for _, group := range groupedReq {
		resp = append(resp, &CreateMultiSpotOrderRequest{
			Timestamp: req.Timestamp,
			Orders:    group,
		})
	}
	return resp
}

// create multi-spot order in same symbol
func doCreateMultiSpotOrder(req *CreateMultiSpotOrderRequest, auth *HashKeyApiAuth) (*CreateMultiSpotOrderResponse, error) {
	urlValues, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	urlString := urlValues.Encode()

	bodyJson, _ := json.Marshal(req.Orders)

	url := fmt.Sprintf("%s?%s&signature=%s", API_BATCH_SPOT_ORDER, urlString, create_signature(urlString, auth.Secret))

	fmt.Printf("req: url: %s\nbody: %s\n", url, bodyJson)

	httpReq, _ := http.NewRequest("POST", url, strings.NewReader(string(bodyJson)))

	httpReq.Header.Add("accept", "application/json")
	httpReq.Header.Add("content-type", "application/json")
	httpReq.Header.Add("X-HK-APIKEY", auth.ApiKey)

	res, _ := http.DefaultClient.Do(httpReq)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Printf("get body:%s\n", body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	resp := &CreateMultiSpotOrderResponse{}
	return resp, getDeserializeJsonFunc(resp)(body)
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
