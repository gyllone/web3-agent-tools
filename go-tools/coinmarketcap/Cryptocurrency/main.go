package main

/*
#cgo CFLAGS: -I../../dependencies
#include <cryptocurrency.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// TODO: 使用symbol参数请求的话，返回格式与id和slug不统一，暂不支持
//
//export query_quotes
func query_quotes(ids C.Optional_String, slug C.Optional_String, convert C.Optional_String, convert_id C.Optional_String, aux C.Optional_String, skip_invalid C.Optional_Bool) C.QuoteResult {
	idsIsSome := bool(ids.is_some)
	slugIsSome := bool(slug.is_some)
	if !(idsIsSome || slugIsSome) {
		errStr := "ids or slug must have at least one"
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			quotes:        C.new_List_List_Float(0),
		}
	}
	convertIsSome := bool(convert.is_some)
	convertIdIsSome := bool(convert_id.is_some)
	auxIsSome := bool(aux.is_some)
	skipInvalidIsSome := bool(skip_invalid.is_some)

	u, err := url.Parse(Quotes)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			quotes:        C.new_List_List_Float(0),
		}
	}
	// 构建查询参数
	params := url.Values{}
	var quoteLen int
	convertLen := 1

	if idsIsSome {
		idStr := C.GoString(ids.value)
		params.Add("id", idStr)
		quoteLen = len(strings.Split(idStr, ","))
	}
	if slugIsSome {
		slugStr := C.GoString(slug.value)
		params.Add("slug", slugStr)
		quoteLen = len(strings.Split(slugStr, ","))
	}
	if convertIsSome {
		convertStr := C.GoString(convert.value)
		params.Add("convert", convertStr)
		convertLen = int(math.Max(float64(convertLen), float64(len(strings.Split(convertStr, ",")))))
	}
	if convertIdIsSome {
		convertIdStr := C.GoString(convert_id.value)
		params.Add("convert_id", convertIdStr)
		convertLen = int(math.Max(float64(convertLen), float64(len(strings.Split(convertIdStr, ",")))))
	}
	if auxIsSome {
		auxStr := C.GoString(aux.value)
		params.Add("aux", auxStr)
	}
	if skipInvalidIsSome {
		skipInvalidStr := strconv.FormatBool(bool(skip_invalid.value))
		params.Add("skip_invalid", skipInvalidStr)
	}

	// 将查询参数添加到 URL 查询字符串中
	u.RawQuery = params.Encode()

	fmt.Println("go print: ", u.String())

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			quotes:        C.new_List_List_Float(0),
		}
	}

	req.Header.Set("X-CMC_PRO_API_KEY", ApiKey)
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			quotes:        C.new_List_List_Float(0),
		}
	}

	defer response.Body.Close()
	var respBody QuoteResp

	err = json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response"
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			quotes:        C.new_List_List_Float(0),
		}
	}

	if response.StatusCode != 200 {
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(respBody.Status.ErrorMessage),
			quotes:        C.new_List_List_Float(0),
		}
	}

	quotesResult := C.QuoteResult{
		is_fail:       C.bool(false),
		error_message: C.CString(""),
	}
	quotesResult.quotes = C.new_List_List_Float(C.size_t(quoteLen))
	priceArrPtr := (*[1 << 30]C.List_Float)(unsafe.Pointer(quotesResult.quotes.values))[:quoteLen:quoteLen]
	for i := 0; i < quoteLen; i++ {
		priceArrPtr[i] = C.new_List_Float(C.size_t(convertLen))
	}

	quoteKeys := reflect.ValueOf(respBody.Data).MapKeys()
	for priceIdx, quote := range quoteKeys {
		currencyKeys := reflect.ValueOf(respBody.Data[quote.String()].Quote).MapKeys()
		currencyArrPtr := (*[1 << 30]C.double)(unsafe.Pointer(priceArrPtr[priceIdx].values))[:convertLen:convertLen]

		for currencyIdx, currency := range currencyKeys {
			currencyArrPtr[currencyIdx] = C.double(respBody.Data[quote.String()].Quote[currency.String()].Price)
		}
	}

	return quotesResult
}

//export query_quotes_release
func query_quotes_release(result C.QuoteResult) {
	C.release_QuoteResult(result)
}

//export query_id_map
func query_id_map(listing_status C.Optional_String, start C.Optional_Int, limit C.Optional_Int, sort C.Optional_String, symbol C.Optional_String, aux C.Optional_String) C.IdMapResult {
	u, err := url.Parse(IdMap)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			id_maps:       C.new_List_Dict_String(0),
		}
	}
	// 构建查询参数
	listingStatusIsSome := bool(listing_status.is_some)
	startIsSome := bool(start.is_some)
	limitIsSome := bool(limit.is_some)
	sortIsSome := bool(sort.is_some)
	symbolIsSome := bool(symbol.is_some)
	auxIsSome := bool(aux.is_some)
	params := url.Values{}
	if listingStatusIsSome {
		params.Add("listing_status", C.GoString(listing_status.value))
	}

	if startIsSome {
		params.Add("start", strconv.Itoa(int(start.value)))
	}
	if limitIsSome {
		params.Add("limit", strconv.Itoa(int(limit.value)))
	}
	if sortIsSome {
		params.Add("sort", C.GoString(sort.value))
	}
	if symbolIsSome {
		params.Add("symbol", C.GoString(symbol.value))
	}
	if auxIsSome {
		params.Add("aux", C.GoString(aux.value))
	}
	// 将查询参数添加到 URL 查询字符串中
	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			id_maps:       C.new_List_Dict_String(0),
		}
	}

	req.Header.Set("X-CMC_PRO_API_KEY", ApiKey)
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		fmt.Println("go print: ", errStr)

		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			id_maps:       C.new_List_Dict_String(0),
		}
	}

	defer response.Body.Close()
	var respBody IdMapResp

	err = json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response"
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			id_maps:       C.new_List_Dict_String(0),
		}
	}

	if response.StatusCode != 200 {
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(respBody.Status.ErrorMessage),
			id_maps:       C.new_List_Dict_String(0),
		}
	}

	data := respBody.Data
	dataLen := len(data)

	idMapResult := C.IdMapResult{
		is_fail:       C.bool(false),
		error_message: C.CString(""),
		id_maps:       C.new_List_Dict_String(C.size_t(dataLen)),
	}
	idMapsArr := (*[1 << 30]C.Dict_String)(unsafe.Pointer(idMapResult.id_maps.values))[:dataLen:dataLen]

	for idx, v := range data {
		fmt.Println("go print: ", v.ID, v.Name, v.Symbol, v.Slug)
		keys := []string{"id", "name", "symbol", "slug"}
		values := []string{strconv.Itoa(v.ID), v.Name, v.Symbol, v.Slug}
		idMap := C.new_Dict_String(C.size_t(len(keys)))
		cKeys := (*[1 << 30]*C.char)(unsafe.Pointer(idMap.keys))[:idMap.len:idMap.len]
		cValues := (*[1 << 30]*C.char)(unsafe.Pointer(idMap.values))[:idMap.len:idMap.len]
		for i := 0; i < len(keys); i++ {
			cKeys[i] = C.CString(keys[i])
			cValues[i] = C.CString(values[i])
		}

		idMapsArr[idx] = idMap
	}
	return idMapResult
}

//export query_id_map_release
func query_id_map_release(result C.IdMapResult) {
	C.release_IdMapResult(result)
}

func main() {}
