package main

/*
#cgo CFLAGS: -I../../dependencies
#include <cryptocurrency.h>
*/
import "C"
import (
	"coinmarketcap/utils"
	"encoding/json"
	"fmt"
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
func query_quotes(id, slug, convert, convert_id, aux C.Optional_String, skip_invalid C.Optional_Bool) (result C.Result_List_List_Float) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.err_List_List_Float(C.CString("go panic"))
		}
	}()
	idsIsSome := bool(id.is_some)
	slugIsSome := bool(slug.is_some)
	if !(idsIsSome || slugIsSome) {
		errStr := "id or slug must have at least one"
		return C.err_List_List_Float(C.CString(errStr))
	}
	convertIsSome := bool(convert.is_some)
	convertIdIsSome := bool(convert_id.is_some)
	auxIsSome := bool(aux.is_some)
	skipInvalidIsSome := bool(skip_invalid.is_some)

	u, err := url.Parse(QuotesUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_List_Float(C.CString(errStr))
	}

	params := url.Values{}
	var quoteLen int
	convertLen := 1

	if idsIsSome {
		idStr := C.GoString(id.value)
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
		return C.err_List_List_Float(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_List_Float(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_List_Float(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody QuoteResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response" + err.Error()
		return C.err_List_List_Float(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_List_Float(C.CString(respBody.Status.ErrorMessage))
	}

	data := C.new_List_List_Float(C.size_t(quoteLen))
	priceArrPtr := (*[1 << 30]C.List_Float)(unsafe.Pointer(data.values))[:quoteLen:quoteLen]
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

	return C.ok_List_List_Float(data)
}

//export query_quotes_release
func query_quotes_release(result C.Result_List_List_Float) {
	C.release_Result_List_List_Float(result)
}

//export query_id_map
func query_id_map(listing_status, sort, symbol, aux C.Optional_String, start, limit C.Optional_Int) (result C.Result_List_Dict_String) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.err_List_Dict_String(C.CString("go panic"))
		}
	}()
	u, err := url.Parse(IdMapUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_Dict_String(C.CString(errStr))
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
	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_List_Dict_String(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_Dict_String(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_Dict_String(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody IdMapResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Dict_String(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_Dict_String(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data
	dataLen := len(respData)

	data := C.new_List_Dict_String(C.size_t(dataLen))
	idMapsArr := (*[1 << 30]C.Dict_String)(unsafe.Pointer(data.values))[:dataLen:dataLen]

	for idx, v := range respData {
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
	return C.ok_List_Dict_String(data)
}

//export query_id_map_release
func query_id_map_release(result C.Result_List_Dict_String) {
	C.release_Result_List_Dict_String(result)
}

// TODO: symbol请求与id和slug格式不符，暂未实现
//
//export query_metadata
func query_metadata(id, slug, address, aux C.Optional_String, skip_invalid C.Optional_Bool) (result C.Result_List_Dict_String) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.err_List_Dict_String(C.CString("go panic"))
		}
	}()
	idIsSome := bool(id.is_some)
	slugIsSome := bool(slug.is_some)
	addressIsSome := bool(address.is_some)
	auxIsSome := bool(aux.is_some)
	skipInvalidIsSome := bool(skip_invalid.is_some)

	if !(idIsSome || slugIsSome) {
		errStr := "id or slug must have at least one"
		return C.err_List_Dict_String(C.CString(errStr))
	}

	u, err := url.Parse(MetadataUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_Dict_String(C.CString(errStr))
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
		return C.err_List_Dict_String(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_Dict_String(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_Dict_String(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody MetadataResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Dict_String(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_Dict_String(C.CString(respBody.Status.ErrorMessage))
	}

	data := C.new_List_Dict_String(C.size_t(len(respBody.Data)))
	metaArr := (*[1 << 30]C.Dict_String)(unsafe.Pointer(data.values))[:data.len:data.len]
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

	return C.ok_List_Dict_String(data)
}

//export query_metadata_release
func query_metadata_release(result C.Result_List_Dict_String) {
	C.release_Result_List_Dict_String(result)
}

func parseMetadata(meta Metadata) (keys []string, values []string) {
	keys = []string{"id", "name", "symbol", "slug", "category", "description", "tags", "urls"}
	values = []string{strconv.Itoa(meta.ID), meta.Name, meta.Symbol, meta.Slug, meta.Category, meta.Description, strings.Join(meta.Tags, ","), ""}

	dataLen := len(keys)

	urls := reflect.ValueOf(meta.URLs)
	for i := 0; i < urls.Type().NumField(); i++ {
		website := urls.Field(i)
		for j := 0; j < website.Len(); j++ {
			value := website.Index(j).String()
			values[dataLen-1] += value + ","
		}
	}
	values[dataLen-1] = strings.TrimSuffix(values[dataLen-1], ",")

	keys = append(keys, "self_reported_circulating_supply")
	values = append(values, strconv.FormatFloat(meta.SelfReportedCirculatingSupply, 'f', -1, 64))

	keys = append(keys, "self_reported_market_cap")
	values = append(values, strconv.FormatFloat(meta.SelfReportedMarketCap, 'f', -1, 64))

	return
}

//export query_listings
func query_listings(start, limit, price_min, price_max, market_cap_min, market_cap_max, volume_24h_min, volume_24h_max, circulating_supply_min, circulating_supply_max, percent_change_24h_min, percent_change_24h_max C.Optional_Int, convert, convert_id, sort, sort_dir, cryptocurrency_type, tag, aux C.Optional_String) (result C.Result_List_MarketData) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.err_List_MarketData(C.CString("go panic"))
		}
	}()
	startIsSome := bool(start.is_some)
	limitIsSome := bool(limit.is_some)
	priceMinIsSome := bool(price_min.is_some)
	priceMaxIsSome := bool(price_max.is_some)
	marketCapMinIsSome := bool(market_cap_min.is_some)
	marketCapMaxIsSome := bool(market_cap_max.is_some)
	volume24hMinIsSome := bool(volume_24h_min.is_some)
	volume24hMaxIsSome := bool(volume_24h_max.is_some)
	circulatingSupplyMinIsSome := bool(circulating_supply_min.is_some)
	circulatingSupplyMaxIsSome := bool(circulating_supply_max.is_some)
	percentChange24hMinIsSome := bool(percent_change_24h_min.is_some)
	percentChange24hMaxIsSome := bool(percent_change_24h_max.is_some)
	convertIsSome := bool(convert.is_some)
	convertIdIsSome := bool(convert_id.is_some)
	sortIsSome := bool(sort.is_some)
	sortDirIsSome := bool(sort_dir.is_some)
	cryptocurrencyTypeIsSome := bool(cryptocurrency_type.is_some)
	tagIsSome := bool(tag.is_some)
	auxIsSome := bool(aux.is_some)

	u, err := url.Parse(ListingsLatestUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_MarketData(C.CString(errStr))
	}

	params := url.Values{}
	if startIsSome {
		params.Add("start", strconv.FormatInt(int64(start.value), 10))
	}
	if limitIsSome {
		params.Add("limit", strconv.FormatInt(int64(limit.value), 10))
	}
	if priceMinIsSome {
		params.Add("price_min", strconv.FormatInt(int64(price_min.value), 10))
	}
	if priceMaxIsSome {
		params.Add("price_max", strconv.FormatInt(int64(price_max.value), 10))
	}
	if marketCapMinIsSome {
		params.Add("market_cap_min", strconv.FormatInt(int64(market_cap_min.value), 10))
	}
	if marketCapMaxIsSome {
		params.Add("market_cap_max", strconv.FormatInt(int64(market_cap_max.value), 10))
	}
	if volume24hMinIsSome {
		params.Add("volume_min", strconv.FormatInt(int64(volume_24h_min.value), 10))
	}
	if volume24hMaxIsSome {
		params.Add("volume_max", strconv.FormatInt(int64(volume_24h_max.value), 10))
	}
	if circulatingSupplyMinIsSome {
		params.Add("market_cap_min", strconv.FormatInt(int64(circulating_supply_min.value), 10))
	}
	if circulatingSupplyMaxIsSome {
		params.Add("market_cap_max", strconv.FormatInt(int64(circulating_supply_max.value), 10))
	}
	if percentChange24hMinIsSome {
		params.Add("volume_min", strconv.FormatInt(int64(percent_change_24h_min.value), 10))
	}
	if percentChange24hMaxIsSome {
		params.Add("volume_max", strconv.FormatInt(int64(percent_change_24h_max.value), 10))
	}
	if convertIsSome {
		params.Add("convert", C.GoString(convert.value))
	}
	if convertIdIsSome {
		params.Add("convert_id", C.GoString(convert_id.value))
	}
	if sortIsSome {
		params.Add("sort", C.GoString(sort.value))
	}
	if sortDirIsSome {
		params.Add("sort_dir", C.GoString(sort_dir.value))
	}
	if cryptocurrencyTypeIsSome {
		params.Add("cryptocurrency_type", C.GoString(cryptocurrency_type.value))
	}
	if tagIsSome {
		params.Add("tag", C.GoString(tag.value))
	}
	if auxIsSome {
		params.Add("aux", C.GoString(aux.value))
	}
	// 将查询参数添加到 URL 查询字符串中
	u.RawQuery = params.Encode()
	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_List_MarketData(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_MarketData(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)

	defer respBodyDecomp.Close()

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_MarketData(C.CString(errStr))
	}

	var respBody ListingsResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_MarketData(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_MarketData(C.CString(respBody.Status.ErrorMessage))
	}

	data := C.new_List_MarketData(C.size_t(len(respBody.Data)))
	marketDataArr := (*[1 << 30]C.MarketData)(unsafe.Pointer(data.values))[:data.len:data.len]
	for idx, marketData := range respBody.Data {
		metaKeys, metaValues := parseMarketMetadata(marketData)

		metaCMap := C.new_Dict_String(C.size_t(len(metaKeys)))
		cMetaKeys := (*[1 << 30]C.String)(unsafe.Pointer(metaCMap.keys))[:metaCMap.len:metaCMap.len]
		cMetaValues := (*[1 << 30]C.String)(unsafe.Pointer(metaCMap.values))[:metaCMap.len:metaCMap.len]

		for i := 0; i < len(metaKeys); i++ {
			cMetaKeys[i] = C.CString(metaKeys[i])
			cMetaValues[i] = C.CString(metaValues[i])
		}

		quoteCMap := C.new_Dict_Dict_String(C.size_t(len(marketData.Quote)))
		cQuoteKeys := (*[1 << 30]C.String)(unsafe.Pointer(quoteCMap.keys))[:quoteCMap.len:quoteCMap.len]
		cQuoteValues := (*[1 << 30]C.Dict_String)(unsafe.Pointer(quoteCMap.values))[:quoteCMap.len:quoteCMap.len]

		quoteIdx := 0
		for key, quote := range marketData.Quote {
			cQuoteKeys[quoteIdx] = C.CString(key)

			quoteKeys, quoteValues := parseMarketQuote(quote)

			cQuoteValues[quoteIdx] = C.new_Dict_String(C.size_t(len(quoteKeys)))
			cQuoteValuesKeys := (*[1 << 30]C.String)(unsafe.Pointer(cQuoteValues[quoteIdx].keys))[:len(quoteKeys):len(quoteKeys)]
			cQuoteValuesValues := (*[1 << 30]C.String)(unsafe.Pointer(cQuoteValues[quoteIdx].values))[:len(quoteValues):len(quoteValues)]

			for i := 0; i < len(quoteKeys); i++ {
				cQuoteValuesKeys[i] = C.CString(quoteKeys[i])
				cQuoteValuesValues[i] = C.CString(quoteValues[i])
			}

			quoteIdx++
		}

		market := C.MarketData{
			metadata: metaCMap,
			quotes:   quoteCMap,
		}

		marketDataArr[idx] = market
	}

	return C.ok_List_MarketData(data)
}

//export query_listings_release
func query_listings_release(result C.Result_List_MarketData) {
	C.release_Result_List_MarketData(result)
}

func parseMarketMetadata(marketData ListingsData) (keys []string, values []string) {
	keys = []string{"id", "name", "symbol", "slug", "cmc_rank", "num_market_pairs", "circulating_supply", "total_supply", "max_supply", "last_updated", "date_added", "tvl_ratio", "tags"}
	values = []string{strconv.Itoa(marketData.ID), marketData.Name, marketData.Symbol, marketData.Slug, strconv.Itoa(marketData.CmcRank), strconv.Itoa(marketData.NumMarketPairs), strconv.FormatFloat(marketData.CirculatingSupply, 'f', -1, 64), strconv.FormatFloat(marketData.TotalSupply, 'f', -1, 64), strconv.FormatFloat(marketData.MaxSupply, 'f', -1, 64), marketData.LastUpdated.String(), marketData.DateAdded.String(), strconv.FormatFloat(marketData.TvlRatio, 'f', -1, 64), strings.Join(marketData.Tags, ",")}

	keys = append(keys, "self_reported_circulating_supply")
	values = append(values, strconv.FormatFloat(marketData.SelfReportedCirculatingSupply, 'f', -1, 64))

	keys = append(keys, "self_reported_market_cap")
	values = append(values, strconv.FormatFloat(marketData.SelfReportedMarketCap, 'f', -1, 64))

	return
}

func parseMarketQuote(quote Quote) (keys []string, values []string) {
	keys = []string{"price", "volume_24h", "volume_change_24h", "percent_change_1h", "percent_change_24h", "percent_change_7d", "percent_change_30d", "market_cap", "market_cap_dominance", "fully_diluted_market_cap", "last_updated"}
	values = []string{strconv.FormatFloat(quote.Price, 'f', -1, 64), strconv.FormatFloat(quote.Volume24h, 'f', -1, 64), strconv.FormatFloat(quote.VolumeChange24h, 'f', -1, 64), strconv.FormatFloat(quote.PercentChange1h, 'f', -1, 64), strconv.FormatFloat(quote.PercentChange24h, 'f', -1, 64),
		strconv.FormatFloat(quote.PercentChange7d, 'f', -1, 64),
		strconv.FormatFloat(quote.PercentChange30d, 'f', -1, 64),
		strconv.FormatFloat(quote.MarketCap, 'f', -1, 64),
		strconv.FormatFloat(quote.MarketCapDominance, 'f', -1, 64),
		strconv.FormatFloat(quote.FullyDilutedMarketCap, 'f', -1, 64), quote.LastUpdated.String()}

	return
}

//export query_categories
func query_categories(start, limit C.Optional_Int, id, slug, symbol C.Optional_String) (result C.Result_List_Category) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.err_List_Category(C.CString("go panic"))
		}
	}()

	u, err := url.Parse(CategoriesUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_Category(C.CString(errStr))
	}

	params := url.Values{}
	if bool(start.is_some) {
		params.Add("start", strconv.FormatInt(int64(start.value), 10))
	}
	if bool(limit.is_some) {
		params.Add("limit", strconv.FormatInt(int64(limit.value), 10))
	}
	if bool(id.is_some) {
		params.Add("id", C.GoString(id.value))
	}
	if bool(slug.is_some) {
		params.Add("slug", C.GoString(slug.value))
	}
	if bool(symbol.is_some) {
		params.Add("symbol", C.GoString(symbol.value))
	}
	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_List_Category(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_Category(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)

	defer respBodyDecomp.Close()

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_Category(C.CString(errStr))
	}

	var respBody CategoriesResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Category(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_Category(C.CString(respBody.Status.ErrorMessage))
	}

	data := C.new_List_Category(C.size_t(len(respBody.Data)))
	categoryArr := (*[1 << 30]C.Category)(unsafe.Pointer(data.values))[:data.len:data.len]

	for i, category := range respBody.Data {
		categoryArr[i] = C.Category{
			id:                C.CString(category.ID),
			name:              C.CString(category.Name),
			description:       C.CString(category.Description),
			num_tokens:        C.longlong(category.NumTokens),
			avg_price_change:  C.double(category.AvgPriceChange),
			market_cap:        C.double(category.MarketCap),
			market_cap_change: C.double(category.MarketCapChange),
			volume:            C.double(category.Volume),
			volume_change:     C.double(category.VolumeChange),
			last_updated:      C.CString(category.LastUpdated.String()),
		}
	}

	return C.ok_List_Category(data)
}

//export query_categories_release
func query_categories_release(result C.Result_List_Category) {
	C.release_Result_List_Category(result)
}

//export query_category
func query_category(id C.String, start, limit C.Optional_Int, convert, convert_id C.Optional_String) (result C.Result_Optional_CategorySingle) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.err_Optional_CategorySingle(C.CString("go panic"))
		}
	}()
	u, err := url.Parse(CategoryUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_Optional_CategorySingle(C.CString(errStr))
	}

	params := url.Values{}
	params.Add("id", C.GoString(id))
	if bool(start.is_some) {
		params.Add("start", strconv.FormatInt(int64(start.value), 10))
	}
	if bool(limit.is_some) {
		params.Add("limit", strconv.FormatInt(int64(limit.value), 10))
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
		return C.err_Optional_CategorySingle(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_Optional_CategorySingle(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)

	defer respBodyDecomp.Close()

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_Optional_CategorySingle(C.CString(errStr))
	}

	var respBody CategoryResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_Optional_CategorySingle(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_Optional_CategorySingle(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data
	data := C.some_CategorySingle(C.CategorySingle{
		id:                C.CString(respData.ID),
		name:              C.CString(respData.Name),
		description:       C.CString(respData.Description),
		num_tokens:        C.Int(respData.NumTokens),
		avg_price_change:  C.Float(respData.AvgPriceChange),
		market_cap:        C.Float(respData.MarketCap),
		market_cap_change: C.Float(respData.MarketCapChange),
		volume:            C.Float(respData.Volume),
		volume_change:     C.Float(respData.VolumeChange),
		last_updated:      C.CString(respData.LastUpdated.String()),
		coins:             C.new_List_Coin(C.size_t(len(respData.Coins))),
	})
	coinArr := (*[1 << 30]C.Coin)(unsafe.Pointer(data.value.coins.values))[:data.value.coins.len:data.value.coins.len]

	for i, coin := range respData.Coins {
		coinArr[i] = C.Coin{
			id:                 C.Int(coin.ID),
			name:               C.CString(coin.Name),
			symbol:             C.CString(coin.Symbol),
			slug:               C.CString(coin.Slug),
			num_market_pairs:   C.Int(coin.NumMarketPairs),
			date_added:         C.CString(coin.DateAdded.String()),
			tags:               C.CString(strings.Join(coin.Tags, ",")),
			max_supply:         C.Int(coin.MaxSupply),
			circulating_supply: C.Int(coin.CirculatingSupply),
			total_supply:       C.Int(coin.TotalSupply),
			is_active:          C.Int(coin.IsActive),
			infinite_supply:    C.Bool(coin.InfiniteSupply),
			cmc_rank:           C.Int(coin.CmcRank),
			is_fiat:            C.Int(coin.IsFiat),
			tvl_ratio:          C.Float(coin.TvlRatio),
			last_updated:       C.CString(coin.LastUpdated.String()),
			quote:              C.new_Dict_Quote(C.size_t(len(coin.Quote))),
		}
		quoteKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(coinArr[i].quote.keys))[:coinArr[i].quote.len:coinArr[i].quote.len]
		quoteValueArr := (*[1 << 30]C.Quote)(unsafe.Pointer(coinArr[i].quote.values))[:coinArr[i].quote.len:coinArr[i].quote.len]

		idx := 0
		for k, quote := range coin.Quote {
			quoteKeyArr[idx] = C.CString(k)
			quoteValueArr[idx] = C.Quote{
				price:                    C.Float(quote.Price),
				volume_24h:               C.Float(quote.Volume24h),
				volume_change_24h:        C.Float(quote.VolumeChange24h),
				percent_change_1h:        C.Float(quote.PercentChange1h),
				percent_change_24h:       C.Float(quote.PercentChange24h),
				percent_change_7d:        C.Float(quote.PercentChange7d),
				percent_change_30d:       C.Float(quote.PercentChange30d),
				percent_change_60d:       C.Float(quote.PercentChange60d),
				percent_change_90d:       C.Float(quote.PercentChange90d),
				market_cap:               C.Float(quote.MarketCap),
				market_cap_dominance:     C.Float(quote.MarketCapDominance),
				fully_diluted_market_cap: C.Float(quote.FullyDilutedMarketCap),
				tvl:                      C.Float(quote.Tvl),
				last_updated:             C.CString(quote.LastUpdated.String()),
			}
			idx++
		}

	}
	return C.ok_Optional_CategorySingle(data)
}

//export query_category_release
func query_category_release(result C.Result_Optional_CategorySingle) {
	C.release_Result_Optional_CategorySingle(result)
}

func main() {

}
