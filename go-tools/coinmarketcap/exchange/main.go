package main

/*
#cgo CFLAGS: -I../../dependencies
#include <exchange.h>
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
	"strings"
	"unsafe"
)

//export query_id_map
func query_id_map(start, limit C.Optional_Int, listing_status, slug, sort, aux, crypto_id C.Optional_String) C.Result_List_Exchange {
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
		return C.err_List_Exchange(C.CString(errStr))
	}

	params := url.Values{}
	if bool(start.is_some) {
		params.Add("start", strconv.FormatInt(int64(start.value), 10))
	}
	if bool(limit.is_some) {
		params.Add("limit", strconv.FormatInt(int64(limit.value), 10))
	}
	if bool(listing_status.is_some) {
		params.Add("listing_status", C.GoString(listing_status.value))
	}
	if bool(slug.is_some) {
		params.Add("slug", C.GoString(slug.value))
	}
	if bool(sort.is_some) {
		params.Add("sort", C.GoString(sort.value))
	}
	if bool(aux.is_some) {
		params.Add("aux", C.GoString(aux.value))
	}
	if bool(crypto_id.is_some) {
		params.Add("crypto_id", C.GoString(crypto_id.value))
	}
	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_List_Exchange(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_Exchange(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_Exchange(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody IdMapResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Exchange(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_Exchange(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.new_List_Exchange(C.size_t(len(respData)))
	exchangeArr := (*[1 << 30]C.Exchange)(unsafe.Pointer(data.values))[:data.len:data.len]

	for i, exchange := range respData {
		exchangeArr[i] = C.Exchange{
			id:                    C.Int(exchange.ID),
			name:                  C.CString(exchange.Name),
			slug:                  C.CString(exchange.Slug),
			is_active:             C.Int(exchange.IsActive),
			is_listed:             C.Int(exchange.IsListed),
			is_redistributable:    C.Int(exchange.IsRedistributable),
			first_historical_data: C.CString(exchange.FirstHistoricalData.String()),
			last_historical_data:  C.CString(exchange.LastHistoricalData.String()),
		}
	}

	return C.ok_List_Exchange(data)
}

//export query_id_map_release
func query_id_map_release(result C.Result_List_Exchange) {
	C.release_Result_List_Exchange(result)
}

//export query_metadata
func query_metadata(id, slug, aux C.Optional_String) C.Result_Dict_Metadata {
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
	if !(idIsSome || slugIsSome) {
		return C.err_Dict_Metadata(C.CString("At least one \"id\" or \"slug\" is required."))
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
	if bool(aux.is_some) {
		params.Add("aux", C.GoString(aux.value))
	}
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

	defer respBodyDecomp.Close()

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_Dict_Metadata(C.CString(errStr))
	}

	var respBody MetadataResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_Dict_Metadata(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_Dict_Metadata(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.new_Dict_Metadata(C.size_t(len(respData)))

	dataKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(data.keys))[:data.len:data.len]
	dataValueArr := (*[1 << 30]C.Metadata)(unsafe.Pointer(data.values))[:data.len:data.len]

	mapIdx := 0

	for k, meta := range respData {
		dataKeyArr[mapIdx] = C.CString(k)
		dataValueArr[mapIdx] = C.Metadata{
			id:                       C.Int(meta.ID),
			name:                     C.CString(meta.Name),
			slug:                     C.CString(meta.Slug),
			description:              C.CString(meta.Description),
			notice:                   C.CString(meta.Notice),
			fiats:                    getMetadataFiats(meta.Fiats),
			urls:                     C.CString(joinUrls(meta.URLs)),
			date_launched:            C.CString(meta.DateLaunched.String()),
			maker_fee:                C.Float(meta.MakerFee),
			taker_fee:                C.Float(meta.TakerFee),
			spot_volume_usd:          C.Float(meta.SpotVolumeUSD),
			spot_volume_last_updated: C.CString(meta.SpotVolumeLastUpdated.String()),
			weekly_visits:            C.Int(meta.WeeklyVisits),
		}
		mapIdx++
	}

	return C.ok_Dict_Metadata(data)
}

func getMetadataFiats(fiats []string) C.List_String {
	res := C.new_List_String(C.size_t(len(fiats)))

	resArr := (*[1 << 30]C.String)(unsafe.Pointer(res.values))[:res.len:res.len]

	for i, fiat := range fiats {
		resArr[i] = C.CString(strings.Trim(fiat, " "))
	}

	return res
}

func joinUrls(urls URLs) string {
	websiteUrls := make([]string, 0)

	if str := strings.Join(urls.Twitter, ","); len(str) > 0 {
		websiteUrls = append(websiteUrls, str)
	}
	if str := strings.Join(urls.Blog, ","); len(str) > 0 {
		websiteUrls = append(websiteUrls, str)
	}
	if str := strings.Join(urls.Website, ","); len(str) > 0 {
		websiteUrls = append(websiteUrls, str)
	}
	if str := strings.Join(urls.Chat, ","); len(str) > 0 {
		websiteUrls = append(websiteUrls, str)
	}
	if str := strings.Join(urls.Actual, ","); len(str) > 0 {
		websiteUrls = append(websiteUrls, str)
	}
	if str := strings.Join(urls.Fee, ","); len(str) > 0 {
		websiteUrls = append(websiteUrls, str)
	}

	return strings.Join(websiteUrls, ",")
}

//export query_metadata_release
func query_metadata_release(result C.Result_Dict_Metadata) {
	C.release_Result_Dict_Metadata(result)
}

//export query_exchange_assets
func query_exchange_assets(id C.String) C.Result_List_Asset {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()
	u, err := url.Parse(ExchangeAssetsUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_Asset(C.CString(errStr))
	}

	params := url.Values{}
	params.Add("id", C.GoString(id))
	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_List_Asset(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_Asset(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)

	defer respBodyDecomp.Close()

	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_Asset(C.CString(errStr))
	}

	var respBody AssetResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Asset(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_Asset(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.new_List_Asset(C.size_t(len(respData)))

	assetArr := (*[1 << 30]C.Asset)(unsafe.Pointer(data.values))[:data.len:data.len]

	for i, asset := range respData {
		assetArr[i] = C.Asset{
			wallet_address: C.CString(asset.WalletAddress),
			balance:        C.Float(asset.Balance),
			platform: C.Platform{
				crypto_id: C.Int(asset.Platform.CryptoID),
				symbol:    C.CString(asset.Platform.Symbol),
				name:      C.CString(asset.Platform.Name),
			},
			currency: C.Currency{
				crypto_id: C.Int(asset.Currency.CryptoID),
				price_usd: C.Float(asset.Currency.PriceUSD),
				symbol:    C.CString(asset.Currency.Symbol),
				name:      C.CString(asset.Currency.Name),
			},
		}
	}

	return C.ok_List_Asset(data)
}

//export query_exchange_assets_release
func query_exchange_assets_release(result C.Result_List_Asset) {
	C.release_Result_List_Asset(result)
}

func main() {}
