package main

/*
#cgo CFLAGS: -I../dependencies
#include <news.h>
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"time"
	"unsafe"
)

//export get_blockbeats_news
func get_blockbeats_news(date C.Optional_String, limit C.Optional_Int) C.NewsResult {
	goDate, goLimit := __parse_input(date, limit)
	newsItems, err := __get_blockbeats_news_by_date(goDate, goLimit)
	if err != nil {
		return C.NewsResult{
			success: C.Bool(false),
		}
	}
	return __convert_success_result(newsItems)
}

//export get_blockbeats_news_release
func get_blockbeats_news_release(res C.NewsResult) {
	C.release_NewsResult(res)
}

//export get_panews
func get_panews(date C.Optional_String, limit C.Optional_Int) C.NewsResult {
	goDate, goLimit := __parse_input(date, limit)
	newsItems, err := __get_panews_by_date(goDate, goLimit)
	if err != nil {
		return C.NewsResult{
			success: C.Bool(false),
		}
	}
	return __convert_success_result(newsItems)
}

//export get_panews_release
func get_panews_release(res C.NewsResult) {
	C.release_NewsResult(res)
}

//export get_blockbeats_news_by_date2json
func get_blockbeats_news_by_date2json(date C.Optional_String, limit C.Optional_Int) C.NewsResultJson {

	goDate, goLimit := __parse_input(date, limit)

	newsItems, err := __get_blockbeats_news_by_date(goDate, goLimit)
	if err != nil {
		return C.NewsResultJson{
			success: C.Bool(false),
		}
	}
	contents, _ := json.Marshal(newsItems)
	// fmt.Printf("Content: %+v\n", C.NewsResultJson{
	// 	success:     C.Bool(true),
	// 	json_string: C.CString(string(contents)),
	// })

	return C.NewsResultJson{
		success:     C.Bool(true),
		json_string: C.CString(string(contents)),
		//json_string: C.CString(string(contents)),
	}
}

//export get_blockbeats_news_by_date2json_release
func get_blockbeats_news_by_date2json_release(res C.NewsResultJson) {
	C.release_NewsResultJson(res)
}

//export get_panews_by_date2json
func get_panews_by_date2json(date C.Optional_String, limit C.Optional_Int) C.NewsResultJson {

	goDate, goLimit := __parse_input(date, limit)

	newsItems, err := __get_panews_by_date(goDate, goLimit)
	if err != nil {
		fmt.Printf("got error: %+v\n", err)
		return C.NewsResultJson{
			success: C.Bool(false),
		}
	}
	contents, _ := json.Marshal(newsItems)
	fmt.Printf("Content: %+v\n", C.NewsResultJson{
		success:     C.Bool(true),
		json_string: C.CString(string(contents)),
	})

	return C.NewsResultJson{
		success:     C.Bool(true),
		json_string: C.CString(string(contents)),
		//json_string: C.CString(string(contents)),
	}
}

//export get_panews_by_date2json_release
func get_panews_by_date2json_release(res C.NewsResultJson) {
	fmt.Println("get_panews_by_date2json_release")
	C.release_NewsResultJson(res)
}

func main() {}

func __parse_input(date C.Optional_String, limit C.Optional_Int) (string, int) {
	goDate := time.Now().Format(time.DateOnly)
	if date.value != nil {
		goDate = C.GoString(date.value)
	}

	goLimit := 5
	if int(limit.value) != 0 {
		goLimit = int(limit.value)
	}

	fmt.Printf("__parse_inputs %s %d\n", goDate, goLimit)
	return goDate, goLimit
}

func __convert_success_result(items []*Custom_FlashNewsItem) C.NewsResult {
	res := C.NewsResult{
		success: C.Bool(true),
		items:   C.new_List_NewsItem(C.ulong(len(items))),
	}

	for i, item := range items {
		item_ptr := (*C.NewsItem)(unsafe.Pointer(uintptr(unsafe.Pointer(res.items.values)) + uintptr(i)*unsafe.Sizeof(*res.items.values)))
		item_ptr.content = C.CString(item.Content)
		item_ptr.title = C.CString(item.Title)
		item_ptr.unixtime = C.Int(item.AddTime)

	}

	return res
}
