package main

import (
	"fmt"
)

func main() {
	s := Stack{}
	s.Push(11)
	fmt.Println(s.Pop())

	s.Push(12)
	s.Push(42)
	fmt.Println(s.Pop())
}
