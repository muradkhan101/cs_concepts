package main

import "fmt"

func main() {
	var n = &BSTNode{Value: 5}
	n.Add(&BSTNode{Value: 3})
	n.Add(&BSTNode{Value: 11})
	node2 := &BSTNode{Value: 19}
	n.Add(node2)
	node1 := &BSTNode{Value: 7}
	n.Add(node1)

	n.Add(&BSTNode{Value: 14})
	n.Add(&BSTNode{Value: 8})

	fmt.Println(lowestAncestor(n, node1, node2))
}
