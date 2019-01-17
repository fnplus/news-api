package main

import "fmt"

type node struct {
	value string
	next  *node
}

// Queue to perform FIFO
type Queue struct {
	front *node
}

// NewQueue creates a new instance of a queue
func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) push(str string) {
	n := node{value: str, next: nil}
	if q.front == nil {
		q.front = &n
	} else {
		q.front.next = &n
	}
}

func (q *Queue) pop() (string, error) {
	if q.front == nil {
		return "", fmt.Errorf("The queue is empty")
	}
	value := q.front.value
	q.front = q.front.next
	return value, nil
}

func (q *Queue) isEmpty() bool {
	return q.front == nil
}
