package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/fnplus/community-news-bot/datastore"
)

var keywordPool []string
var keywordQueue *Queue
var wg sync.WaitGroup

func main() {
	sConfig := getConfig()
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

	for i := 0; i < sConfig.NumberOfWorkers; i++ {
		worker := NewWorker(sConfig.NewsAPIToken)
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

type config struct {
	NumberOfWorkers int    `yaml:"numberOfWorkers"`
	NewsAPIToken    string `yaml:"newsAPIToken"`
}

func getConfig() *config {
	contents, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Not able to read the config file. Error: %s", err.Error())
	}

	var sConfig config
	err = yaml.Unmarshal(contents, &sConfig)
	if err != nil {
		log.Fatalf("Not able to parse the config file. Error: %s", err.Error())
	}
	return &sConfig
}
