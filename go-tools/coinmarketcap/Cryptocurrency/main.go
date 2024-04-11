package main

/*
#cgo CFLAGS: -I../../dependencies
#include <cryptocurrency.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"go-tools/dependencies"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"unsafe"
)

// TODO: 使用symbol参数请求的话，返回格式与id和slug不统一，暂不支持
//
//export query_quotes
func query_quotes(ids, slug, convert, convert_id, aux C.Optional_String, skip_invalid C.Optional_Bool) (result C.QuoteResult) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.QuoteResult{
				is_fail:       C.bool(true),
				error_message: C.CString("go panic"),
				quotes:        C.new_List_List_Float(0),
			}
		}
	}()
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

	u, err := url.Parse(QuotesUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			quotes:        C.new_List_List_Float(0),
		}
	}

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
	req.Header.Set("Accept-Encoding", "deflate, gzip")

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

	respBodyDecomp, err3 := dependencies.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			quotes:        C.new_List_List_Float(0),
		}
	}

	var respBody QuoteResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response" + err.Error()
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
func query_id_map(listing_status, sort, symbol, aux C.Optional_String, start, limit C.Optional_Int) (result C.IdMapResult) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.IdMapResult{
				is_fail:       C.bool(true),
				error_message: C.CString("go panic"),
				id_maps:       C.new_List_Dict_String(0),
			}
		}
	}()
	u, err := url.Parse(IdMapUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			id_maps:       C.new_List_Dict_String(0),
		}
	}

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
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			id_maps:       C.new_List_Dict_String(0),
		}
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := dependencies.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			id_maps:       C.new_List_Dict_String(0),
		}
	}

	var respBody IdMapResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
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
		keys := []string{"id", "name", "symbol", "slug"}
		values := []string{strconv.Itoa(v.ID), v.Name, v.Symbol, v.Slug}
		idMap := C.new_Dict_String(C.size_t(len(keys)))
		cKeys := (*[1 << 30]C.String)(unsafe.Pointer(idMap.keys))[:idMap.len:idMap.len]
		cValues := (*[1 << 30]C.String)(unsafe.Pointer(idMap.values))[:idMap.len:idMap.len]
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

// TODO: symbol请求与id和slug格式不符，暂未实现
//
//export query_metadata
func query_metadata(id, slug, address, aux C.Optional_String, skip_invalid C.Optional_Bool) (result C.MetadataResult) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.MetadataResult{
				is_fail:       C.bool(true),
				error_message: C.CString("go panic"),
				metas:         C.new_List_Dict_String(0),
			}
		}
	}()
	idIsSome := bool(id.is_some)
	slugIsSome := bool(slug.is_some)
	addressIsSome := bool(address.is_some)
	auxIsSome := bool(aux.is_some)
	skipInvalidIsSome := bool(skip_invalid.is_some)

	if !(idIsSome || slugIsSome) {
		errStr := "id or slug must have at least one"
		return C.MetadataResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			metas:         C.new_List_Dict_String(0),
		}
	}

	u, err := url.Parse(MetadataUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.MetadataResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			metas:         C.new_List_Dict_String(0),
		}
	}

	params := url.Values{}
	if idIsSome {
		params.Add("id", C.GoString(id.value))
	}
	if slugIsSome {
		params.Add("slug", C.GoString(slug.value))
	}
	if addressIsSome {
		params.Add("address", C.GoString(address.value))
	}
	if skipInvalidIsSome {
		params.Add("skip_invalid", strconv.FormatBool(bool(skip_invalid.value)))
	}
	if auxIsSome {
		params.Add("aux", C.GoString(aux.value))
	}
	// 将查询参数添加到 URL 查询字符串中
	u.RawQuery = params.Encode()
	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.MetadataResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			metas:         C.new_List_Dict_String(0),
		}
	}

	req.Header.Set("X-CMC_PRO_API_KEY", ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.MetadataResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			metas:         C.new_List_Dict_String(0),
		}
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := dependencies.DecompressResponse(response)

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.MetadataResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			metas:         C.new_List_Dict_String(0),
		}
	}

	var respBody MetadataResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.MetadataResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
			metas:         C.new_List_Dict_String(0),
		}
	}

	if response.StatusCode != 200 {
		return C.MetadataResult{
			is_fail:       C.bool(true),
			error_message: C.CString(respBody.Status.ErrorMessage),
			metas:         C.new_List_Dict_String(0),
		}
	}

	result = C.MetadataResult{
		is_fail:       C.bool(false),
		error_message: C.CString(""),
		metas:         C.new_List_Dict_String(C.size_t(len(respBody.Data))),
	}

	metaArr := (*[1 << 30]C.Dict_String)(unsafe.Pointer(result.metas.values))[:result.metas.len:result.metas.len]
	idxMeta := 0
	for _, meta := range respBody.Data {
		keys, values := parseMetadata(meta)

		metaCMap := C.new_Dict_String(C.size_t(len(keys)))
		cKeys := (*[1 << 30]C.String)(unsafe.Pointer(metaCMap.keys))[:metaCMap.len:metaCMap.len]
		cValues := (*[1 << 30]C.String)(unsafe.Pointer(metaCMap.values))[:metaCMap.len:metaCMap.len]

		for i := 0; i < len(keys); i++ {
			cKeys[i] = C.CString(keys[i])
			cValues[i] = C.CString(values[i])
		}
		metaArr[idxMeta] = metaCMap

		idxMeta++
	}

	return result
}

//export query_metadata_release
func query_metadata_release(result C.MetadataResult) {
	C.release_MetadataResult(result)
}

func parseMetadata(meta Metadata) (keys []string, values []string) {
	keys = []string{"id", "name", "symbol", "slug", "category", "description", "tags", "urls"}
	values = []string{strconv.Itoa(meta.ID), meta.Name, meta.Symbol, meta.Slug, meta.Category, meta.Description, "", ""}

	dataLen := len(keys)
	for _, tag := range meta.Tags {
		values[dataLen-2] += tag + ","
	}
	values[dataLen-2] = strings.TrimSuffix(values[dataLen-2], ",")

	urls := reflect.ValueOf(meta.URLs)
	for i := 0; i < urls.Type().NumField(); i++ {
		website := urls.Field(i)
		for j := 0; j < website.Len(); j++ {
			value := website.Index(j).String()
			values[dataLen-1] += value + ","
		}
	}
	values[dataLen-1] = strings.TrimSuffix(values[dataLen-1], ",")

	if meta.SelfReportedCirculatingSupply != nil {
		keys = append(keys, "self_reported_circulating_supply")
		values = append(values, strconv.FormatFloat(*meta.SelfReportedCirculatingSupply, 'f', -1, 64))
	}

	if meta.SelfReportedCirculatingSupply != nil {
		keys = append(keys, "self_reported_market_cap")
		values = append(values, strconv.FormatFloat(*meta.SelfReportedMarketCap, 'f', -1, 64))
	}

	return
}

func main() {}
