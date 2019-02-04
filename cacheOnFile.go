package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var cache = []string{}

func fetchCache() {
	conts, err := ioutil.ReadFile(".newscache")
	if err != nil {
		fmt.Println("newsInCache: ", err.Error())
	}

	for _, title := range strings.Split(string(conts), "\n") {
		cache = append(cache, title)
	}
}

func newsInCache(title string) bool {
	title = strings.TrimSpace(title)
	title = strings.ToLower(title)
	for _, cTitle := range cache {
		cTitle = strings.TrimSpace(cTitle)
		cTitle = strings.ToLower(cTitle)
		if cTitle == title {
			return true
		}
	}
	return false
}

func cacheNews(title string) {
	err := ioutil.WriteFile(".newscache", []byte(strings.ToLower(title)), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
