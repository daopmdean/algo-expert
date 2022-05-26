package easy

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	sum := 0

	traverse(root, 0, &sum)

	return sum
}

func traverse(node *BinaryTree, currentDepth int, sum *int) {
	if node == nil {
		return
	}

	*sum += currentDepth
	currentDepth++

	traverse(node.Left, currentDepth, sum)
	traverse(node.Right, currentDepth, sum)
}
