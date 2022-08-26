package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type TreeInfo struct {
	diameter int
	height   int
}

func BinaryTreeDiameter(tree *BinaryTree) int {
	return getTreeInfo(tree).diameter
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{0, 0}
	}

	leftTree := getTreeInfo(tree.Left)
	rightTree := getTreeInfo(tree.Right)

	longestPath := leftTree.height + rightTree.height
	maxDiameter := max(leftTree.diameter, rightTree.diameter)
	currentDiameter := max(longestPath, maxDiameter)
	currentHeight := 1 + max(leftTree.height, rightTree.height)

	return TreeInfo{currentDiameter, currentHeight}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
