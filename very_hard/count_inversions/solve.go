package main

func CountInversions(array []int) int {
	return countSubArrayInversions(array, 0, len(array))
}

func countSubArrayInversions(array []int, start, end int) int {
	if end-start <= 1 {
		return 0
	}

	middle := start + (end-start)/2
	leftInversions := countSubArrayInversions(array, start, middle)
	rightInversions := countSubArrayInversions(array, middle, end)
	mergedArrayInversions := mergeSortAndCountInversions(array, start, middle, end)
	return leftInversions + rightInversions + mergedArrayInversions
}

func mergeSortAndCountInversions(array []int, start, middle, end int) int {
	sortedArray := make([]int, 0)
	left := start
	right := middle
	inversions := 0

	for left < middle && right < end {
		if array[left] <= array[right] {
			sortedArray = append(sortedArray, array[left])
			left++
		} else {
			inversions += middle - left
			sortedArray = append(sortedArray, array[right])
			right++
		}
	}

	sortedArray = append(sortedArray, array[left:middle]...)
	sortedArray = append(sortedArray, array[right:end]...)
	for idx := range sortedArray {
		num := sortedArray[idx]
		array[start+idx] = num
	}

	return inversions
}
