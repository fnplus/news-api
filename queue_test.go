package main

import (
	"fmt"
	"testing"
)

func TestQueue_push(t *testing.T) {
	q := Queue{}
	q.push("test-push")
	q.push("test-push")
	fmt.Println(q.front)
}

func TestQueue_pop(t *testing.T) {
	var q Queue
	q.push("sample-push")
	fmt.Println(q.front)
	q.pop()
	fmt.Println(q.front)
}
