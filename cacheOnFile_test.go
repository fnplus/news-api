package main

import (
	"testing"
)

func Test_cacheNews(t *testing.T) {
	cacheNews("tesgij;oidfg;ohg")
}

func Test_newsInCache(t *testing.T) {
	newsInCache("test")
}
