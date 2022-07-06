package main

// This is the class of the input root. Do not edit it.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func FlattenBinaryTree(root *BinaryTree) *BinaryTree {
	leftMost, _ := flattenTree(root)
	return leftMost
}

func flattenTree(node *BinaryTree) (leftMost, rightMost *BinaryTree) {
	leftMost = node
	if node.Left != nil {
		leftSubtreeLeftMost, leftSubtreeRightMost := flattenTree(node.Left)
		connectNodes(leftSubtreeRightMost, node)
		leftMost = leftSubtreeLeftMost
	}

	rightMost = node
	if node.Right != nil {
		rightSubtreeLeftMost, rightSubtreeRightMost := flattenTree(node.Right)
		connectNodes(node, rightSubtreeLeftMost)
		rightMost = rightSubtreeRightMost
	}

	return leftMost, rightMost
}

func connectNodes(left, right *BinaryTree) {
	left.Right = right
	right.Left = left
}
