package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	sortedNodeValues := make([]int, 0)
	inOrderTraverse(tree, &sortedNodeValues)
	return sortedNodeValues[len(sortedNodeValues)-k]
}

func inOrderTraverse(node *BST, sortedNodeValues *[]int) {
	if node == nil {
		return
	}

	inOrderTraverse(node.Left, sortedNodeValues)
	*sortedNodeValues = append(*sortedNodeValues, node.Value)
	inOrderTraverse(node.Right, sortedNodeValues)
}
