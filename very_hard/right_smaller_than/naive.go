package main

func RightSmallerThan(array []int) []int {
	result := []int{}

	for i, v := range array {
		count := 0
		for j := i + 1; j < len(array); j++ {
			if v > array[j] {
				count++
			}
		}
		result = append(result, count)
	}

	return result
}
