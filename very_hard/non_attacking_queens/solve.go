package main

func NonAttackingQueens(n int) int {
	columnPlacements := make([]int, n)
	return getNumberOfNonAttackingQueenPlacements(0, columnPlacements, n)
}

func getNumberOfNonAttackingQueenPlacements(row int, columnPlacements []int, boardSize int) int {
	if row == boardSize {
		return 1
	}

	validPlacements := 0
	for col := 0; col < boardSize; col++ {
		if isNonAttackingPlacement(row, col, columnPlacements) {
			columnPlacements[row] = col
			validPlacements += getNumberOfNonAttackingQueenPlacements(row+1, columnPlacements, boardSize)
		}
	}
	return validPlacements
}

func isNonAttackingPlacement(row, col int, columnPlacements []int) bool {
	for previousRow := 0; previousRow < row; previousRow++ {
		columnToCheck := columnPlacements[previousRow]
		sameColumn := columnToCheck == col
		onDiagonal := abs(columnToCheck-col) == row-previousRow
		if sameColumn || onDiagonal {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
