package main

import (
	"fmt"
)

func main() {
	// s := ArrayStack{}
	// s.Push(11)
	// fmt.Println(s.Pop())

	// s.Push(12)
	// s.Push(42)
	// fmt.Println(s.Pop())

	// s.Push(1)
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	q := QueueWithStacks{}
	q.Enqueue(4)
	q.Enqueue(3)
	q.Enqueue(11)
	q.Enqueue(22)
	fmt.Print("Dequeing: ")
	fmt.Println(q.Dequeue())
	fmt.Print("Dequeing: ")
	fmt.Println(q.Dequeue())
}
