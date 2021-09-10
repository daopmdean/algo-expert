package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 5, 6, 8, 9}
	arr = SortedSquaredArray(arr)
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
}

func SortedSquaredArray(array []int) []int {
	for i, val := range array {
		array[i] *= val
	}
	sort.Ints(array)
	return array
}
