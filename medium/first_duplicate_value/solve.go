package main

func FirstDuplicateValue(array []int) int {
	minIndex := -1
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				if minIndex == -1 {
					minIndex = j
				} else if j < minIndex {
					minIndex = j
				}
			}
		}
	}
	if minIndex == -1 {
		return minIndex
	}

	return array[minIndex]
}
