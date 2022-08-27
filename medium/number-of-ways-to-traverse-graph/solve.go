package main

func NumberOfWaysToTraverseGraph(width int, height int) int {
	return factor(width+height-2) / factor(width-1) / factor(height-1)
}

func factor(n int) int {
	if n <= 1 {
		return 1
	}

	result := 1
	for i := n; i >= 1; i-- {
		result *= i
	}
	return result
}
