package main

/*
#cgo CFLAGS: -I../dependencies
#include <twitter.h>
*/
import "C"
import (
	"context"
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"strings"
	"sync"
	"unsafe"
)

//export query_users_timeline
func query_users_timeline(usernames C.String, max_results_per_user C.Int) C.Result_Dict_Result_List_TimelineInfo {
	names := strings.Split(C.GoString(usernames), ",")
	userIds, err := getUserIds(names)

	switch {
	case err != nil:
		return C.err_Dict_Result_List_TimelineInfo(C.CString(err.Error()))
	case len(userIds) == 0:
		return C.err_Dict_Result_List_TimelineInfo(C.CString("No user found"))
	}

	data := C.new_Dict_Result_List_TimelineInfo(C.size_t(len(userIds)))
	dataKeyArr := (*[1 << 30]C.String)(unsafe.Pointer(data.keys))[:data.len:data.len]
	dataValueArr := (*[1 << 30]C.Result_List_TimelineInfo)(unsafe.Pointer(data.values))[:data.len:data.len]

	var wg sync.WaitGroup
	wg.Add(len(userIds))

	for i, userId := range userIds {
		go func(i int, userId string) {
			defer wg.Done()
			resp := getUserTimeline(userId, int(max_results_per_user))
			dataKeyArr[i] = C.CString(names[i])
			dataValueArr[i] = resp
			return
		}(i, userId)
	}

	wg.Wait()

	return C.ok_Dict_Result_List_TimelineInfo(data)
}

//export query_users_timeline_release
func query_users_timeline_release(result C.Result_Dict_Result_List_TimelineInfo) {
	C.release_Result_Dict_Result_List_TimelineInfo(result)
}

func getUserTimeline(userId string, maxResultsPerUser int) C.Result_List_TimelineInfo {
	opts := twitter.UserTweetTimelineOpts{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldAuthorID, twitter.TweetFieldConversationID, twitter.TweetFieldPublicMetrics, twitter.TweetFieldContextAnnotations},
		UserFields:  []twitter.UserField{twitter.UserFieldUserName},
		Expansions:  []twitter.Expansion{twitter.ExpansionAuthorID},
		MaxResults:  maxResultsPerUser,
	}

	timeline, err := client.UserTweetTimeline(context.Background(), userId, opts)
	if err != nil {
		return C.err_List_TimelineInfo(C.CString(fmt.Sprintf("User Tweet timeline error: %v", err)))
	}

	dictionaries := timeline.Raw.TweetDictionaries()

	data := C.new_List_TimelineInfo(C.size_t(len(dictionaries)))
	if data.len == 0 {
		return C.ok_List_TimelineInfo(data)
	}

	dataArr := (*[1 << 30]C.TimelineInfo)(unsafe.Pointer(data.values))[:data.len:data.len]

	postIdx := 0
	for _, post := range dictionaries {
		dataArr[postIdx] = C.TimelineInfo{
			id:         C.CString(post.Tweet.ID),
			text:       C.CString(post.Tweet.Text),
			created_at: C.CString(post.Tweet.CreatedAt),
		}
		postIdx++
	}

	return C.ok_List_TimelineInfo(data)
}

func getUserIds(usernames []string) ([]string, error) {
	ids := make([]string, 0)
	opts := twitter.UserLookupOpts{
		Expansions: []twitter.Expansion{twitter.ExpansionPinnedTweetID},
	}

	userResponse, err := client.UserNameLookup(context.Background(), usernames, opts)
	if err != nil {
		return ids, err
	}

	dictionaries := userResponse.Raw.UserDictionaries()

	for _, v := range dictionaries {
		ids = append(ids, v.User.ID)
	}

	return ids, nil
}

func main() {}
