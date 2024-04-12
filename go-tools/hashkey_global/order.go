package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func createSpotOrder(order *SpotOrderRequest, auth *HashKeyAuth) (*SpotOrderResponse, error) {

	urlValues, err := query.Values(order)
	if err != nil {
		return nil, err
	}

	urlString := urlValues.Encode()
	signature := create_signature(urlString, auth.Secret)
	url := fmt.Sprintf("%s?%s&signature=%s", API_SPOT_ORDER, urlString, signature)
	fmt.Printf("createSpotOrder url: %s\n", url)
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-HK-APIKEY", auth.ApiKey)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	return nil, nil
}
