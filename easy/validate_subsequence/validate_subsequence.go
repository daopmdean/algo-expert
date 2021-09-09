package main

import "fmt"

func main() {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, -1}
	fmt.Println(IsValidSubsequence(arr, seq))
}

func IsValidSubsequence(array []int, sequence []int) bool {
	for i := 0; i < len(sequence); i++ {
		found := false
		for j := 0; j < len(array); j++ {
			if sequence[i] == array[j] {
				found = true
				array = array[j+1:]
			}
		}
		if !found {
			return false
		}
	}
	return true
}
