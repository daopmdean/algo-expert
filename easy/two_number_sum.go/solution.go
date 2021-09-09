package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	fmt.Println(TwoNumberSumSolution1(array, target))
	fmt.Println(TwoNumberSumSolution2(array, target))
}

func TwoNumberSumSolution1(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return []int{}
}

func TwoNumberSumSolution2(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentValue := array[left] + array[right]
		if currentValue == target {
			return []int{array[left], array[right]}
		} else if currentValue > target {
			right -= 1
		} else {
			left += 1
		}
	}
	return []int{}
}
