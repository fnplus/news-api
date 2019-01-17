package datastore

// Client is an interface to interact with the firebase firestore
type Client interface {
	init() error
	GetKeywordPool() ([]string, error)
	AddNewsItems(newsItems []NewsItem) error
}

// NewsItem is a modle of the news item in the database
type NewsItem struct {
	Title   string `json:"title"`
	Source  string `json:"source"`
	Author  string `json:"author"`
	URL     string `json:"url"`
	Content string `json:"content"`
}

// NewClient creates a new instance of the client
func NewClient() Client {
	var firestoreClient *FirestoreClient
	firestoreClient.init()
	return firestoreClient
}
