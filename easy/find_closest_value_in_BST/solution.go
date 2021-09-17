package easy

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	closestValue := tree.Value
	valueRange := diffAbs(tree.Value, target)

	for tree.Left != nil || tree.Right != nil {
		if target == tree.Value {
			return target
		} else if target > tree.Value && tree.Right != nil {
			tree = tree.Right
			if diffAbs(tree.Value, target) < valueRange {
				closestValue = tree.Value
				valueRange = diffAbs(tree.Value, target)
			}
		} else if target < tree.Value && tree.Left != nil {
			tree = tree.Left
			if diffAbs(tree.Value, target) < valueRange {
				closestValue = tree.Value
				valueRange = diffAbs(tree.Value, target)
			}
		} else {
			break
		}
	}

	return closestValue
}

func diffAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
