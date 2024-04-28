package main

/*
#cgo CFLAGS: -I../dependencies
#include <odaily.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"unsafe"
)

//export query_newsflashes
func query_newsflashes(per_page C.Optional_Int, is_import C.Optional_Bool) C.Result_List_News {
	u, err := url.Parse(NewsflashesUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_News(C.CString(errStr))
	}

	params := url.Values{}
	if bool(per_page.is_some) {
		params.Add("per_page", strconv.FormatInt(int64(per_page.value), 10))
	}
	if bool(is_import.is_some) {
		params.Add("is_import", strconv.FormatBool(bool(is_import.value)))
	}
	u.RawQuery = params.Encode()

	resp, err1 := http.Get(u.String())
	if err1 != nil {
		errStr := "Failed to send request"
		return C.err_List_News(C.CString(errStr))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return C.err_List_News(
			C.CString(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)))
	}

	var respBody NewsResp

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_News(C.CString(errStr))
	}

	if len(respBody.Message) != 0 {
		return C.err_List_News(C.CString(respBody.Message))
	}

	data := C.new_List_News(C.size_t(len(respBody.Data.Items)))
	if data.len == 0 {
		return C.ok_List_News(data)
	}

	dataArr := (*[1 << 30]C.News)(unsafe.Pointer(data.values))[:data.len:data.len]

	for i, v := range respBody.Data.Items {
		dataArr[i] = C.News{
			id:              C.Int(v.ID),
			is_top:          C.Int(v.IsTop),
			title:           C.CString(v.Title),
			description:     C.CString(v.Description),
			cover:           C.CString(v.Cover),
			news_url:        C.CString(v.NewsURL),
			extraction_tags: C.CString(v.ExtractionTags),
			updated_at:      C.CString(v.UpdatedAt),
		}
	}

	return C.ok_List_News(data)
}

//export query_newsflashes_release
func query_newsflashes_release(result C.Result_List_News) {
	C.release_Result_List_News(result)
}

//export query_post_list
func query_post_list(type_str C.Optional_String) C.Result_List_Post {
	u, err := url.Parse(PostListUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_Post(C.CString(errStr))
	}

	params := url.Values{}
	if bool(type_str.is_some) {
		params.Add("type", C.GoString(type_str.value))
	}
	u.RawQuery = params.Encode()

	resp, err1 := http.Get(u.String())

	if err1 != nil {
		errStr := "Failed to send request"
		return C.err_List_Post(C.CString(errStr))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return C.err_List_Post(C.CString(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)))
	}

	var respBody PostResp
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_List_Post(C.CString(errStr))
	}

	if respBody.Message != "success" {
		return C.err_List_Post(C.CString(respBody.Message))
	}
	data := C.new_List_Post(C.size_t(len(respBody.Data.Items)))

	if data.len == 0 {
		return C.ok_List_Post(data)
	}

	dataArr := (*[1 << 30]C.Post)(unsafe.Pointer(data.values))[:data.len:data.len]

	for i, v := range respBody.Data.Items {
		dataArr[i] = C.Post{
			id:         C.Int(v.ID),
			title:      C.CString(v.Title),
			summary:    C.CString(v.Summary),
			updated_at: C.CString(v.UpdatedAt),
		}
	}

	return C.ok_List_Post(data)
}

//export query_post_list_release
func query_post_list_release(result C.Result_List_Post) {
	C.release_Result_List_Post(result)
}

func main() {}
