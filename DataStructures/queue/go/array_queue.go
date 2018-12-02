package main

import (
	"fmt"
)

// ArrayQueue is a quaue implemented with an array
type ArrayQueue struct {
	size  int
	count int
	list  *[]int
	first int
	last  int
}

// Init creates queue
func (q *ArrayQueue) Init(size int) {
	q.size = size
	newQueue := make([]int, size, size)
	q.list = &newQueue
	q.first = 0
	q.last = 0
}

func (q *ArrayQueue) checkInit() {
	if q.list == nil {
		panic("Have not initialized ArrayQueue using 'init' function")
	}
}

func (q *ArrayQueue) nextIndex(i int) int {
	return (i + 1) % q.size
}

// Enqueue adds item to queue
func (q *ArrayQueue) Enqueue(value int) {
	q.checkInit()
	if q.count == q.size {
		panic("Queue has reached capacity")
	}
	nextIndex := q.nextIndex(q.last)
	queue := *q.list
	queue[q.last] = value
	q.count++
	q.last = nextIndex
}

// Dequeue removes item from queue
func (q *ArrayQueue) Dequeue() int {
	q.checkInit()
	if q.count == 0 {
		panic("Queue is empty")
	}
	queue := *q.list
	item := queue[q.first]

	q.first = q.nextIndex(q.first)
	q.count--
	return item
}

// Print items in queue
func (q *ArrayQueue) Print() {
	fmt.Println("Printing items in queue")
	start := q.first
	count := q.count
	queue := *q.list
	for count != 0 {
		fmt.Print(queue[start])
		fmt.Print(" -> ")
		start = q.nextIndex(start)
		count--
	}
	fmt.Print("END")
}
