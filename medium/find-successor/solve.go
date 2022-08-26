package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	inOrderTraversal := []*BinaryTree{}
	getInOrderTraversal(tree, &inOrderTraversal)

	for idx, currentNode := range inOrderTraversal {
		if currentNode != node {
			continue
		}

		if idx == len(inOrderTraversal)-1 {
			return nil
		}
		return inOrderTraversal[idx+1]
	}
	return nil
}

func getInOrderTraversal(node *BinaryTree, order *[]*BinaryTree) {
	if node == nil {
		return
	}

	getInOrderTraversal(node.Left, order)
	*order = append(*order, node)
	getInOrderTraversal(node.Right, order)
}
