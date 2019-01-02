package main

import "github.com/fnplus/api-news-curator/curator"

func main() {
	worker := curator.NewWorker("3111edae78f6492983bd0a6df945356e") // Change API key
	runWorker(worker)
}

func runWorker(worker curator.IWorker) {
}
