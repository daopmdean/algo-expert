package main

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	result := [][]int{}
	if len(array) < 3 {
		return result
	}

	sort.Ints(array)

	for i := 0; i < len(array)-2; i++ {
		for j := i + 1; j < len(array)-1; j++ {
			for k := j + 1; k < len(array); k++ {
				if array[i]+array[j]+array[k] == target {
					sumFound := []int{array[i], array[j], array[k]}
					result = append(result, sumFound)
				}
			}
		}
	}

	return result
}
