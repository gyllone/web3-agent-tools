package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

func main() {
	klinereq := QuoteKlineRequest{}
	fmt.Println(query.Values(klinereq))
	url := fmt.Sprintf("%s?symbol=BTCUSDT&interval=1m&limit=100", API_KLINE)
	fmt.Println("query:", url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
