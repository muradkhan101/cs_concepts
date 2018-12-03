package main

import "fmt"

// BSTNode is a Binary Search Tree node
type BSTNode struct {
	Value int
	// Parent *BSTNode
	Left  *BSTNode
	Right *BSTNode
}

// InOrder traverses tree
func InOrder(t *BSTNode) {
	if t.Left != nil {
		InOrder(t.Left)
	}
	fmt.Println(t.Value)
	if t.Right != nil {
		InOrder(t.Right)
	}
}

func (node *BSTNode) hasLeftChild() bool {
	return node.Left != nil
}
func (node *BSTNode) hasRightChild() bool {
	return node.Right != nil
}
func (node *BSTNode) getChildCount() int {
	count := 0
	if node.Left != nil {
		count++
	}
	if node.Right != nil {
		count++
	}
	return count
}

// Add nodo to BST
func (node *BSTNode) Add(child *BSTNode) {
	if child.Value < node.Value {
		if node.hasLeftChild() {
			node.Left.Add(child)
		} else {
			// child.Parent = node
			node.Left = child
		}
	} else {
		if node.hasRightChild() {
			node.Right.Add(child)
		} else {
			// child.Parent = node
			node.Right = child
		}
	}
}

// Remove Node from BST
func (node *BSTNode) Remove(child int, parent *BSTNode) bool {
	if node.Value == child {
		childCount := node.getChildCount()
		if childCount == 0 {
			if parent.Left == node {
				parent.Left = nil
			} else if parent.Right == node {
				parent.Right = nil
			}
			return true
		} else if childCount == 1 {
			var replacement *BSTNode
			if node.Left != nil {
				replacement = node.Left
			} else {
				replacement = node.Right
			}
			if parent.Left == node {
				parent.Left = replacement
			} else if parent.Right == node {
				parent.Right = replacement
			}
			return true
		} else {
			minNode := node.Right
			for minNode.hasLeftChild() {
				minNode = minNode.Left
			}
			node.Value = minNode.Value
			return node.Right.Remove(minNode.Value, node)
		}
	} else if child < node.Value {
		if node.hasLeftChild() {
			return node.Left.Remove(child, node)
		} else {
			return false
		}
	} else {
		if node.hasRightChild() {
			return node.Right.Remove(child, node)
		} else {
			return false
		}
	}
}
