package datastore

import (
	"context"
	"strings"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

// FirestoreClient implements Client
type FirestoreClient struct {
	app *firebase.App
}

func (fc *FirestoreClient) init() error {
	opt := option.WithCredentialsFile("firebase-creds.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	fc.app = app
	return nil
}

// GetKeywordPool gets the pool of keywords from the database
func (fc *FirestoreClient) GetKeywordPool() ([]string, error) {
	var keywords []string
	cl, err := fc.app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}
	defer cl.Close()

	snap, err := cl.Collection("community-news").Doc("keyword-pool").Get(context.Background())
	if err != nil {
		return nil, err
	}

	keywordSnap, err := snap.DataAt("keywords")
	if err != nil {
		return nil, err
	}

	for _, word := range keywordSnap.([]interface{}) {
		keywords = append(keywords, strings.TrimSpace(word.(string)))
	}
	return keywords, nil
}

// AddNewsItems add a list of news items to the database
func (fc *FirestoreClient) AddNewsItems(newsItems []NewsItem) error {
	cl, err := fc.app.Firestore(context.Background())
	if err != nil {
		return err
	}
	defer cl.Close()

	ref := cl.Collection("community-news/news/items")
	for _, nItem := range newsItems {
		ref.Add(context.Background(), nItem)
	}
	return nil
}
