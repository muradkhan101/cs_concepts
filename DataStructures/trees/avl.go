package main

// AVLNode is a node in a self-balanced AVL tree
type AVLNode struct {
	Value   int
	Left    *AVLNode
	Right   *AVLNode
	Parent  *AVLNode
	balance int
}

func (node *AVLNode) hasLeftChild() bool {
	return node.Left != nil
}
func (node *AVLNode) hasRightChild() bool {
	return node.Right != nil
}
func (node *AVLNode) hasParent() bool {
	return node.Parent != nil
}
func (node *AVLNode) getChildCount() int {
	count := 0
	if node.Left != nil {
		count++
	}
	if node.Right != nil {
		count++
	}
	return count
}

// Add node to AVL tree
func (node *AVLNode) Add(child *AVLNode) {
	if child.Value < node.Value {
		if node.hasLeftChild() {
			node.Left.Add(child)
		} else {
			node.Left = child
			node.Left.updateBalance()
		}
	} else {
		if node.hasRightChild() {
			node.Right.Add(child)
		} else {
			node.Right = child
			node.Right.updateBalance()
		}
	}
}

func (node *AVLNode) Remove(value int) bool {
	if value == node.Value {

		switch childCount := node.getChildCount(); childCount {
		case 0:
			if node.Parent.Left == node {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
			node.Parent.updateBalance()
		case 1:
			if node.Parent.Left == node {
				if node.hasRightChild() {
					node.Parent.Left = node.Right
					node.Right.Parent = node.Parent
				} else {
					node.Parent.Left = node.Left
					node.Left.Parent = node.Parent
				}
				node.Parent.updateBalance()
			}
		case 2:
			minNode := node.Right
			for minNode.Left != nil {
				minNode = minNode.Left
			}
			node.Value = minNode.Value
			node.Right.Remove(value)

		}
	} else if value < node.Value {
		if node.hasLeftChild() {
			return node.Left.Remove(value)
		} else {
			return false
		}
	} else {
		if node.hasRightChild() {
			return node.Right.Remove(value)
		} else {
			return false
		}
	}
	return false
}
func (node *AVLNode) updateBalance() {
	if node.balance > 1 || node.balance < -1 {
		node.rebalance()
		return
	}
	if node.hasParent() {
		if node.Parent.Left == node {
			node.Parent.balance++
		} else if node.Parent.Right == node {
			node.Parent.balance--
		}

		if node.Parent.balance != 0 {
			node.Parent.updateBalance()
		}
	}
}

func (node *AVLNode) rotateLeft() {
	var newRoot = node.Right
	newRoot.Parent = node.Parent
	if node.hasParent() {
		if node.Parent.Left == node {
			node.Parent.Left = newRoot
		} else if node.Parent.Right == node {
			node.Parent.Right = newRoot
		}
	}
	node.Parent = newRoot
	node.Right = newRoot.Left
	if newRoot.hasLeftChild() {
		newRoot.Left.Parent = node
	}
	newRoot.Left = node

	node.balance = node.balance + 1 - min(newRoot.balance, 0)
	newRoot.balance = newRoot.balance + 1 + max(node.balance, 0)
}

func (node *AVLNode) rotateRight() {
	var newRoot = node.Left
	newRoot.Parent = node.Parent
	if node.hasParent() {
		if node.Parent.Left == node {
			node.Parent.Left = newRoot
		} else if node.Parent.Right == node {
			node.Parent.Right = newRoot
		}
	}
	node.Parent = newRoot
	node.Left = newRoot.Right
	if newRoot.hasRightChild() {
		newRoot.Right.Parent = node
	}
	newRoot.Right = node

	node.balance = node.balance + 1 - min(newRoot.balance, 0)
	newRoot.balance = newRoot.balance + 1 + max(node.balance, 0)
}

func (node *AVLNode) rebalance() {
	if node.balance < 0 {
		if node.Right.balance > 0 {
			node.Right.rotateRight()
		}
		node.rotateLeft()
	} else if node.balance > 0 {
		if node.Left.balance < 0 {
			node.Left.rotateLeft()
		}
		node.rotateRight()
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
