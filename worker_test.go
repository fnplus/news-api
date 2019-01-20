package main

import (
	"testing"

	"github.com/fnplus/community-news-bot/datastore"
)

var sConfig = getConfig()

func TestWorker_GetNewsWithKeywords(t *testing.T) {
	worker := NewWorker(sConfig.NewsAPIToken)
	_, err := worker.GetNewsWithKeywords("golang")
	if err != nil {
		t.Fatalf("Test failed!\nError: %s", err.Error())
	}
}

func TestWorker_PushToDB(t *testing.T) {
	worker := NewWorker(sConfig.NewsAPIToken)
	newsitems := []datastore.NewsItem{
		datastore.NewsItem{Keyword: "sample-title-1"},
		datastore.NewsItem{Keyword: "sample-title-2"},
		datastore.NewsItem{Keyword: "sample-title-3"},
	}
	err := worker.PushToDB(newsitems)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}
