package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{1, 2, 4, 8}
	// 1, 1, 4, 5, 6, 8, 9
	coins2 := []int{5, 7, 1, 1, 2, 3, 22}
	fmt.Println(NonConstructibleChangeSolution2(coins))
	fmt.Println(NonConstructibleChangeSolution2(coins2))
}

func NonConstructibleChangeSolution2(coins []int) int {
	sort.Ints(coins)

	currentChangeCreated := 0
	for _, coin := range coins {
		if coin > currentChangeCreated+1 {
			return currentChangeCreated + 1
		}
		currentChangeCreated += coin
	}
	return currentChangeCreated + 1
}

func NonConstructibleChangeNotFinishYet(coins []int) int {
	minSum := 1
	sort.Ints(coins)

	for true {
		foundChanges := findChanges(coins, minSum)
		if !foundChanges {
			break
		}
		minSum++
	}

	return minSum
}

func findChanges(coins []int, target int) bool {
	maxSum := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] > target {
			break
		}
		maxSum += coins[i]
		if maxSum > target {
			break
		} else if maxSum == target {
			return true
		}
	}
	return false
}
