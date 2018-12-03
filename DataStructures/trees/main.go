package main

import (
	"fmt"
)

func main() {
	var n = &BSTNode{5, nil, nil}
	n.Add(&BSTNode{3, nil, nil})
	n.Add(&BSTNode{11, nil, nil})
	n.Add(&BSTNode{19, nil, nil})
	n.Add(&BSTNode{7, nil, nil})
	n.Add(&BSTNode{14, nil, nil})
	n.Add(&BSTNode{8, nil, nil})
	InOrder(n)
	fmt.Println("Post remove----")
	n.Remove(11, nil)
	InOrder(n)
}
