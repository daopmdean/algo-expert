package main

func MoveElementToEnd(array []int, toMove int) []int {
	lastIdx := len(array) - 1

	for i := 0; i < lastIdx; i++ {
		if array[i] == toMove {
			moveElement(array, i, &lastIdx)
		}
	}

	return array
}

func moveElement(array []int, idx int, lastIdx *int) {
	for {
		if array[*lastIdx] == array[idx] {
			if idx < *lastIdx {
				*lastIdx--
			} else {
				break
			}
		} else {
			array[idx], array[*lastIdx] = array[*lastIdx], array[idx]
			break
		}
	}
}
