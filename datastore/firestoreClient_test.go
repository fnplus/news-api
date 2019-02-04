package datastore

import (
	"fmt"
	"testing"

	"firebase.google.com/go"
)

func TestFirestoreClient_init(t *testing.T) {
	type fields struct {
		app *firebase.App
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Test 1",
			fields:  fields{app: nil},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &FirestoreClient{
				app: tt.fields.app,
			}
			if err := fc.init(); (err != nil) != tt.wantErr {
				t.Errorf("FirestoreClient.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFirestoreClient_GetKeywordPool(t *testing.T) {
	client := FirestoreClient{}
	client.init()
	pool, err := client.GetKeywordPool()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	fmt.Println(pool)
}

func TestFirestoreClient_AddNewsItems(t *testing.T) {
	cl := FirestoreClient{}
	cl.init()

	newsitems := []NewsItem{
		NewsItem{Keyword: "sample-title-1"},
		NewsItem{Keyword: "sample-title-2"},
		NewsItem{Keyword: "sample-title-3"},
	}

	err := cl.AddNewsItems(newsitems)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestFirestoreClient_CacheNewsTitle(t *testing.T) {
	cl := FirestoreClient{}
	cl.init()

	err := cl.CacheNewsTitle([]string{"test1"})
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}
