package easy

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	result := 0
	currentWaiting := 0

	for i := 0; i < len(queries); i++ {
		if i != 0 {
			currentWaiting += queries[i-1]
		}
		result += currentWaiting
	}

	return result
}
