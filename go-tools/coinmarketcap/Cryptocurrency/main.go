package main

/*
#include <stdlib.h>
#include <stdbool.h>

typedef struct {
	bool is_some;
	char* value;
} OptionalStr;

typedef struct {
	bool is_some;
	long long value;
} OptionalInt;

typedef struct {
	bool is_some;
	bool value;
} OptionalBool;

typedef struct {
	size_t len;
	double* currencies;
} PriceArr;

typedef struct {
	size_t len;
	PriceArr* prices;
} QuoteArr;

typedef struct {
	bool is_fail;
	char* error_message;
	QuoteArr quotes;
} QuoteResult;

typedef struct {
	size_t len;
	char** keys;
	char** values;
} IdMap;

typedef struct {
	size_t len;
	IdMap* id_maps;
} IdMapArr;

typedef struct {
	bool is_fail;
	char* error_message;
	IdMapArr id_maps;
} IdMapResult;

typedef struct {
	size_t len;
	char** keys;
	char** values;
} Cryptocurrency;

typedef struct {
	size_t len;
	Cryptocurrency* cryptocurrency;
} CryptocurrencyArr;

typedef struct {
	bool is_fail;
	char* error_message;
	CryptocurrencyArr cryptocurrencies;
} ListingsLatestResult
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
// func query_quotes(ids C.OptionalStr, slug C.OptionalStr, symbol C.OptionalStr, convert C.OptionalStr, convert_id C.OptionalStr, aux C.OptionalStr, skip_invalid C.OptionalBool) C.QuoteResult {
//
//export query_quotes
func query_quotes(ids C.OptionalStr, slug C.OptionalStr, convert C.OptionalStr, convert_id C.OptionalStr, aux C.OptionalStr, skip_invalid C.OptionalBool) C.QuoteResult {
	idsIsSome := bool(ids.is_some)
	slugIsSome := bool(slug.is_some)
	if !(idsIsSome || slugIsSome) {
		errStr := "ids or slug must have at least one"
		fmt.Println("go print: ", errStr)
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
		}
	}
	convertIsSome := bool(convert.is_some)
	convertIdIsSome := bool(convert_id.is_some)
	auxIsSome := bool(aux.is_some)
	skipInvalidIsSome := bool(skip_invalid.is_some)

	u, err := url.Parse(Quotes)
	if err != nil {
		errStr := "Failed to parse URL"
		fmt.Println("go print: ", errStr)
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
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
		fmt.Println("go print: ", errStr)
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
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
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
		}
	}

	defer response.Body.Close()
	var respBody QuoteResp

	err = json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response"
		fmt.Println("go print: ", errStr)
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
		}
	}

	if response.StatusCode != 200 {
		fmt.Println("go print: ", respBody.Status.ErrorMessage)
		return C.QuoteResult{
			is_fail:       C.bool(true),
			error_message: C.CString(respBody.Status.ErrorMessage),
		}
	}

	quotesResult := C.QuoteResult{
		is_fail:       C.bool(false),
		error_message: C.CString(""),
	}
	quotesResult.quotes = C.QuoteArr{}
	quotesResult.quotes.len = C.size_t(quoteLen)
	quotesResult.quotes.prices = (*C.PriceArr)(C.malloc(C.size_t(quoteLen) * C.sizeof_PriceArr))
	priceArrPtr := (*[1 << 30]C.PriceArr)(unsafe.Pointer(quotesResult.quotes.prices))[:quoteLen:quoteLen]
	for i := 0; i < quoteLen; i++ {
		prices := C.PriceArr{
			len:        C.size_t(convertLen),
			currencies: (*C.double)(C.malloc(C.sizeof_double * C.size_t(convertLen))),
		}
		priceArrPtr[i] = prices
	}

	quoteKeys := reflect.ValueOf(respBody.Data).MapKeys()
	for priceIdx, quote := range quoteKeys {
		currencyKeys := reflect.ValueOf(respBody.Data[quote.String()].Quote).MapKeys()
		currencyArrPtr := (*[1 << 30]C.double)(unsafe.Pointer(priceArrPtr[priceIdx].currencies))[:convertLen:convertLen]

		for currencyIdx, currency := range currencyKeys {
			currencyArrPtr[currencyIdx] = C.double(respBody.Data[quote.String()].Quote[currency.String()].Price)
		}
	}

	return quotesResult
}

//export query_quotes_release
func query_quotes_release(result C.QuoteResult) {
	fmt.Println("go print: execute query_quotes_release")
	C.free(unsafe.Pointer(result.error_message))
	if !bool(result.is_fail) {
		priceArrPtr := (*[1 << 30]C.PriceArr)(unsafe.Pointer(result.quotes.prices))[:result.quotes.len:result.quotes.len]
		for i := 0; i < int(result.quotes.len); i++ {
			C.free(unsafe.Pointer(priceArrPtr[i].currencies))
		}
		C.free(unsafe.Pointer(result.quotes.prices))
	}
}

//export query_id_map
func query_id_map(listing_status C.OptionalStr, start C.OptionalInt, limit C.OptionalInt, sort C.OptionalStr, symbol C.OptionalStr, aux C.OptionalStr) C.IdMapResult {
	u, err := url.Parse(IdMap)
	if err != nil {
		errStr := "Failed to parse URL"
		fmt.Println("go print: ", errStr)
		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
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
		fmt.Println("go print: ", errStr)

		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
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
		}
	}

	defer response.Body.Close()
	var respBody IdMapResp

	err = json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response"
		fmt.Println("go print: ", errStr)

		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(errStr),
		}
	}
	//fmt.Println(respBody)

	if response.StatusCode != 200 {
		fmt.Println("go print: ", "response.StatusCode != 200")

		return C.IdMapResult{
			is_fail:       C.bool(true),
			error_message: C.CString(respBody.Status.ErrorMessage),
		}
	}

	data := respBody.Data
	dataLen := len(data)

	idMapResult := C.IdMapResult{
		is_fail:       C.bool(false),
		error_message: C.CString(""),
		id_maps: C.IdMapArr{
			len: C.size_t(dataLen),
		},
	}
	idMapResult.id_maps.id_maps = (*C.IdMap)(C.malloc(C.sizeof_IdMap * C.size_t(dataLen)))
	idMapsArr := (*[1 << 30]C.IdMap)(unsafe.Pointer(idMapResult.id_maps.id_maps))[:dataLen:dataLen]

	for idx, v := range data {
		fmt.Println("go print: ", v.ID, v.Name, v.Symbol, v.Slug)
		keys := []string{"id", "name", "symbol", "slug"}
		values := []string{strconv.Itoa(v.ID), v.Name, v.Symbol, v.Slug}
		idMap := C.IdMap{
			len: C.size_t(len(keys)),
		}
		idMap.keys = (**C.char)(C.malloc(C.size_t(idMap.len) * C.size_t(unsafe.Sizeof((*C.char)(nil)))))
		idMap.values = (**C.char)(C.malloc(C.size_t(idMap.len) * C.size_t(unsafe.Sizeof((*C.char)(nil)))))

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
	fmt.Println("go print: execute query_id_map_release")
	C.free(unsafe.Pointer(result.error_message))
	if !bool(result.is_fail) {
		IdMapArrPtr := (*[1 << 30]C.IdMap)(unsafe.Pointer(result.id_maps.id_maps))[:result.id_maps.len:result.id_maps.len]
		for i := 0; i < int(result.id_maps.len); i++ {
			keysArr := (*[1 << 30]*C.char)(unsafe.Pointer(IdMapArrPtr[i].keys))[:IdMapArrPtr[i].len:IdMapArrPtr[i].len]
			valuesArr := (*[1 << 30]*C.char)(unsafe.Pointer(IdMapArrPtr[i].values))[:IdMapArrPtr[i].len:IdMapArrPtr[i].len]
			for j := 0; j < int(IdMapArrPtr[i].len); j++ {
				C.free(unsafe.Pointer(keysArr[j]))
				C.free(unsafe.Pointer(valuesArr[j]))
			}
			C.free(unsafe.Pointer(IdMapArrPtr[i].keys))
			C.free(unsafe.Pointer(IdMapArrPtr[i].values))
		}
		C.free(unsafe.Pointer(result.id_maps.id_maps))
	}
}

func query_listings_latest(start C.OptionalInt, limit C.OptionalInt, price_min C.OptionalStr, price_max C.OptionalStr,
	market_cap_min C.OptionalStr, market_cap_max C.OptionalStr,
	volume_24h_min C.OptionalStr, volume_24h_max C.OptionalStr, circulating_supply_min C.OptionalStr,
	circulating_supply_max C.OptionalStr, percent_change_24h_min C.OptionalInt, percent_change_24h_max C.OptionalInt, convert C.OptionalStr, convert_id C.OptionalStr, sort C.OprionalStr, sort_dir C.OptionalStr, cryptocurrency_type C.OptionalStr, tag C.OptionalStr, aux C.OptionalStr) C.ListingsLatestResult {
	return nil
	//u, err := url.Parse(ListingsLatest)
	//if err != nil {
	//	errStr := "Failed to parse URL"
	//	fmt.Println("go print: ", errStr)
	//	return C.IdMapResult{
	//		is_fail:       C.bool(true),
	//		error_message: C.CString(errStr),
	//	}
	//}
	//// 构建查询参数
	//listingStatusIsSome := bool(listing_status.is_some)
	//startIsSome := bool(start.is_some)
	//limitIsSome := bool(limit.is_some)
	//sortIsSome := bool(sort.is_some)
	//symbolIsSome := bool(symbol.is_some)
	//auxIsSome := bool(aux.is_some)
	//params := url.Values{}
	//
	//if listingStatusIsSome {
	//	params.Add("listing_status", C.GoString(listing_status.value))
	//}
	//
	//if startIsSome {
	//	params.Add("start", strconv.Itoa(int(start.value)))
	//}
	//if limitIsSome {
	//	params.Add("limit", strconv.Itoa(int(limit.value)))
	//}
	//if sortIsSome {
	//	params.Add("sort", C.GoString(sort.value))
	//}
	//if symbolIsSome {
	//	params.Add("symbol", C.GoString(symbol.value))
	//}
	//if auxIsSome {
	//	params.Add("aux", C.GoString(aux.value))
	//}
	//// 将查询参数添加到 URL 查询字符串中
	//u.RawQuery = params.Encode()
	//
	//req, err1 := http.NewRequest("GET", u.String(), nil)
	//if err1 != nil {
	//	errStr := "Failed to create request"
	//	fmt.Println("go print: ", errStr)
	//
	//	return
	//}
	//
	//req.Header.Set("X-CMC_PRO_API_KEY", ApiKey)
	//req.Header.Set("Accept", "application/json")
	////req.Header.Set("Accept-Encoding", "deflate, gzip")
	//
	//client := &http.Client{}
	//response, err2 := client.Do(req)
	//if err2 != nil {
	//	errStr := "Failed to send request"
	//	fmt.Println("go print: ", errStr)
	//
	//	return
	//}
	//
	//defer response.Body.Close()
	//var respBody ListingsLatestResp
	//
	//err = json.NewDecoder(response.Body).Decode(&respBody)
	//if err != nil {
	//	errStr := "Failed to decode response"
	//	fmt.Println("go print: ", errStr)
	//
	//	return
	//}
	////fmt.Println(respBody)
	//
	//if response.StatusCode != 200 {
	//	fmt.Println("go print: ", "response.StatusCode != 200")
	//
	//	return
	//}
	//
	//fmt.Println(respBody)
}

func main() {
	//query_listings_latest()
}
