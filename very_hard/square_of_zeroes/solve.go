package main

func SquareOfZeroes(matrix [][]int) bool {
	n := len(matrix)
	for topRow := 0; topRow < n; topRow++ {
		for leftCol := 0; leftCol < n; leftCol++ {
			squareLength := 2
			for squareLength <= n-leftCol && squareLength <= n-topRow {
				bottomRow := topRow + squareLength - 1
				rightCol := leftCol + squareLength - 1
				if isSquare(matrix, topRow, leftCol, bottomRow, rightCol) {
					return true
				}
				squareLength++
			}
		}
	}
	return false
}

func isSquare(matrix [][]int, r1, c1, r2, c2 int) bool {
	for row := r1; row < r2+1; row++ {
		if matrix[row][c1] != 0 || matrix[row][c2] != 0 {
			return false
		}
	}

	for col := c1; col < c2+1; col++ {
		if matrix[r1][col] != 0 || matrix[r2][col] != 0 {
			return false
		}
	}

	return true
}
