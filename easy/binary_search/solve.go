package main

func BinarySearch(array []int, target int) int {
	start := 0
	end := len(array) - 1

	for start <= end {
		i := start + (end-start)/2
		if array[i] == target {
			return i
		}
		if array[i] > target {
			end = i - 1
		}
		if array[i] < target {
			start = i + 1
		}
	}

	return -1
}

// 0 1 21 33 45
