package main

func LongestPeak(array []int) int {
	longest := 0
	i := 1
	for i < len(array)-1 {
		isPeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isPeak {
			i++
			continue
		}

		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
			leftIdx--
		}

		rightIdx := i + 2
		for rightIdx < len(array) && array[rightIdx] < array[rightIdx-1] {
			rightIdx++
		}

		currentPeakLength := rightIdx - leftIdx - 1
		if currentPeakLength > longest {
			longest = currentPeakLength
		}
		i = rightIdx
	}
	return longest
}
