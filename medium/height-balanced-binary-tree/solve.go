package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type treeInfo struct {
	isBalanced bool
	height     int
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	treeInfo := getTreeInfo(tree)
	return treeInfo.isBalanced
}

func getTreeInfo(node *BinaryTree) treeInfo {
	if node == nil {
		return treeInfo{isBalanced: true, height: -1}
	}

	leftSubtreeInfo := getTreeInfo(node.Left)
	rightSubtreeInfo := getTreeInfo(node.Right)

	cond := abs(leftSubtreeInfo.height-rightSubtreeInfo.height) <= 1
	isBalanced := leftSubtreeInfo.isBalanced && rightSubtreeInfo.isBalanced && cond
	height := max(leftSubtreeInfo.height, rightSubtreeInfo.height) + 1

	return treeInfo{
		isBalanced: isBalanced,
		height:     height,
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
