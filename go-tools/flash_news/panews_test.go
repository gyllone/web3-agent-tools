package main

import (
	"testing"
	"time"
)

func TestRequestPaNews(t *testing.T) {
	t.Log(request_panews(1712741238))
}

func Test__get_panews_by_date(t *testing.T) {
	today := time.Now().Format(time.DateOnly)
	t.Log(__get_panews_by_date(today, 10))
}
