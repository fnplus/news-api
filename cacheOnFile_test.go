package main

import (
	"fmt"
	"testing"
)

func Test_cacheNews(t *testing.T) {
	cacheNews("tesgij;oidfg;ohg")
}

func Test_newsInCache(t *testing.T) {
	cache = []string{"tes", "Test"}
	fmt.Println(newsInCache("testhh"))
}
