package main

type news struct {
	ID             int64  `json:"id"`
	IsTop          int64  `json:"is_top"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Cover          string `json:"cover"`
	NewsURL        string `json:"news_url"`
	ExtractionTags string `json:"extraction_tags"`
	UpdatedAt      string `json:"updated_at"`
}

type newsData struct {
	Items []news `json:"items"`
}

type NewsResp struct {
	Data    newsData `json:"data"`
	Message string   `json:"message"`
}

type Post struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	UpdatedAt string `json:"updated_at"`
}

type postData struct {
	Items []Post `json:"items"`
}

type PostResp struct {
	Data    postData `json:"data"`
	Message string   `json:"msg"`
}
