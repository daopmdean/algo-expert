package main

func SolveSudoku(board [][]int) [][]int {
	solvePartialSudoku(0, 0, board)
	return board
}

func solvePartialSudoku(row, col int, board [][]int) bool {
	var currentRow = row
	var currentCol = col

	if currentCol == len(board[currentRow]) {
		currentRow++
		currentCol = 0
		if currentRow == len(board) {
			return true
		}
	}

	if board[currentRow][currentCol] == 0 {
		return tryDigitsAtPosition(currentRow, currentCol, board)
	}

	return solvePartialSudoku(currentRow, currentCol+1, board)
}

func tryDigitsAtPosition(row, col int, board [][]int) bool {
	for digit := 1; digit < 10; digit++ {
		if isValidAtPosition(digit, row, col, board) {
			board[row][col] = digit
			if solvePartialSudoku(row, col+1, board) {
				return true
			}
		}
	}

	board[row][col] = 0
	return false
}

func isValidAtPosition(value, row, col int, board [][]int) bool {
	rowIsValid := !rowContains(board, row, value)
	columnIsValid := !colContains(board, col, value)

	if !rowIsValid || !columnIsValid {
		return false
	}

	subgridRowStart := (row / 3) * 3
	subgridColStart := (col / 3) * 3
	for rowIdx := 0; rowIdx < 3; rowIdx++ {
		for colIdx := 0; colIdx < 3; colIdx++ {
			rowToCheck := subgridRowStart + rowIdx
			colToCheck := subgridColStart + colIdx
			existingValue := board[rowToCheck][colToCheck]

			if existingValue == value {
				return false
			}
		}
	}

	return true
}

func rowContains(board [][]int, row, value int) bool {
	for _, ele := range board[row] {
		if value == ele {
			return true
		}
	}
	return false
}

func colContains(board [][]int, col, value int) bool {
	for _, row := range board {
		if value == row[col] {
			return true
		}
	}
	return false
}
