package main

import "fmt"

func main() {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, -1}
	fmt.Println(IsValidSubsequenceSolution1(arr, seq))
	fmt.Println(IsValidSubsequenceSolution2(arr, seq))
	arr2 := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq2 := []int{22, 25, 6}
	fmt.Println(IsValidSubsequenceSolution1(arr2, seq2))
	fmt.Println(IsValidSubsequenceSolution2(arr2, seq2))
}

func IsValidSubsequenceSolution1(array []int, sequence []int) bool {
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

func IsValidSubsequenceSolution2(array []int, sequence []int) bool {
	seqIndex := 0
	for _, val := range array {
		if seqIndex == len(sequence) {
			return true
		}
		if sequence[seqIndex] == val {
			seqIndex++
		}
	}
	return seqIndex == len(sequence)
}
