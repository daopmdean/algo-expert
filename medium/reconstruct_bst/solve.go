package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	tree := BST{
		Value: preOrderTraversalValues[0],
	}

	for i, v := range preOrderTraversalValues {
		if i == 0 {
			continue
		}
		tree.insert(v)
	}

	return &tree
}

func (t *BST) insert(value int) {
	if t.Value > value {
		if t.Left == nil {
			t.Left = &BST{
				Value: value,
			}
		} else {
			t.Left.insert(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &BST{
				Value: value,
			}
		} else {
			t.Right.insert(value)
		}
	}
}
