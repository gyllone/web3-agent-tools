package main

/*
#cgo CFLAGS: -I../dependencies
#include <news.h>
*/
import "C"
import (
	"time"
	"unsafe"
)

//export get_blockbeats_news
func get_blockbeats_news(date C.Optional_String, limit C.Optional_Int) C.Result_List_NewsItem {
	goDate, goLimit := __parse_input(date, limit)
	newsItems, err := __get_blockbeats_news_by_date(goDate, goLimit)
	if err != nil {
		return C.err_List_NewsItem(C.CString(err.Error()))
	}
	return __convert_success_result(newsItems)
}

//export get_blockbeats_news_release
func get_blockbeats_news_release(res C.Result_List_NewsItem) {
	C.release_Result_List_NewsItem(res)
}

//export get_panews
func get_panews(date C.Optional_String, limit C.Optional_Int) C.Result_List_NewsItem {
	goDate, goLimit := __parse_input(date, limit)
	newsItems, err := __get_panews_by_date(goDate, goLimit)
	if err != nil {
		return C.err_List_NewsItem(C.CString(err.Error()))
	}
	return __convert_success_result(newsItems)
}

//export get_panews_release
func get_panews_release(res C.Result_List_NewsItem) {
	C.release_Result_List_NewsItem(res)
}

////export get_blockbeats_news_by_date2json
//func get_blockbeats_news_by_date2json(date C.Optional_String, limit C.Optional_Int) C.NewsResultJson {
//
//	goDate, goLimit := __parse_input(date, limit)
//
//	newsItems, err := __get_blockbeats_news_by_date(goDate, goLimit)
//	if err != nil {
//		return C.NewsResultJson{
//			success: C.Bool(false),
//		}
//	}
//	contents, _ := json.Marshal(newsItems)
//	// fmt.Printf("Content: %+v\n", C.NewsResultJson{
//	// 	success:     C.Bool(true),
//	// 	json_string: C.CString(string(contents)),
//	// })
//
//	return C.NewsResultJson{
//		success:     C.Bool(true),
//		json_string: C.CString(string(contents)),
//		//json_string: C.CString(string(contents)),
//	}
//}
//
////export get_blockbeats_news_by_date2json_release
//func get_blockbeats_news_by_date2json_release(res C.NewsResultJson) {
//	C.release_NewsResultJson(res)
//}
//
////export get_panews_by_date2json
//func get_panews_by_date2json(date C.Optional_String, limit C.Optional_Int) C.NewsResultJson {
//
//	goDate, goLimit := __parse_input(date, limit)
//
//	newsItems, err := __get_panews_by_date(goDate, goLimit)
//	if err != nil {
//		fmt.Printf("got error: %+v\n", err)
//		return C.NewsResultJson{
//			success: C.Bool(false),
//		}
//	}
//	contents, _ := json.Marshal(newsItems)
//	fmt.Printf("Content: %+v\n", C.NewsResultJson{
//		success:     C.Bool(true),
//		json_string: C.CString(string(contents)),
//	})
//
//	return C.NewsResultJson{
//		success:     C.Bool(true),
//		json_string: C.CString(string(contents)),
//		//json_string: C.CString(string(contents)),
//	}
//}
//
////export get_panews_by_date2json_release
//func get_panews_by_date2json_release(res C.NewsResultJson) {
//	fmt.Println("get_panews_by_date2json_release")
//	C.release_NewsResultJson(res)
//}

func main() {}

func __parse_input(date C.Optional_String, limit C.Optional_Int) (string, int) {
	goDate := time.Now().Format(time.DateOnly)
	if bool(date.is_some) {
		goDate = C.GoString(date.value)
	}

	goLimit := 5
	if bool(limit.is_some) {
		goLimit = int(limit.value)
	}

	//fmt.Printf("__parse_inputs %s %d\n", goDate, goLimit)
	return goDate, goLimit
}

func __convert_success_result(items []*Custom_FlashNewsItem) C.Result_List_NewsItem {
	newsItems := C.new_List_NewsItem(C.size_t(len(items)))
	itemsSlice := (*[1 << 30]C.NewsItem)(unsafe.Pointer(newsItems.values))
	for i, item := range items {
		addTime := time.Unix(item.AddTime, 0).Format(time.RFC3339)
		itemsSlice[i] = C.NewsItem{
			title:     C.CString(item.Title),
			content:   C.CString(item.Content),
			timestamp: C.CString(addTime),
		}
	}
	return C.ok_List_NewsItem(newsItems)
}
