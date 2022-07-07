package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func AllKindsOfNodeDepths(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	return AllKindsOfNodeDepths(root.Left) + AllKindsOfNodeDepths(root.Right) + nodeDepths(root, 0)
}

func nodeDepths(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}
	return depth + nodeDepths(node.Left, depth+1) + nodeDepths(node.Right, depth+1)
}
