package main

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middleIndex := len(array) / 2
	leftHalf := MergeSort(array[:middleIndex])
	rightHalf := MergeSort(array[middleIndex:])
	return mergeSortedArray(leftHalf, rightHalf)
}

func mergeSortedArray(leftHalf, rightHalf []int) []int {
	sortedArray := make([]int, len(leftHalf)+len(rightHalf))
	k, i, j := 0, 0, 0
	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] <= rightHalf[j] {
			sortedArray[k] = leftHalf[i]
			i++
		} else {
			sortedArray[k] = rightHalf[j]
			j++
		}
		k++
	}

	for i < len(leftHalf) {
		sortedArray[k] = leftHalf[i]
		k++
		i++
	}

	for j < len(rightHalf) {
		sortedArray[k] = rightHalf[j]
		k++
		j++
	}

	return sortedArray
}
