package main

import (
	"fmt"
)

func main() {
	var n = &BSTNode{5, nil, nil}
	n.add(&BSTNode{3, nil, nil})
	n.add(&BSTNode{7, nil, nil})
	fmt.Printf("Before delete: %+v", n)
	fmt.Print("\n")
	n.remove(3, nil)
	fmt.Printf("After delete: %+v", n)
	fmt.Print("\n")
}
