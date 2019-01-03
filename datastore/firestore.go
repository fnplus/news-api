package datastore

// Client is an interface to interact with the firebase firestore
type Client interface {
}

// NewClient creates a new instance of the client
func NewClient() Client {
	return &FirestoreClient{}
}

// FirestoreClient implements Client
type FirestoreClient struct{}
