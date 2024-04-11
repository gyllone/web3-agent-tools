package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func __get_panews_by_date(date string, limit int) ([]*Custom_FlashNewsItem, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	startTs, endTs := int(t.Unix()), int(t.Unix())+86400
	end := false
	items := make([]*PA_News, 0, 0)
	curTs := endTs
	for !end {
		rsp, err := request_panews(curTs)
		if err != nil {
			return nil, err
		}
		for i, item := range rsp.Data.FlashNews[0].List {
			if int(item.PublishTime) > startTs {
				items = append(items, &rsp.Data.FlashNews[0].List[i])
				if int(item.PublishTime) < curTs {
					curTs = int(item.PublishTime)
				}
			} else {
				end = true
			}
		}

		if limit > 0 && len(items) >= limit {
			items = items[:limit]
			end = true
		}
	}

	res := make([]*Custom_FlashNewsItem, 0, len(items))
	for _, item := range items {
		res = append(res, &Custom_FlashNewsItem{
			ArticleID: item.ID,
			Title:     item.Title,
			Content:   item.Desc,
			AddTime:   int(item.PublishTime),
			Url:       "",
		})
	}
	return res, nil
}

func request_panews(end_time int) (*PA_Response, error) {
	url := fmt.Sprintf("https://www.panewslab.com/webapi/flashnews?LId=1&LastTime=%d&Rn=50&tw=0", end_time)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cookie", "_ga=GA1.1.2014066543.1712110885; _ga_KHBYDL8DMV=GS1.1.1712457989.2.1.1712458122.0.0.0")
	req.Header.Add("referer", "https://www.panewslab.com/zh/news/index.html")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
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

	rsp := &PA_Response{}
	if err := json.Unmarshal(body, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
