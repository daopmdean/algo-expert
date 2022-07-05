package main

import "fmt"

func LineThroughPoints(points [][]int) int {
	maxNumberOfPointsOnLine := 1

	for idx1, p1 := range points {
		slopes := map[string]int{}
		for idx2 := idx1 + 1; idx2 < len(points); idx2++ {
			p2 := points[idx2]
			rise, run := getSlopeOfLineBetweenPoints(p1, p2)
			slopeKey := createHashTableKeyForRational(rise, run)
			if slopes[slopeKey] == 0 {
				slopes[slopeKey] = 1
			}
			slopes[slopeKey] += 1
		}

		currentMaxNumberOfPointsOnLine := maxSlope(slopes)
		maxNumberOfPointsOnLine = max(maxNumberOfPointsOnLine, currentMaxNumberOfPointsOnLine)
	}

	return maxNumberOfPointsOnLine
}

func getSlopeOfLineBetweenPoints(p1, p2 []int) (int, int) {
	p1x, p1y := p1[0], p1[1]
	p2x, p2y := p2[0], p2[1]

	if p1x == p2x {
		return 1, 0
	}

	var xDiff = p1x - p2x
	var yDiff = p1y - p2y
	gcd := getGreatestCommonDivisor(abs(xDiff), abs(yDiff))
	xDiff = xDiff / gcd
	yDiff = yDiff / gcd
	if xDiff < 0 {
		xDiff *= -1
		yDiff *= -1
	}

	return yDiff, xDiff
}

func getGreatestCommonDivisor(num1, num2 int) int {
	for {
		if num1 == 0 {
			return num2
		}
		if num2 == 0 {
			return num1
		}
		num1, num2 = num2, num1%num2
	}
}

func createHashTableKeyForRational(numerator int, denominator int) string {
	return fmt.Sprintf("%d:&d", numerator, denominator)
}

func maxSlope(slopes map[string]int) int {
	currentMax := 0
	for _, slope := range slopes {
		currentMax = max(slope, currentMax)
	}
	return currentMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
