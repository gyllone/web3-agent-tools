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
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"
	"unsafe"
)

// TODO: 使用symbol参数请求的话，返回格式与id和slug不统一，暂不支持
//
//export query_quotes_latest
func query_quotes_latest(id, slug, convert, convert_id, aux C.Optional_String, skip_invalid C.Optional_Bool) C.Result_Dict_QuoteData {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()
	idsIsSome := bool(id.is_some)
	slugIsSome := bool(slug.is_some)
	if !(idsIsSome || slugIsSome) {
		errStr := "id or slug must have at least one"
		return C.err_Dict_QuoteData(C.CString(errStr))
	}
	convertIsSome := bool(convert.is_some)
	convertIdIsSome := bool(convert_id.is_some)
	auxIsSome := bool(aux.is_some)
	skipInvalidIsSome := bool(skip_invalid.is_some)

	u, err := url.Parse(QuotesUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_Dict_QuoteData(C.CString(errStr))
	}

	params := url.Values{}
	if idsIsSome {
		idStr := C.GoString(id.value)
		params.Add("id", idStr)
	}
	if slugIsSome {
		slugStr := C.GoString(slug.value)
		params.Add("slug", slugStr)
	}
	if convertIsSome {
		convertStr := C.GoString(convert.value)
		params.Add("convert", convertStr)
	}
	if convertIdIsSome {
		convertIdStr := C.GoString(convert_id.value)
		params.Add("convert_id", convertIdStr)
	}
	if auxIsSome {
		auxStr := C.GoString(aux.value)
		params.Add("aux", auxStr)
	}
	if skipInvalidIsSome {
		skipInvalidStr := strconv.FormatBool(bool(skip_invalid.value))
		params.Add("skip_invalid", skipInvalidStr)
	}

	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_Dict_QuoteData(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_Dict_QuoteData(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_Dict_QuoteData(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody QuoteResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response" + err.Error()
		return C.err_Dict_QuoteData(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_Dict_QuoteData(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.new_Dict_QuoteData(C.size_t(len(respData)))

	// C.new_xxx(0) 并不会分配内存, 所以下面unsafe.Pointer() 会pointer panic
	if data.len == 0 {
		return C.ok_Dict_QuoteData(data)
	}

	dataKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(data.keys))[:data.len:data.len]
	dataValueArr := (*[1 << 30]C.QuoteData)(unsafe.Pointer(data.values))[:data.len:data.len]

	dataIdx := 0
	for k, v := range respData {
		dataKeyArr[dataIdx] = C.CString(k)
		dataValueArr[dataIdx] = C.QuoteData{
			id:                               C.Int(v.ID),
			name:                             C.CString(v.Name),
			symbol:                           C.CString(v.Symbol),
			slug:                             C.CString(v.Slug),
			num_market_pairs:                 C.Int(v.NumMarketPairs),
			date_added:                       C.CString(v.DateAdded.String()),
			tags:                             getCQuoteDataTags(v.Tags),
			max_supply:                       C.Int(v.MaxSupply),
			circulating_supply:               C.Float(v.CirculatingSupply),
			total_supply:                     C.Float(v.TotalSupply),
			is_active:                        C.Int(v.IsActive),
			infinite_supply:                  C.Bool(v.InfiniteSupply),
			platform:                         getCPlatform(v.Platform),
			cmc_rank:                         C.Int(v.CmcRank),
			is_fiat:                          C.Int(v.IsFiat),
			self_reported_circulating_supply: C.Float(v.SelfReportedCirculatingSupply),
			self_reported_market_cap:         C.Float(v.SelfReportedMarketCap),
			tvl_ratio:                        C.Float(v.TvlRatio),
			last_updated:                     C.CString(v.LastUpdated.String()),
			quote:                            getCQuoteDict(v.Quote),
		}

		dataIdx++
	}

	return C.ok_Dict_QuoteData(data)
}

func getCPlatform(platform *Platform) C.Optional_Platform {
	if platform == nil {
		return C.none_Platform()
	}

	return C.some_Platform(C.Platform{
		id:            C.Int(int64(platform.ID)),
		name:          C.CString(platform.Name),
		symbol:        C.CString(platform.Symbol),
		slug:          C.CString(platform.Slug),
		token_address: C.CString(platform.TokenAddress),
	})
}

func getCQuoteDataTags(tags []Tags) C.List_Tag {
	res := C.new_List_Tag(C.size_t(len(tags)))
	if res.len == 0 {
		return res
	}
	tagArr := (*[1 << 30]C.Tag)(unsafe.Pointer(res.values))[:res.len:res.len]

	for i, tag := range tags {
		tagArr[i] = C.Tag{
			name:     C.CString(tag.Name),
			slug:     C.CString(tag.Slug),
			category: C.CString(tag.Category),
		}
	}

	return res
}

func getCQuoteDict(quotes map[string]Quote) C.Dict_Quote {
	res := C.new_Dict_Quote(C.size_t(len(quotes)))

	if res.len == 0 {
		return res
	}

	resKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(res.keys))[:res.len:res.len]
	resValueArr := (*[1 << 30]C.Quote)(unsafe.Pointer(res.values))[:res.len:res.len]

	quoteIdx := 0
	for quoteKey, quote := range quotes {
		resKeyArr[quoteIdx] = C.CString(quoteKey)
		resValueArr[quoteIdx] = C.Quote{
			last_updated:             C.CString(quote.LastUpdated.String()),
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
		}
		quoteIdx++
	}

	return res
}

//export query_quotes_latest_release
func query_quotes_latest_release(result C.Result_Dict_QuoteData) {
	C.release_Result_Dict_QuoteData(result)
}

//export query_id_map
func query_id_map(listing_status, sort, symbol, aux C.Optional_String, start, limit C.Optional_Int) C.Result_List_Cryptocurrency {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()
	u, err := url.Parse(IdMapUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_Cryptocurrency(C.CString(errStr))
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
		params.Add("start", strconv.FormatInt(int64(start.value), 10))
	}
	if limitIsSome {
		params.Add("limit", strconv.FormatInt(int64(limit.value), 10))
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
		return C.err_List_Cryptocurrency(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_Cryptocurrency(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_Cryptocurrency(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody IdMapResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Cryptocurrency(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_Cryptocurrency(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.new_List_Cryptocurrency(C.size_t(len(respData)))

	if data.len == 0 {
		return C.ok_List_Cryptocurrency(data)
	}

	dataArr := (*[1 << 30]C.Cryptocurrency)(unsafe.Pointer(data.values))[:data.len:data.len]

	for idx, v := range respData {
		dataArr[idx] = C.Cryptocurrency{
			id:                    C.Int(v.ID),
			rank:                  C.Int(v.Rank),
			name:                  C.CString(v.Name),
			symbol:                C.CString(v.Symbol),
			slug:                  C.CString(v.Slug),
			is_active:             C.Int(v.IsActive),
			first_historical_data: C.CString(v.FirstHistoricalData.String()),
			last_historical_data:  C.CString(v.LastHistoricalData.String()),
			platform:              getCPlatform(v.Platform),
		}
	}

	return C.ok_List_Cryptocurrency(data)
}

//export query_id_map_release
func query_id_map_release(result C.Result_List_Cryptocurrency) {
	C.release_Result_List_Cryptocurrency(result)
}

// TODO: symbol请求与id和slug格式不符，暂未实现
//
//export query_metadata
func query_metadata(id, slug, address, aux C.Optional_String, skip_invalid C.Optional_Bool) C.Result_Dict_Metadata {
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
	slugIsSome := bool(slug.is_some)
	addressIsSome := bool(address.is_some)
	auxIsSome := bool(aux.is_some)
	skipInvalidIsSome := bool(skip_invalid.is_some)

	if !(idIsSome || slugIsSome) {
		errStr := "id or slug must have at least one"
		return C.err_Dict_Metadata(C.CString(errStr))
	}

	u, err := url.Parse(MetadataUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_Dict_Metadata(C.CString(errStr))
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
		return C.err_Dict_Metadata(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_Dict_Metadata(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_Dict_Metadata(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody MetadataResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_Dict_Metadata(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_Dict_Metadata(C.CString(respBody.Status.ErrorMessage))
	}

	data := C.new_Dict_Metadata(C.size_t(len(respBody.Data)))

	if data.len == 0 {
		return C.ok_Dict_Metadata(data)
	}

	dataKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(data.keys))[:data.len:data.len]
	dataValueArr := (*[1 << 30]C.Metadata)(unsafe.Pointer(data.values))[:data.len:data.len]

	dataIdx := 0

	for k, v := range respBody.Data {
		dataKeyArr[dataIdx] = C.CString(k)
		dataValueArr[dataIdx] = C.Metadata{
			id:          C.Int(v.ID),
			name:        C.CString(v.Name),
			symbol:      C.CString(v.Symbol),
			category:    C.CString(v.Category),
			description: C.CString(v.Description),
			slug:        C.CString(v.Slug),
			logo:        C.CString(v.Logo),
			subreddit:   C.CString(v.Subreddit),
			notice:      C.CString(v.Notice),
			tags:        getCStringList(v.Tags),
			tag_names:   getCStringList(v.TagNames),
			tag_groups:  getCStringList(v.TagGroups),
			urls: C.URLs{
				website:       getCStringList(v.URLs.Website),
				twitter:       getCStringList(v.URLs.Twitter),
				message_board: getCStringList(v.URLs.MessageBoard),
				chat:          getCStringList(v.URLs.Chat),
				facebook:      getCStringList(v.URLs.Facebook),
				explorer:      getCStringList(v.URLs.Explorer),
				reddit:        getCStringList(v.URLs.Reddit),
				technical_doc: getCStringList(v.URLs.TechnicalDoc),
				source_code:   getCStringList(v.URLs.SourceCode),
				announcement:  getCStringList(v.URLs.Announcement),
			},
			platform:                         getCPlatform(v.Platform),
			date_added:                       C.CString(v.DateAdded.String()),
			twitter_username:                 C.CString(v.TwitterUsername),
			is_hidden:                        C.Int(v.IsHidden),
			date_launched:                    C.CString(v.DateLaunched.String()),
			self_reported_circulating_supply: C.Float(v.SelfReportedCirculatingSupply),
			self_reported_market_cap:         C.Float(v.SelfReportedMarketCap),
			infinite_supply:                  C.Bool(v.InfiniteSupply),
		}

		dataIdx++
	}

	return C.ok_Dict_Metadata(data)
}

func getCStringList(tags []string) C.List_String {
	res := C.new_List_String(C.size_t(len(tags)))

	if res.len == 0 {
		return res
	}

	resArr := (*[1 << 30]C.String)(unsafe.Pointer(res.values))[:res.len:res.len]

	for i, tag := range tags {
		resArr[i] = C.CString(tag)
	}
	return res
}

//export query_metadata_release
func query_metadata_release(result C.Result_Dict_Metadata) {
	C.release_Result_Dict_Metadata(result)
}

//export query_listings
func query_listings(start, limit, price_min, price_max, market_cap_min, market_cap_max, volume_24h_min, volume_24h_max, circulating_supply_min, circulating_supply_max, percent_change_24h_min, percent_change_24h_max C.Optional_Int, convert, convert_id, sort, sort_dir, cryptocurrency_type, tag, aux C.Optional_String) C.Result_List_MarketData {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()

	u, err := url.Parse(ListingsLatestUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_MarketData(C.CString(errStr))
	}

	params := url.Values{}
	if bool(start.is_some) {
		params.Add("start", strconv.FormatInt(int64(start.value), 10))
	}
	if bool(limit.is_some) {
		params.Add("limit", strconv.FormatInt(int64(limit.value), 10))
	}
	if bool(price_min.is_some) {
		params.Add("price_min", strconv.FormatInt(int64(price_min.value), 10))
	}
	if bool(price_max.is_some) {
		params.Add("price_max", strconv.FormatInt(int64(price_max.value), 10))
	}
	if bool(market_cap_min.is_some) {
		params.Add("market_cap_min", strconv.FormatInt(int64(market_cap_min.value), 10))
	}
	if bool(market_cap_max.is_some) {
		params.Add("market_cap_max", strconv.FormatInt(int64(market_cap_max.value), 10))
	}
	if bool(volume_24h_min.is_some) {
		params.Add("volume_min", strconv.FormatInt(int64(volume_24h_min.value), 10))
	}
	if bool(volume_24h_max.is_some) {
		params.Add("volume_max", strconv.FormatInt(int64(volume_24h_max.value), 10))
	}
	if bool(circulating_supply_min.is_some) {
		params.Add("market_cap_min", strconv.FormatInt(int64(circulating_supply_min.value), 10))
	}
	if bool(circulating_supply_max.is_some) {
		params.Add("market_cap_max", strconv.FormatInt(int64(circulating_supply_max.value), 10))
	}
	if bool(percent_change_24h_min.is_some) {
		params.Add("volume_min", strconv.FormatInt(int64(percent_change_24h_min.value), 10))
	}
	if bool(percent_change_24h_max.is_some) {
		params.Add("volume_max", strconv.FormatInt(int64(percent_change_24h_max.value), 10))
	}
	if bool(convert.is_some) {
		params.Add("convert", C.GoString(convert.value))
	}
	if bool(convert_id.is_some) {
		params.Add("convert_id", C.GoString(convert_id.value))
	}
	if bool(sort.is_some) {
		params.Add("sort", C.GoString(sort.value))
	}
	if bool(sort_dir.is_some) {
		params.Add("sort_dir", C.GoString(sort_dir.value))
	}
	if bool(cryptocurrency_type.is_some) {
		params.Add("cryptocurrency_type", C.GoString(cryptocurrency_type.value))
	}
	if bool(tag.is_some) {
		params.Add("tag", C.GoString(tag.value))
	}
	if bool(aux.is_some) {
		params.Add("aux", C.GoString(aux.value))
	}

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
	if data.len == 0 {
		return C.ok_List_MarketData(data)
	}

	marketDataArr := (*[1 << 30]C.MarketData)(unsafe.Pointer(data.values))[:data.len:data.len]

	for idx, v := range respBody.Data {
		marketDataArr[idx] = C.MarketData{
			id:                               C.Int(v.ID),
			name:                             C.CString(v.Name),
			symbol:                           C.CString(v.Symbol),
			slug:                             C.CString(v.Slug),
			num_market_pairs:                 C.Int(v.NumMarketPairs),
			date_added:                       C.CString(v.DateAdded.String()),
			tags:                             getCStringList(v.Tags),
			max_supply:                       C.Int(v.MaxSupply),
			circulating_supply:               C.Int(v.CirculatingSupply),
			total_supply:                     C.Int(v.TotalSupply),
			infinite_supply:                  C.Bool(v.InfiniteSupply),
			platform:                         getCPlatform(v.Platform),
			cmc_rank:                         C.Int(v.CmcRank),
			self_reported_circulating_supply: C.Float(v.SelfReportedCirculatingSupply),
			self_reported_market_cap:         C.Float(v.SelfReportedMarketCap),
			tvl_ratio:                        C.Float(v.TvlRatio),
			last_updated:                     C.CString(v.LastUpdated.String()),
			quote:                            getCQuoteDict(v.Quote),
		}
	}

	return C.ok_List_MarketData(data)
}

//export query_listings_release
func query_listings_release(result C.Result_List_MarketData) {
	C.release_Result_List_MarketData(result)
}

//export query_categories
func query_categories(start, limit C.Optional_Int, id, slug, symbol C.Optional_String) C.Result_List_Category {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
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

	if data.len == 0 {
		return C.ok_List_Category(data)
	}
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
func query_category(id C.String, start, limit C.Optional_Int, convert, convert_id C.Optional_String) C.Result_Optional_CategorySingle {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
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
		coins:             getCCoinList(respData.Coins),
	})

	return C.ok_Optional_CategorySingle(data)
}

func getCCoinList(coins []Coin) C.List_Coin {
	res := C.new_List_Coin(C.size_t(len(coins)))

	if res.len == 0 {
		return res
	}
	coinArr := (*[1 << 30]C.Coin)(unsafe.Pointer(res.values))[:res.len:res.len]

	for i, coin := range coins {
		coinArr[i] = C.Coin{
			id:                 C.Int(coin.ID),
			name:               C.CString(coin.Name),
			symbol:             C.CString(coin.Symbol),
			slug:               C.CString(coin.Slug),
			num_market_pairs:   C.Int(coin.NumMarketPairs),
			date_added:         C.CString(coin.DateAdded.String()),
			tags:               getCStringList(coin.Tags),
			max_supply:         C.Int(coin.MaxSupply),
			circulating_supply: C.Int(coin.CirculatingSupply),
			total_supply:       C.Int(coin.TotalSupply),
			is_active:          C.Int(coin.IsActive),
			infinite_supply:    C.Bool(coin.InfiniteSupply),
			cmc_rank:           C.Int(coin.CmcRank),
			is_fiat:            C.Int(coin.IsFiat),
			tvl_ratio:          C.Float(coin.TvlRatio),
			last_updated:       C.CString(coin.LastUpdated.String()),
			quote:              getCQuoteDict(coin.Quote),
		}
	}

	return res
}

//export query_category_release
func query_category_release(result C.Result_Optional_CategorySingle) {
	C.release_Result_Optional_CategorySingle(result)
}

func main() {

}
