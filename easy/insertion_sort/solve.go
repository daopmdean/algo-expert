package main

func InsertionSort(array []int) []int {
	for i, v := range array {
		if i == 0 {
			continue
		}

		if v >= array[i-1] {
			continue
		}

		currentIndex := i
		for j := i - 1; j >= 0; j-- {
			if v < array[j] {
				array = swap(array, currentIndex, j)
				currentIndex = j
			} else {
				break
			}
		}
	}
	return array
}

func swap(array []int, i, j int) []int {
	array[i], array[j] = array[j], array[i]
	return array
}
