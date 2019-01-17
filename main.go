package main

import (
	"log"
	"strings"
	"sync"

	"github.com/fnplus/community-news-bot/datastore"
)

var keywordPool []string
var keywordQueue *Queue
var wg sync.WaitGroup

const maxNumberOfRoutines = 3

func main() {
	storeClient := datastore.NewClient()
	log.Println("Fetching keywords pool")
	keywordPool, err := storeClient.GetKeywordPool()
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	log.Println("Queuing keywords")
	keywordQueue = NewQueue()
	for _, word := range keywordPool {
		keywordQueue.push(strings.TrimSpace(word))
	}

	for i := 0; i < maxNumberOfRoutines; i++ {
		worker := NewWorker("3111edae78f6492983bd0a6df945356e")
		wg.Add(1)
		go runWorker(worker)
	}
	wg.Wait()
}

func runWorker(worker IWorker) {
	defer wg.Done()
	for !keywordQueue.isEmpty() {
		var finalNewsItems []datastore.NewsItem
		word, _ := keywordQueue.pop()

		log.Println("Fetching news items for \"", word, "\"")
		newsItems, err := worker.GetNewsWithKeywords(word)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		} else {
			for _, n := range newsItems {
				finalNewsItems = append(finalNewsItems, datastore.NewsItem{
					Item: n, Keyword: word,
				})
			}
		}

		log.Printf("Pushing %d news items to database", len(finalNewsItems))
		err = worker.PushToDB(finalNewsItems)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
	}
}
