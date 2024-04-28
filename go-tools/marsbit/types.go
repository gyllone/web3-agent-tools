package main

type news struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Synopsis string `json:"synopsis"`
}

type newsData struct {
	NewsType    string `json:"infoType"`
	News        news   `json:"infoObj"`
	PublishTime int64  `json:"publishTime"`
}

type NewsResp struct {
	Message string     `json:"msg"`
	Data    []newsData `json:"obj"`
}

type searchedNews struct {
	news
	PublishTime int64 `json:"publishTime"`
}

type searchedNewsData struct {
	InforList []searchedNews `json:"inforList"`
}

type searchedNewsObj struct {
	News          searchedNewsData `json:"News"`
	Lives         searchedNewsData `json:"Lives"`
	ExcellentNews searchedNewsData `json:"ExcellentNews"`
}

type SearchedNewsResp struct {
	Message string          `json:"msg"`
	Data    searchedNewsObj `json:"obj"`
}
