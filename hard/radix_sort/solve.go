package main

func RadixSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	maxNumber := max(array)

	digit := 0
	for maxNumber/pow(10, digit) > 0 {
		countingSort(array, digit)
		digit++
	}

	return array
}

func countingSort(array []int, digit int) {
	sortedArray := make([]int, len(array))
	countArray := make([]int, 10)

	digitColumn := pow(10, digit)
	for _, num := range array {
		countIndex := (num / digitColumn) % 10
		countArray[countIndex] += 1
	}

	for idx := 1; idx < 10; idx++ {
		countArray[idx] += countArray[idx-1]
	}

	for idx := len(array) - 1; idx >= 0; idx-- {
		countIndex := (array[idx] / digitColumn) % 10
		countArray[countIndex] -= 1
		sortedIndex := countArray[countIndex]
		sortedArray[sortedIndex] = array[idx]
	}

	for idx := range array {
		array[idx] = sortedArray[idx]
	}
}

func max(array []int) int {
	max := array[0]
	for _, v := range array {
		if max < v {
			max = v
		}
	}
	return max
}

func pow(a int, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= a
	}
	return result
}
