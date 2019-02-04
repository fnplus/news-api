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
	for _, cTitle := range cache {
		if strings.TrimSpace(cTitle) == strings.TrimSpace(title) {
			fmt.Println(title, "\n----\n", cTitle)
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
