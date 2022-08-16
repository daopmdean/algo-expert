package main

import "math"

func SameBsts(arrayOne, arrayTwo []int) bool {
	return areSameBsts(arrayOne, arrayTwo, 0, 0, math.MinInt32, math.MaxInt32)
}

func areSameBsts(arrayOne, arrayTwo []int, rootIdxOne, rootIdxTwo int, minVal, maxVal int) bool {
	if rootIdxOne == -1 || rootIdxTwo == -1 {
		return rootIdxOne == rootIdxTwo
	}

	if arrayOne[rootIdxOne] != arrayTwo[rootIdxTwo] {
		return false
	}

	leftRootIdxOne := getIdxOfFirstSmaller(arrayOne, rootIdxOne, minVal)
	leftRootIdxTwo := getIdxOfFirstSmaller(arrayTwo, rootIdxTwo, minVal)
	rightRootIdxOne := getIdxOfFirstBiggerOrEqual(arrayOne, rootIdxOne, maxVal)
	rightRootIdxTwo := getIdxOfFirstBiggerOrEqual(arrayTwo, rootIdxTwo, maxVal)

	currentValue := arrayOne[rootIdxOne]
	leftAreSame := areSameBsts(arrayOne, arrayTwo, leftRootIdxOne, leftRootIdxTwo, minVal, currentValue)
	rightAreSame := areSameBsts(arrayOne, arrayTwo, rightRootIdxOne, rightRootIdxTwo, currentValue, maxVal)

	return leftAreSame && rightAreSame
}

func getIdxOfFirstSmaller(array []int, startingIdx, minVal int) int {
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] < array[startingIdx] && array[i] >= minVal {
			return i
		}
	}
	return -1
}

func getIdxOfFirstBiggerOrEqual(array []int, startingIdx, maxVal int) int {
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] >= array[startingIdx] && array[i] < maxVal {
			return i
		}
	}
	return -1
}
