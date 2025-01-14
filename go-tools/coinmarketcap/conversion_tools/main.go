package main

/*
#cgo CFLAGS: -I../../dependencies
#include <conversion_tools.h>
*/
import "C"
import (
	"coinmarketcap/utils"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"unsafe"
)

//export query_price_conversion
func query_price_conversion(amount C.Float, id, symbol, time, convert, convert_id C.Optional_String) C.Result_PriceConversion {
	idIsSome := bool(id.is_some)
	symbolIsSome := bool(symbol.is_some)
	if !(idIsSome || symbolIsSome) {
		errStr := "id or symbol must have at least one"
		return C.err_PriceConversion(C.CString(errStr))
	}

	u, err := url.Parse(PriceConversionUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_PriceConversion(C.CString(errStr))
	}

	params := url.Values{}
	params.Add("amount", strconv.FormatFloat(float64(amount), 'f', -1, 64))
	if idIsSome {
		params.Add("id", C.GoString(id.value))
	}
	if symbolIsSome {
		params.Add("symbol", C.GoString(symbol.value))
	}
	if bool(time.is_some) {
		params.Add("time", C.GoString(time.value))
	}
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
		return C.err_PriceConversion(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_PriceConversion(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_PriceConversion(C.CString(errStr))
	}

	var respBody PriceConversionResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response" + err.Error()
		return C.err_PriceConversion(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_PriceConversion(C.CString(respBody.Status.ErrorMessage))
	}

	defer respBodyDecomp.Close()

	respData := respBody.Data
	data := C.PriceConversion{
		id:           C.Int(respData.ID),
		symbol:       C.CString(respData.Symbol),
		name:         C.CString(respData.Name),
		amount:       C.Float(respData.Amount),
		last_updated: C.CString(respData.LastUpdated.String()),
		quote:        getCQuoteDict(respData.Quote),
	}

	return C.ok_PriceConversion(data)
}

func getCQuoteDict(quote map[string]Quote) C.Dict_Quote {
	res := C.new_Dict_Quote(C.size_t(len(quote)))

	if res.len == 0 {
		return res
	}
	quoteKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(res.keys))[:res.len:res.len]
	quoteValueArr := (*[1 << 30]C.Quote)(unsafe.Pointer(res.values))[:res.len:res.len]

	quoteIdx := 0

	for k, v := range quote {
		quoteKeyArr[quoteIdx] = C.CString(k)
		quoteValueArr[quoteIdx] = C.Quote{
			price:        C.Float(v.Price),
			last_updated: C.CString(v.LastUpdated.String()),
		}
		quoteIdx++
	}

	return res
}

//export query_price_conversion_release
func query_price_conversion_release(result C.Result_PriceConversion) {
	C.release_Result_PriceConversion(result)
}

func main() {}
