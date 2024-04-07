package main

// Response 结构体对应整个 JSON 对象
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

// Data 结构体对应 JSON 对象中的 data 字段
type Data struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	List  []Item `json:"list"`
}

// Item 结构体对应 data.list 中的对象
type Item struct {
	ID               int      `json:"id"`
	ArticleID        int      `json:"article_id"`
	ContentID        int      `json:"content_id"`
	Type             int      `json:"type"`
	IsShowHome       int      `json:"is_show_home"`
	IsDetective      int      `json:"is_detective"`
	IsTop            int      `json:"is_top"`
	IsOriginal       int      `json:"is_original"`
	SpecialID        int      `json:"special_id"`
	TopicID          int      `json:"topic_id"`
	IOS              int      `json:"ios"`
	IsFirst          int      `json:"is_first"`
	IsHot            int      `json:"is_hot"`
	AddTime          int      `json:"add_time"`
	ImgURL           string   `json:"img_url"`
	URL              string   `json:"url"`
	CryptoToken      string   `json:"crypto_token"`
	Title            string   `json:"title"`
	Lang             string   `json:"lang"`
	PID              int      `json:"p_id"`
	Abstract         *string  `json:"abstract"` // 使用指针是因为 abstract 可能为 null
	Content          string   `json:"content"`
	TagList          []string `json:"tag_list"`
	CollectionStatus int      `json:"collection_status"`
}

type FlashNewsItem struct {
	ArticleID int
	Title     string
	Content   string
	AddTime   int
	Url       string
}
