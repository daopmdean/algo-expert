package main

import "math"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	_, max := findMaxSum(tree)
	return max
}

func findMaxSum(tree *BinaryTree) (int, int) {
	if tree == nil {
		return 0, math.MinInt32
	}

	leftMaxSumAsBranch, leftMaxPathSum := findMaxSum(tree.Left)
	rightMaxSumAsBranch, rightMaxPathSum := findMaxSum(tree.Right)
	maxChildSumAsBranch := max(leftMaxSumAsBranch, rightMaxSumAsBranch)

	value := tree.Value
	maxSumAsBranch := max(maxChildSumAsBranch+value, value)
	maxSumAsRootNode := max(leftMaxSumAsBranch+value+rightMaxSumAsBranch, maxSumAsBranch)
	maxPathSum := max(leftMaxPathSum, rightMaxPathSum, maxSumAsRootNode)

	return maxSumAsBranch, maxPathSum
}

func max(first int, vals ...int) int {
	for _, val := range vals {
		if val > first {
			first = val
		}
	}
	return first
}
