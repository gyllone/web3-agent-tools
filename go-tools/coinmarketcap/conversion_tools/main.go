package main

/*
#cgo CFLAGS: -I../../dependencies
#include <conversion_tools.h>
*/
import "C"
import (
	"coinmarketcap/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"
	"unsafe"
)

//export query_price_conversion
func query_price_conversion(amount C.Float, id, symbol, time, convert, convert_id C.Optional_String) C.Result_Optional_PriceConversion {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()

	idIsSome := bool(id.is_some)
	symbolIsSome := bool(symbol.is_some)
	if !(idIsSome || symbolIsSome) {
		errStr := "id or symbol must have at least one"
		return C.err_Optional_PriceConversion(C.CString(errStr))
	}

	u, err := url.Parse(PriceConversionUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_Optional_PriceConversion(C.CString(errStr))
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
		return C.err_Optional_PriceConversion(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_Optional_PriceConversion(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_Optional_PriceConversion(C.CString(errStr))
	}

	var respBody PriceConversionResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response" + err.Error()
		return C.err_Optional_PriceConversion(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_Optional_PriceConversion(C.CString(respBody.Status.ErrorMessage))
	}

	defer respBodyDecomp.Close()

	respData := respBody.Data
	data := C.some_PriceConversion(C.PriceConversion{
		id:           C.Int(respData.ID),
		symbol:       C.CString(respData.Symbol),
		name:         C.CString(respData.Name),
		amount:       C.Float(respData.Amount),
		last_updated: C.CString(respData.LastUpdated.String()),
		quote:        C.new_Dict_Quote(C.size_t(len(respData.Quote))),
	})

	quoteKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(data.value.quote.keys))[:data.value.quote.len:data.value.quote.len]
	quoteValueArr := (*[1 << 30]C.Quote)(unsafe.Pointer(data.value.quote.values))[:data.value.quote.len:data.value.quote.len]

	quoteIdx := 0

	for k, v := range respData.Quote {
		quoteKeyArr[quoteIdx] = C.CString(k)
		quoteValueArr[quoteIdx] = C.Quote{
			price:        C.Float(v.Price),
			last_updated: C.CString(v.LastUpdated.String()),
		}
		quoteIdx++
	}

	return C.ok_Optional_PriceConversion(data)
}

//export query_price_conversion_release
func query_price_conversion_release(result C.Result_Optional_PriceConversion) {
	C.release_Result_Optional_PriceConversion(result)
}

func main() {}
