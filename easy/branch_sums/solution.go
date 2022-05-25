package easy

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	result := []int{}

	traverseTree(root, 0, &result)

	return result
}

func traverseTree(node *BinaryTree, currentSum int, result *[]int) {
	if node == nil {
		return
	}

	currentSum += node.Value

	if node.Left == nil && node.Right == nil {
		*result = append(*result, currentSum)
	}

	traverseTree(node.Left, currentSum, result)
	traverseTree(node.Right, currentSum, result)
}
