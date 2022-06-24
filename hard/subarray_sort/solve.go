package main

import "math"

func SubarraySort(array []int) []int {
	minOutOfOrder, maxOutOfOrder := math.MaxInt32, math.MinInt32
	for i, num := range array {
		if isOutOfOrder(i, num, array) {
			minOutOfOrder = min(minOutOfOrder, num)
			maxOutOfOrder = max(maxOutOfOrder, num)
		}
	}
	if minOutOfOrder == math.MaxInt32 {
		return []int{-1, -1}
	}

	subArrayLeft := 0
	for minOutOfOrder >= array[subArrayLeft] {
		subArrayLeft += 1
	}

	subArrayRight := len(array) - 1
	for maxOutOfOrder <= array[subArrayRight] {
		subArrayRight -= 1
	}

	return []int{subArrayLeft, subArrayRight}
}

func isOutOfOrder(i, num int, array []int) bool {
	if i == 0 {
		return num > array[i+1]
	}
	if i == len(array)-1 {
		return num < array[i-1]
	}
	return num > array[i+1] || num < array[i-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
