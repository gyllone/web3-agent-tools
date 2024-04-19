package main

/*
#cgo CFLAGS: -I../dependencies
#include <marsbit.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"runtime/debug"
	"strconv"
	"time"
	"unsafe"
)

//export query_news
func query_news(query_time, page_size C.Optional_Int, lang C.Optional_String) C.Result_List_News {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()

	u, err := url.Parse(ShowInfoUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_List_News(C.CString(errStr))
	}

	params := url.Values{}
	if bool(query_time.is_some) {
		params.Add("queryTime", strconv.FormatInt(int64(query_time.value), 10))
	}
	if bool(page_size.is_some) {
		params.Add("pageSize", strconv.FormatInt(int64(page_size.value), 10))
	}
	if bool(lang.is_some) {
		params.Add("lang", C.GoString(lang.value))
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

	if respBody.Message != "ok" {
		return C.err_List_News(C.CString(respBody.Message))
	}

	data := C.new_List_News(C.size_t(len(respBody.Data)))
	if data.len == 0 {
		return C.ok_List_News(data)
	}

	dataArr := (*[1 << 30]C.News)(unsafe.Pointer(data.values))[:data.len:data.len]

	for i, v := range respBody.Data {
		dataArr[i] = C.News{
			news_type:    C.CString(v.NewsType),
			title:        C.CString(v.News.Title),
			content:      C.CString(removeHTMLTags(v.News.Content)),
			synopsis:     C.CString(v.News.Synopsis),
			publish_time: C.CString(unix2UTC(v.PublishTime)),
		}
	}

	return C.ok_List_News(data)
}

//export query_news_release
func query_news_release(result C.Result_List_News) {
	C.release_Result_List_News(result)
}

func removeHTMLTags(input string) string {
	r, _ := regexp.Compile("<[^>]*>")
	return r.ReplaceAllString(input, "")
}

func unix2UTC(unix int64) string {
	return time.Unix(unix/1000, 0).UTC().Format(time.RFC3339)
}

//export query_multisearch
func query_multisearch(q C.String, page, page_size C.Optional_Int) C.Result_Optional_SearchedNewsObj {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("------go print start------")
			fmt.Println("Recovered from panic:", r)
			fmt.Println("Stack trace:")
			fmt.Println("------go print end------")
			debug.PrintStack()
		}
	}()

	u, err := url.Parse(MultiSearchUrl)
	if err != nil {
		errStr := "Failed to parse URL"
		return C.err_Optional_SearchedNewsObj(C.CString(errStr))
	}

	params := url.Values{}
	params.Add("type", "0")
	params.Add("excellentFlag", "0")
	params.Add("q", C.GoString(q))
	if bool(page.is_some) {
		params.Add("page", strconv.FormatInt(int64(page.value), 10))
	}
	if bool(page_size.is_some) {
		params.Add("pageSize", strconv.FormatInt(int64(page_size.value), 10))
	}

	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		errStr := "Failed to send request"
		return C.err_Optional_SearchedNewsObj(C.CString(errStr))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return C.err_Optional_SearchedNewsObj(
			C.CString(fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)))
	}

	var respBody SearchedNewsResp

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		errStr := "Failed to decode response\n" + err.Error()
		return C.err_Optional_SearchedNewsObj(C.CString(errStr))
	}

	if respBody.Message != "success" {
		return C.err_Optional_SearchedNewsObj(C.CString(respBody.Message))
	}

	data := C.some_SearchedNewsObj(C.SearchedNewsObj{
		News:          getCSearchedNewsList(respBody.Data.News.InforList),
		Lives:         getCSearchedNewsList(respBody.Data.Lives.InforList),
		ExcellentNews: getCSearchedNewsList(respBody.Data.ExcellentNews.InforList),
	})

	return C.ok_Optional_SearchedNewsObj(data)
}

func getCSearchedNewsList(list []searchedNews) C.List_SearchedNews {
	res := C.new_List_SearchedNews(C.size_t(len(list)))

	if res.len == 0 {
		return res
	}

	resArr := (*[1 << 30]C.SearchedNews)(unsafe.Pointer(res.values))[:res.len:res.len]

	for i, v := range list {
		resArr[i] = C.SearchedNews{
			title:        C.CString(v.Title),
			content:      C.CString(v.Content),
			synopsis:     C.CString(v.Synopsis),
			publish_time: C.CString(unix2UTC(v.PublishTime)),
		}
	}

	return res
}

//export query_multisearch_release
func query_multisearch_release(result C.Result_Optional_SearchedNewsObj) {
	C.release_Result_Optional_SearchedNewsObj(result)
}

func main() {}
