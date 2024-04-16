package main

/*
#cgo CFLAGS: -I../../dependencies
#include <fiat.h>
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

//export query_id_map
func query_id_map(start, limit C.Optional_Int, sort C.Optional_String, include_metals C.Optional_Bool) (result C.Result_List_Fiat) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
			result = C.err_List_Fiat(C.CString("go panic"))
		}
	}()

	u, err := url.Parse(IdMapUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_Fiat(C.CString(errStr))
	}

	params := url.Values{}
	if bool(start.is_some) {
		params.Add("start", strconv.FormatInt(int64(start.value), 10))
	}
	if bool(limit.is_some) {
		params.Add("limit", strconv.FormatInt(int64(limit.value), 10))
	}
	if bool(sort.is_some) {
		params.Add("sort", C.GoString(sort.value))
	}
	if bool(include_metals.is_some) {
		params.Add("include_metals", strconv.FormatBool(bool(include_metals.value)))
	}
	u.RawQuery = params.Encode()

	req, err1 := http.NewRequest("GET", u.String(), nil)
	if err1 != nil {
		errStr := "Failed to create request"
		return C.err_List_Fiat(C.CString(errStr))
	}

	req.Header.Set("X-CMC_PRO_API_KEY", utils.ApiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")

	client := &http.Client{}
	response, err2 := client.Do(req)
	if err2 != nil {
		errStr := "Failed to send request"
		return C.err_List_Fiat(C.CString(errStr))
	}

	defer response.Body.Close()

	respBodyDecomp, err3 := utils.DecompressResponse(response)
	if err3 != nil {
		errStr := "Failed to decompress response"
		return C.err_List_Fiat(C.CString(errStr))
	}

	defer respBodyDecomp.Close()

	var respBody FiatResp

	err = json.NewDecoder(respBodyDecomp).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Fiat(C.CString(errStr))
	}

	if response.StatusCode != 200 {
		return C.err_List_Fiat(C.CString(respBody.Status.ErrorMessage))
	}

	respData := respBody.Data

	data := C.new_List_Fiat(C.size_t(len(respData)))
	fiatArr := (*[1 << 30]C.Fiat)(unsafe.Pointer(data.values))[:data.len:data.len]

	for i, fiat := range respData {
		fiatArr[i] = C.Fiat{
			id:     C.Int(fiat.ID),
			name:   C.CString(fiat.Name),
			sign:   C.CString(fiat.Sign),
			symbol: C.CString(fiat.Symbol),
		}
	}

	return C.ok_List_Fiat(data)
}

//export query_id_map_release
func query_id_map_release(result C.Result_List_Fiat) {
	C.release_Result_List_Fiat(result)
}

func main() {}
