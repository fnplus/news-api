package main

import (
	"testing"

	"github.com/fnplus/community-news-bot/datastore"
)

func TestWorker_GetNewsWithKeywords(t *testing.T) {
	apikey := "3111edae78f6492983bd0a6df945356e"
	worker := NewWorker(apikey)
	_, err := worker.GetNewsWithKeywords("golang")
	if err != nil {
		t.Fatalf("Test failed!\nError: %s", err.Error())
	}
}

func TestWorker_PushToDB(t *testing.T) {
	worker := NewWorker("3111edae78f6492983bd0a6df945356e")
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
