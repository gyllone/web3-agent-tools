package main

import (
	"testing"

	query "github.com/google/go-querystring/query"
)

func TestTypes(t *testing.T) {
	req := &QuoteKlineRequest{}
	values, _ := query.Values(req)

	t.Log(values.Encode())
	t.Log(req)
}

// func TestBlockBeats(t *testing.T) {
// 	today := time.Now().Format(time.DateOnly)
// 	t.Log(__get_blockbeats_news_by_date(today, 10))
// }
