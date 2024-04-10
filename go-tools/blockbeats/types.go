package main

// BlockBeats types
// BB_Response 结构体对应整个 JSON 对象
type BB_Response struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data BB_Data `json:"data"`
}

// BB_Data 结构体对应 JSON 对象中的 data 字段
type BB_Data struct {
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
	List  []BB_Item `json:"list"`
}

// BB_Item 结构体对应 data.list 中的对象
type BB_Item struct {
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

// PANews types
type PA_Response struct {
	Errno int     `json:"errno"`
	Msg   string  `json:"msg"`
	Data  BB_Data `json:"data"`
}

type PA_Data struct {
	HotArticles interface{}    `json:"hotArticles"` // 或者根据实际情况定义更具体的类型
	FlashNews   []PA_FlashNews `json:"flashNews"`
}

type PA_FlashNews struct {
	Date  string    `json:"date"`
	Week  string    `json:"week"`
	Month string    `json:"month"`
	Unix  int64     `json:"unix"`
	List  []PA_News `json:"list"`
}

type PA_News struct {
	ID          string        `json:"id"`
	Type        int           `json:"type"`
	PublishTime int64         `json:"publishTime"`
	Img         string        `json:"img"`
	Title       string        `json:"title"`
	Desc        string        `json:"desc"`
	Readnum     int           `json:"readnum"`
	Tags        []string      `json:"tags"` // 注意这里，原始JSON中"tags"是null，但通常我们会定义为字符串数组
	Author      PA_NewsAuthor `json:"author"`
}

type PA_NewsAuthor struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Img    string `json:"img"`
	Tag    string `json:"tag"`
	Follow int    `json:"follow"`
}

type Custom_FlashNewsItem struct {
	ArticleID int
	Title     string
	Content   string
	AddTime   int
	Url       string
}
