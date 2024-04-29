package main

/*
#cgo CFLAGS: -I../dependencies
#include <twitter.h>
*/
import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"strings"
	"unsafe"
)

//export query_users_by_username
func query_users_by_username(username C.String) C.Result_List_User {
	name := C.GoString(username)
	names := strings.Split(name, ",")

	opts := twitter.UserLookupOpts{
		Expansions: []twitter.Expansion{twitter.ExpansionPinnedTweetID},
	}

	userResponse, err := client.UserNameLookup(context.Background(), names, opts)
	if err != nil {
		return C.err_List_User(C.CString(err.Error()))
	}

	dictionaries := userResponse.Raw.UserDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		return C.err_List_User(C.CString(err.Error()))
	}

	var resp UserIdResp
	err = json.Unmarshal(enc, &resp)
	if err != nil {
		return C.err_List_User(C.CString(err.Error()))
	}

	data := C.new_List_User(C.size_t(len(resp)))
	if data.len == 0 {
		return C.ok_List_User(data)
	}

	dataArr := (*[1 << 30]C.User)(unsafe.Pointer(data.values))[:data.len:data.len]

	idx := 0
	for _, user := range resp {
		dataArr[idx] = C.User{
			id:       C.CString(user.User.ID),
			name:     C.CString(user.User.Name),
			username: C.CString(user.User.Username),
		}
		idx++
	}

	return C.ok_List_User(data)
}

//export query_users_by_username_release
func query_users_by_username_release(result C.Result_List_User) {
	C.release_Result_List_User(result)
}

//export query_user_timeline
func query_user_timeline(id C.String) C.Result_List_TimelineInfo {
	userID := C.GoString(id)

	opts := twitter.UserTweetTimelineOpts{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldAuthorID, twitter.TweetFieldConversationID, twitter.TweetFieldPublicMetrics, twitter.TweetFieldContextAnnotations},
		UserFields:  []twitter.UserField{twitter.UserFieldUserName},
		Expansions:  []twitter.Expansion{twitter.ExpansionAuthorID},
		MaxResults:  5,
	}

	timeline, err := client.UserTweetTimeline(context.Background(), userID, opts)
	if err != nil {
		return C.err_List_TimelineInfo(C.CString(fmt.Sprintf("User Tweet timeline error: %v", err)))
	}

	dictionaries := timeline.Raw.TweetDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		return C.err_List_TimelineInfo(C.CString(err.Error()))
	}

	var resp TimelineResp

	if err = json.Unmarshal(enc, &resp); err != nil {
		return C.err_List_TimelineInfo(C.CString(err.Error()))
	}

	data := C.new_List_TimelineInfo(C.size_t(len(resp)))
	if data.len == 0 {
		return C.ok_List_TimelineInfo(data)
	}

	dataArr := (*[1 << 30]C.TimelineInfo)(unsafe.Pointer(data.values))[:data.len:data.len]

	postIdx := 0
	for _, post := range resp {
		dataArr[postIdx] = C.TimelineInfo{
			tweet: C.Tweet{
				id:                  C.CString(post.Tweet.ID),
				text:                C.CString(post.Tweet.Text),
				context_annotations: getCList(post.Tweet.ContextAnnotations),
				created_at:          C.CString(post.Tweet.CreatedAt),
			},
			author: C.User{
				id:       C.CString(post.Author.ID),
				name:     C.CString(post.Author.Name),
				username: C.CString(post.Author.Username),
			},
		}
		postIdx++
	}

	return C.ok_List_TimelineInfo(data)
}

//export query_user_timeline_release
func query_user_timeline_release(result C.Result_List_TimelineInfo) {
	C.release_Result_List_TimelineInfo(result)
}

//export query_user_mention_timeline
func query_user_mention_timeline(userId C.String) C.Result_List_TimelineInfo {
	userID := C.GoString(userId)

	opts := twitter.UserMentionTimelineOpts{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldAuthorID, twitter.TweetFieldConversationID, twitter.TweetFieldPublicMetrics, twitter.TweetFieldContextAnnotations},
		UserFields:  []twitter.UserField{twitter.UserFieldUserName},
		Expansions:  []twitter.Expansion{twitter.ExpansionAuthorID},
		MaxResults:  5,
	}

	timeline, err := client.UserMentionTimeline(context.Background(), userID, opts)
	if err != nil {
		return C.err_List_TimelineInfo(C.CString(fmt.Sprintf("User Tweet timeline error: %v", err)))
	}

	dictionaries := timeline.Raw.TweetDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		return C.err_List_TimelineInfo(C.CString(err.Error()))
	}
	var resp TimelineResp

	if err = json.Unmarshal(enc, &resp); err != nil {
		return C.err_List_TimelineInfo(C.CString(err.Error()))
	}

	data := C.new_List_TimelineInfo(C.size_t(len(resp)))
	if data.len == 0 {
		return C.ok_List_TimelineInfo(data)
	}

	dataArr := (*[1 << 30]C.TimelineInfo)(unsafe.Pointer(data.values))[:data.len:data.len]

	postIdx := 0
	for _, post := range resp {
		dataArr[postIdx] = C.TimelineInfo{
			tweet: C.Tweet{
				id:                  C.CString(post.Tweet.ID),
				text:                C.CString(post.Tweet.Text),
				context_annotations: getCList(post.Tweet.ContextAnnotations),
				created_at:          C.CString(post.Tweet.CreatedAt),
			},
			author: C.User{
				id:       C.CString(post.Author.ID),
				name:     C.CString(post.Author.Name),
				username: C.CString(post.Author.Username),
			},
		}
		postIdx++
	}

	return C.ok_List_TimelineInfo(data)
}

//export query_user_mention_timeline_release
func query_user_mention_timeline_release(result C.Result_List_TimelineInfo) {
	C.release_Result_List_TimelineInfo(result)
}

func getCList(annotations []ContextAnnotation) C.List_ContextAnnotation {
	res := C.new_List_ContextAnnotation(C.size_t(len(annotations)))
	if res.len == 0 {
		return res
	}

	dataArr := (*[1 << 30]C.ContextAnnotation)(unsafe.Pointer(res.values))[:res.len:res.len]

	for idx, annotation := range annotations {
		dataArr[idx] = C.ContextAnnotation{
			domain: C.Domain{
				id:          C.CString(annotation.Domain.ID),
				name:        C.CString(annotation.Domain.Name),
				description: C.CString(annotation.Domain.Description),
			},
			entity: C.Entity{
				id:          C.CString(annotation.Entity.ID),
				name:        C.CString(annotation.Entity.Name),
				description: C.CString(annotation.Entity.Description),
			},
		}
	}

	return res
}

func main() {

}
