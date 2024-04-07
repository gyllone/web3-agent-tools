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
	long value;
} OptionalInt;

typedef struct {
	char* json_string;
} BlockBeatsFlashNewsJsonResult;
*/
import "C"
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"unsafe"
)


//export get_blockbeats_news_by_date2json
func get_blockbeats_news_by_date2json(cDate *C.OptionalStr, cLimit *C.OptionalInt) *C.BlockBeatsFlashNewsJsonResult {
    date := C.GoString(cDate.value)
    fmt.Printf("get_blockbeats_news_by_date2json %s %d", date, int(cLimit.value))
    newsItems, err := __get_blockbeats_news_by_date(date, int(cLimit.value))
    if err != nil {
        return nil
    }
    contents, _ := json.Marshal(newsItems)
    fmt.Printf("Content: %s", string(contents))
    result := (*C.BlockBeatsFlashNewsJsonResult)(C.malloc(C.sizeof_BlockBeatsFlashNewsJsonResult))
    result.json_string = C.CString(string(contents))
    return result
}

//export get_blockbeats_news_by_date2json_release
func get_blockbeats_news_by_date2json_release(res *C.BlockBeatsFlashNewsJsonResult) {
	C.free(unsafe.Pointer(res))
}

// e.g. '2024-01-01'
func __get_blockbeats_news_by_date(date string, limit int) ([]*FlashNewsItem, error) {

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	startTs, endTs := int(t.Unix()), int(t.Unix())+86400
	end := false
	items := make([]*Item, 0, 0)
	curTs := endTs
	for !end {
		rsp, err := request(curTs)
		if err != nil {
			return nil, err
		}
		for i, item := range rsp.Data.List {
			if item.AddTime > startTs {
				items = append(items, &rsp.Data.List[i])
				curTs = min(item.AddTime, curTs)
			} else {
				end = true
			}
		}

		if limit > 0 && len(items) >= limit {
			items = items[:limit]
			end = true
		}
	}

	res := make([]*FlashNewsItem, 0, len(items))
	for _, item := range items {
		res = append(res, &FlashNewsItem{
			ArticleID: item.ArticleID,
			Title:     item.Title,
			Content:   item.Content,
			AddTime:   item.AddTime,
			Url:       item.URL,
		})
	}
	return res, nil
}

// page参数没有用, end_time开集合，取时间<end_time的条数
func request(end_time int) (*Response, error) {
	url := fmt.Sprintf("https://api.theblockbeats.info/v1/newsflash/list?page=1&limit=50&ios=-2&end_time=%d&detective=-2", end_time)
	fmt.Printf("request: %s\n", url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("lang", "cn")
	req.Header.Add("origin", "https://www.theblockbeats.info")
	req.Header.Add("referer", "https://www.theblockbeats.info/")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("token", "")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rsp := &Response{}
	if err := json.Unmarshal(body, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}


func main() {
}