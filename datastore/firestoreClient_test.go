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
