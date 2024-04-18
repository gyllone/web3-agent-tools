package main

/*
#cgo CFLAGS: -I../../dependencies
#include <global_metrics.h>
*/
import "C"
import (
	"coinmarketcap/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"unsafe"
)

// TODO: 返回err_Optional_Metric 报panic
//
//export query_quotes_latest
func query_quotes_latest(convert, convert_id C.Optional_String) C.Result_Optional_Metric {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()
	fmt.Println("------go print start------")
	u, err := url.Parse(QuotesLatestUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		fmt.Println("go print", errStr)
		return C.err_Optional_Metric(C.CString(errStr))
	}

	params := url.Values{}
	if bool(convert.is_some) {
		params.Add("convert", C.GoString(convert.value))
	}
	if bool(convert_id.is_some) {
		params.Add("convert_id", C.GoString(convert_id.value))
	}
	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		fmt.Println("go print", errStr)
		return C.err_Optional_Metric(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	defer response.Body.Close()

	if err2 != nil {
		errStr := "Failed to send request"
		fmt.Println("go print", errStr)
		return C.err_Optional_Metric(C.CString(errStr))
	}

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	defer respBodyDecomp.Close()

	if err3 != nil {
		errStr := "Failed to decompress response"
		fmt.Println("go print", errStr)
		return C.err_Optional_Metric(C.CString(errStr))
	}

	var respBody QuotesLatestResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		fmt.Println("go print", errStr)
		return C.err_Optional_Metric(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		fmt.Println("go print", respBody.Status.ErrorMessage)
		return C.err_Optional_Metric(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.some_Metric(C.Metric{
		active_cryptocurrencies:             C.Int(respData.ActiveCryptocurrencies),
		total_cryptocurrencies:              C.Int(respData.TotalCryptocurrencies),
		active_market_pairs:                 C.Int(respData.ActiveMarketPairs),
		active_exchanges:                    C.Int(respData.ActiveExchanges),
		total_exchanges:                     C.Int(respData.TotalExchanges),
		eth_dominance:                       C.Float(respData.EthDominance),
		btc_dominance:                       C.Float(respData.BtcDominance),
		eth_dominance_yesterday:             C.Float(respData.EthDominanceYesterday),
		btc_dominance_yesterday:             C.Float(respData.BtcDominanceYesterday),
		eth_dominance_24h_percentage_change: C.Float(respData.EthDominance24hPercentageChange),
		btc_dominance_24h_percentage_change: C.Float(respData.BtcDominance24hPercentageChange),
		defi_volume_24h:                     C.Float(respData.DefiVolume24h),
		defi_volume_24h_reported:            C.Float(respData.DefiVolume24hReported),
		defi_market_cap:                     C.Float(respData.DefiMarketCap),
		defi_24h_percentage_change:          C.Float(respData.Defi24hPercentageChange),
		stablecoin_volume_24h:               C.Float(respData.StablecoinVolume24h),
		stablecoin_volume_24h_reported:      C.Float(respData.StablecoinVolume24hReported),
		stablecoin_market_cap:               C.Float(respData.StablecoinMarketCap),
		stablecoin_24h_percentage_change:    C.Float(respData.Stablecoin24hPercentageChange),
		derivatives_volume_24h:              C.Float(respData.DerivativesVolume24h),
		derivatives_volume_24h_reported:     C.Float(respData.DerivativesVolume24hReported),
		derivatives_24h_percentage_change:   C.Float(respData.Derivatives24hPercentageChange),
		quote:                               getCQuoteDict(respData.Quote),
		last_updated:                        C.CString(respData.LastUpdated.String()),
	})

	return C.ok_Optional_Metric(data)
}

func getCQuoteDict(quotes map[string]Quote) C.Dict_Quote {
	res := C.new_Dict_Quote(C.size_t(len(quotes)))

	if res.len == 0 {
		return res
	}
	quoteKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(res.keys))[:res.len:res.len]
	quoteValueArr := (*[1 << 30]C.Quote)(unsafe.Pointer(res.values))[:res.len:res.len]

	quoteIdx := 0

	for key, value := range quotes {
		quoteKeyArr[quoteIdx] = C.CString(key)
		quoteValueArr[quoteIdx] = C.Quote{
			total_market_cap:                             C.Float(value.TotalMarketCap),
			total_volume_24h:                             C.Float(value.TotalVolume24h),
			total_volume_24h_reported:                    C.Float(value.TotalVolume24hReported),
			altcoin_volume_24h:                           C.Float(value.AltcoinVolume24h),
			altcoin_volume_24h_reported:                  C.Float(value.AltcoinVolume24hReported),
			altcoin_market_cap:                           C.Float(value.AltcoinMarketCap),
			defi_volume_24h:                              C.Float(value.DefiVolume24h),
			defi_volume_24h_reported:                     C.Float(value.DefiVolume24hReported),
			defi_24h_percentage_change:                   C.Float(value.Defi24hPercentageChange),
			defi_market_cap:                              C.Float(value.DefiMarketCap),
			stablecoin_volume_24h:                        C.Float(value.StablecoinVolume24h),
			stablecoin_volume_24h_reported:               C.Float(value.StablecoinVolume24hReported),
			stablecoin_24h_percentage_change:             C.Float(value.Stablecoin24hPercentageChange),
			stablecoin_market_cap:                        C.Float(value.StablecoinMarketCap),
			derivatives_volume_24h:                       C.Float(value.DerivativesVolume24h),
			derivatives_volume_24h_reported:              C.Float(value.DerivativesVolume24hReported),
			derivatives_24h_percentage_change:            C.Float(value.Derivatives24hPercentageChange),
			total_market_cap_yesterday:                   C.Float(value.TotalMarketCapYesterday),
			total_volume_24h_yesterday:                   C.Float(value.TotalVolume24hYesterday),
			total_market_cap_yesterday_percentage_change: C.Float(value.TotalMarketCapYesterdayPercentageChange),
			total_volume_24h_yesterday_percentage_change: C.Float(value.TotalVolume24hYesterdayPercentageChange),
			last_updated:                                 C.CString(value.LastUpdated.String()),
		}
		quoteIdx++
	}

	return res
}

//export query_quotes_latest_release
func query_quotes_latest_release(result C.Result_Optional_Metric) {
	fmt.Println("123")
	C.release_Result_Optional_Metric(result)
}

func main() {}
