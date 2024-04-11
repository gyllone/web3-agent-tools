package main

import (
	"testing"
	"time"
)

func TestBlockBeats(t *testing.T) {
	today := time.Now().Format(time.DateOnly)
	t.Log(__get_blockbeats_news_by_date(today, 10))
}
