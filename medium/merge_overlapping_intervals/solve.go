package main

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	intervals = sort(intervals)
	cursor := 0

	for cursor < len(intervals)-1 {
		if intervals[cursor+1][0] <= intervals[cursor][1] {
			intervals[cursor][1] = max(intervals[cursor][1], intervals[cursor+1][1])
			intervals = remove(intervals, cursor+1)
		} else {
			cursor++
		}
	}

	return intervals
}

func sort(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	return intervals
}

func remove(intervals [][]int, i int) [][]int {
	return append(intervals[:i], intervals[i+1:]...)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
