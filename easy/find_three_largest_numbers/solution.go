package easy

func FindThreeLargestNumbers(array []int) []int {
	result := []int{array[0], array[1], array[2]}
	reorderArray(result)

	for i, val := range array {
		if i <= 2 {
		} else {
			if val > result[0] {
				if val < result[1] {
					result[0] = val
				} else if val < result[2] {
					result[0] = result[1]
					result[1] = val
				} else {
					result[0] = result[1]
					result[1] = result[2]
					result[2] = val
				}
			}
		}
	}

	return result
}

func reorderArray(result []int) {
	if result[0] > result[1] {
		result[0], result[1] = result[1], result[0]
	}
	if result[1] > result[2] {
		result[1], result[2] = result[2], result[1]
	}
	if result[0] > result[1] {
		result[0], result[1] = result[1], result[0]
	}
}
