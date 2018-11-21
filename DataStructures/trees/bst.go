package main

// BSTNode is a Binary Search Tree node
type BSTNode struct {
	Value int
	// Parent *BSTNode
	Left  *BSTNode
	Right *BSTNode
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
func (node *BSTNode) add(child *BSTNode) {
	if child.Value < node.Value {
		if node.hasLeftChild() {
			node.Left.add(child)
		} else {
			// child.Parent = node
			node.Left = child
		}
	} else {
		if node.hasRightChild() {
			node.Right.add(child)
		} else {
			// child.Parent = node
			node.Right = child
		}
	}
}

func (node *BSTNode) remove(child int, parent *BSTNode) bool {
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
			minNode := node.Left
			for minNode.hasLeftChild() {
				minNode = minNode.Left
			}
			node.Value = minNode.Value
			return node.Left.remove(minNode.Value, node)
		}
	} else if child < node.Value {
		if node.hasLeftChild() {
			return node.Left.remove(child, node)
		} else {
			return false
		}
	} else {
		if node.hasRightChild() {
			return node.Right.remove(child, node)
		} else {
			return false
		}
	}
}
