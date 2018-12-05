package main

func findSuccessor(node *AVLNode) *AVLNode {
	if node.hasRightChild() {
		min := node.Right
		for min.hasLeftChild() {
			min = min.Left
		}
		return min
	}

	if node.hasParent() {
		parent := node.Parent
		if parent.Left == node {
			return parent
		}
		for true {
			if !parent.hasParent() {
				break
			}
			next := parent.Parent
			if next.Left == parent {
				return next
			}
			parent = next
		}
	}
	return nil
}

func lowestAncestor(root, node1, node2 *BSTNode) *BSTNode {
	leftHasNode1 := isChild(root.Left, node1)
	leftHasNode2 := isChild(root.Left, node2)
	if leftHasNode1 && leftHasNode2 {
		return lowestAncestor(root.Left, node1, node2)
	}
	rightHasNode1 := isChild(root.Right, node1)
	rightHasNode2 := isChild(root.Right, node2)
	if rightHasNode1 && rightHasNode2 {
		return lowestAncestor(root.Right, node1, node2)
	}
	return root
}

func isChild(parent, child *BSTNode) bool {
	if parent == nil {
		return false
	}
	if parent == child {
		return true
	}

	return isChild(parent.Left, child) || isChild(parent.Right, child)
}
