package main

func NumberOfBinaryTreeTopologies(n int) int {
	cache := []int{1}

	for i := 1; i < n+1; i++ {
		numberOfTrees := 0
		for leftTreeSize := 0; leftTreeSize < i; leftTreeSize++ {
			rightTreeSize := i - 1 - leftTreeSize
			numberOfLeftTrees := cache[leftTreeSize]
			numberOfRightTrees := cache[rightTreeSize]
			numberOfTrees += numberOfLeftTrees * numberOfRightTrees
		}
		cache = append(cache, numberOfTrees)
	}

	return cache[n]
}
