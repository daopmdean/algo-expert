package main

import "math"

type Block map[string]bool

func ApartmentHunting(blocks []Block, reqs []string) int {
	minDistancesFromBlocks := [][]int{}
	for _, req := range reqs {
		minDistancesFromBlocks = append(minDistancesFromBlocks, getMinDistances(blocks, req))
	}
	maxDistancesAtBlocks := getMaxDistancesAtBlocks(blocks, minDistancesFromBlocks)

	var optimalBlockIdx int
	smallestMaxDistance := math.MaxInt32
	for i, currentDistance := range maxDistancesAtBlocks {
		if currentDistance < smallestMaxDistance {
			smallestMaxDistance = currentDistance
			optimalBlockIdx = i
		}
	}

	return optimalBlockIdx
}

func getMinDistances(blocks []Block, req string) []int {
	minDistances := make([]int, len(blocks))
	closestReq := math.MaxInt32

	for i := range blocks {
		if val, found := blocks[i][req]; found && val {
			closestReq = i
		}
		minDistances[i] = distanceBetween(i, closestReq)
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if val, found := blocks[i][req]; found && val {
			closestReq = i
		}
		minDistances[i] = min(minDistances[i], distanceBetween(i, closestReq))
	}

	return minDistances
}

func getMaxDistancesAtBlocks(blocks []Block, minDistancesFromBlocks [][]int) []int {
	maxDistancesAtBlocks := make([]int, len(blocks))
	for i := range blocks {
		minDistancesAtBlocks := []int{}
		for _, distances := range minDistancesFromBlocks {
			minDistancesAtBlocks = append(minDistancesAtBlocks, distances[i])
		}
		maxDistancesAtBlocks[i] = max(minDistancesAtBlocks)
	}
	return maxDistancesAtBlocks
}

func distanceBetween(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(array []int) int {
	if len(array) == 0 {
		return 0
	}

	max := array[0]
	for _, val := range array {
		if max < val {
			max = val
		}
	}
	return max
}
