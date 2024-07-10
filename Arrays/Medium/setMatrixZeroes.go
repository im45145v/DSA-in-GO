package main

import "fmt"

func setZeroes(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])

	zeroRows := make([]bool, rows)
	zeroCols := make([]bool, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = true
				zeroCols[j] = true
			}
		}
	}

	for i := 0; i < rows; i++ {
		if zeroRows[i] {
			for j := 0; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < cols; j++ {
		if zeroCols[j] {
			for i := 0; i < rows; i++ {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	setZeroes(matrix)

	fmt.Println(matrix)
}
