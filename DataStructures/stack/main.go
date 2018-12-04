package main

import (
	"fmt"
)

func main() {
	s := MinStack{}
	s.Push(11)
	fmt.Println(s.MinValue())

	s.Push(12)
	s.Push(42)
	fmt.Println(s.MinValue())

	s.Push(1)
	fmt.Println(s.MinValue())
	s.Pop()
	fmt.Println(s.MinValue())
}
