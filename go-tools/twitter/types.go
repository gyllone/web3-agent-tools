package main

import (
	"fmt"
	"net/http"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

type UserInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type User struct {
	User UserInfo `json:"User"`
}

type UserIdResp map[string]User

// timelines struct
//type Domain struct {
//	ID          string `json:"id"`
//	Name        string `json:"name"`
//	Description string `json:"description"`
//}
//
//type Entity struct {
//	ID          string `json:"id"`
//	Name        string `json:"name"`
//	Description string `json:"description"`
//}

//type ContextAnnotation struct {
//	Domain Domain `json:"Domain"`
//	Entity Entity `json:"Entity"`
//}
//
//type PublicMetrics struct {
//	ImpressionCount   int `json:"impression_count"`
//	URLLinkClicks     int `json:"url_link_clicks"`
//	UserProfileClicks int `json:"user_profile_clicks"`
//	LikeCount         int `json:"like_count"`
//	ReplyCount        int `json:"reply_count"`
//	RetweetCount      int `json:"retweet_count"`
//	QuoteCount        int `json:"quote_count"`
//}

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	//AuthorID           string              `json:"author_id"`
	//ContextAnnotations []ContextAnnotation `json:"context_annotations"`
	//ConversationID     string              `json:"conversation_id"`
	CreatedAt string `json:"created_at"`
	//PublicMetrics      PublicMetrics       `json:"public_metrics"`
}

type TimelineData struct {
	Tweet Tweet `json:"Tweet"`
	//Author UserInfo `json:"Author"`
}

type TimelineResp map[string]TimelineData
