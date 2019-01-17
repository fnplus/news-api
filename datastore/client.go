package datastore

import (
	"github.com/barthr/newsapi"
)

// Client is an interface to interact with the firebase firestore
type Client interface {
	init() error
	GetKeywordPool() ([]string, error)
	AddNewsItems(newsItems []NewsItem) error
}

// NewsItem is a modle of the news item in the database
type NewsItem struct {
	Item    newsapi.Article `json:"item"`
	Keyword string          `json:"keyword"`
}

// NewClient creates a new instance of the client
func NewClient() Client {
	var firestoreClient = FirestoreClient{nil}
	firestoreClient.init()
	return &firestoreClient
}
