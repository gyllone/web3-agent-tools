package main

import (
	"encoding/json"
	"testing"
)

func TestParseKline(t *testing.T) {
	klines := `[[1712534400000,"69341.86","72701.79","69130.38","71628.59","6.99969",0,"500999.8100448",724,"0","0"],[1712620800000,"71638.64","71731.93","68320.92","69159.21","12.98089",0,"905745.4832861",898,"0","0"],[1712707200000,"69126.61","71064","67611.38","70597.99","16.9925",0,"1169212.517422",1256,"0","0"],[1712793600000,"70563.2","71195.19","70318.3","71004.99","13.0298",0,"919737.5073175",718,"0","0"]]`

	res, err := parseKline([]byte(klines))
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(res)
	t.Logf("%+v\n", string(bytes))
}

func TestRequestKline(t *testing.T) {
	resp, err := requestKline(&QuoteKlineRequest{
		Symbol:   "ETHUSDT",
		Interval: QuoteKlineInterval_1d,
	})
	if err != nil {
		t.Error(err)
	}
	bytes, _ := json.Marshal(resp)
	t.Logf("%+v\n", string(bytes))
}
