package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func CompareLeafTraversal(tree1 *BinaryTree, tree2 *BinaryTree) bool {
	t1 := []*BinaryTree{tree1}
	t2 := []*BinaryTree{tree2}

	for len(t1) > 0 && len(t2) > 0 {
		t1Leaf := getNextLeafNode(&t1)
		t2Leaf := getNextLeafNode(&t2)

		if t1Leaf.Value != t2Leaf.Value {
			return false
		}
	}

	return len(t1) == 0 && len(t2) == 0
}

func getNextLeafNode(traversalStack *[]*BinaryTree) *BinaryTree {
	var currentNode *BinaryTree
	currentNode, *traversalStack = (*traversalStack)[len(*traversalStack)-1], (*traversalStack)[:len(*traversalStack)-1]

	for !isLeafNode(currentNode) {
		if currentNode.Right != nil {
			*traversalStack = append(*traversalStack, currentNode.Right)
		}

		if currentNode.Left != nil {
			*traversalStack = append(*traversalStack, currentNode.Left)
		}
		currentNode, *traversalStack = (*traversalStack)[len(*traversalStack)-1], (*traversalStack)[:len(*traversalStack)-1]
	}

	return currentNode
}

func isLeafNode(node *BinaryTree) bool {
	return node.Left == nil && node.Right == nil
}
