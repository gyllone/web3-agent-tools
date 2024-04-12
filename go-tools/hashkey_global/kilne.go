package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	query "github.com/google/go-querystring/query"
)

func parseKline(klines []byte) (QuoteKlineResponse, error) {
	var object [][]interface{}
	if err := json.Unmarshal(klines, &object); err != nil {
		return nil, err
	}

	res := make(QuoteKlineResponse, 0, len(object))
	for _, array := range object {
		point := &QuoteKlinePoint{
			OpenTime:                 int64(array[0].(float64)),
			OpenPrice:                array[1].(string),
			HighPrice:                array[2].(string),
			LowPrice:                 array[3].(string),
			ClosePrice:               array[4].(string),
			Volume:                   array[5].(string),
			CloseTime:                int64(array[6].(float64)),
			QuoteAssetVolume:         array[7].(string),
			NumberOfTrades:           int(array[8].(float64)),
			TakerBuyBaseAssetVolume:  array[9].(string),
			TakerBuyQuoteAssetVolume: array[10].(string),
		}
		res = append(res, point)
	}
	return res, nil
}

func requestKline(quoteReq *QuoteKlineRequest) (QuoteKlineResponse, error) {

	urlValues, err := query.Values(quoteReq)
	if err != nil {
		return nil, err
	}
	urlQuery := urlValues.Encode()
	url := fmt.Sprintf("%s?%s", API_KLINE, urlQuery)

	fmt.Println("query:", url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	return parseKline(body)

}
