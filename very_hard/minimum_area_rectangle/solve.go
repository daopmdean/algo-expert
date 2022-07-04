package main

import (
	"fmt"
	"math"
)

func MinimumAreaRectangle(points [][]int) int {
	pointSet := createPointSet(points)
	var minimumAreaFound = math.MaxInt32

	for currentIdx := range points {
		p2x, p2y := points[currentIdx][0], points[currentIdx][1]
		for previousIdx := 0; previousIdx < currentIdx; previousIdx++ {
			p1x, p1y := points[previousIdx][0], points[previousIdx][1]
			pointsShareValue := p1x == p2x || p1y == p2y

			if pointsShareValue {
				continue
			}

			point1OnOppositeDiagonalExists := pointSet[convertPointToString(p1x, p2y)]
			point2OnOppositeDiagonalExists := pointSet[convertPointToString(p2x, p1y)]
			oppositeDiagonalExists := point1OnOppositeDiagonalExists && point2OnOppositeDiagonalExists

			if oppositeDiagonalExists {
				currentArea := abs(p2x-p1x) * abs(p2y-p1y)
				minimumAreaFound = min(minimumAreaFound, currentArea)
			}
		}
	}

	if minimumAreaFound != math.MaxInt32 {
		return minimumAreaFound
	}

	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func createPointSet(points [][]int) map[string]bool {
	pointSet := map[string]bool{}

	for _, point := range points {
		x, y := point[0], point[1]
		pointString := convertPointToString(x, y)
		pointSet[pointString] = true
	}
	return pointSet
}

func convertPointToString(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}
