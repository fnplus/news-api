package curator

import (
	"context"

	"github.com/barthr/newsapi"
)

func getNewsWithKeywords(keywords string) ([]newsapi.Article, error) {
	client := newsapi.NewClient("3111edae78f6492983bd0a6df945356e")
	ar, err := client.GetEverything(context.Background(), &newsapi.EverythingParameters{
		Keywords: keywords,
		Language: "en",
	})
	if err != nil {
		return nil, err
	}
	return ar.Articles, nil
}
