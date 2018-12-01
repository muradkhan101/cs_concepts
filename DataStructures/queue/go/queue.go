package main

import (
	"fmt"
)

// Node holds data for queue
type Node struct {
	Value int
	Next  *Node
}

// Queue is a queue type
type Queue struct {
	first *Node
	last  *Node
}

// Enqueue adds item to queue
func (q *Queue) Enqueue(val int) {
	if q.first == nil {
		n := &Node{val, nil}
		q.first = n
		q.last = n
	} else {
		q.last.Next = &Node{val, nil}
		q.last = q.last.Next
	}
}

// Dequeue removes item from queue
func (q *Queue) Dequeue() int {
	if q.first == nil {
		panic("Nothing in the queue!")
	}
	val := q.first.Value
	q.first = q.first.Next
	if q.first == nil {
		q.last = nil
	}
	return val
}

// Print prints items in queue
func (q *Queue) Print() {
	fmt.Println("Printing Queue!")
	node := q.first
	for node != nil {
		fmt.Print(node.Value)
		fmt.Print(" -> ")
		node = node.Next
	}
	fmt.Println("END")
}
