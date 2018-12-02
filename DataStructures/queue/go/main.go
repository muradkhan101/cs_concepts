package main

import (
	"fmt"
)

func main() {
	q := ArrayQueue{}
	q.Init(4)
	q.Enqueue(4)
	q.Enqueue(3)
	q.Enqueue(11)
	q.Enqueue(22)
	q.Print()
	fmt.Print("Dequeing: ")
	fmt.Println(q.Dequeue())
	fmt.Print("Dequeing: ")
	fmt.Println(q.Dequeue())
	q.Print()
}
