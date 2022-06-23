package main

func LargestRange(array []int) []int {
	best := []int{}
	longestLength := 0
	nums := map[int]bool{}

	for _, num := range array {
		nums[num] = true
	}

	for _, num := range array {
		if !nums[num] {
			continue
		}

		nums[num] = false
		currentLength, left, right := 1, num-1, num+1
		for nums[left] {
			currentLength += 1
			nums[left] = false
			left -= 1
		}

		for nums[right] {
			currentLength += 1
			nums[right] = false
			right += 1
		}

		if currentLength > longestLength {
			longestLength = currentLength
			best = []int{left + 1, right - 1}
		}
	}

	return best
}
