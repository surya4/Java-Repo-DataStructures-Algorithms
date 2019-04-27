package main

import (
	"fmt"
)

func setZero(matrix [][]int) [][]int {
	row := -1
	col := -1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row = i
				col = j
				break
			}
		}
		if row > -1 {
			break
		}
	}

	if row > -1 {
		for i := 0; i < len(matrix[row]); i++ {
			matrix[row][i] = 0
		}
		for i := 0; i < len(matrix[col]); i++ {
			matrix[i][col] = 0
		}

	}
	return matrix
}

func main() {
	input1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 0}, {13, 14, 15, 16}}
	input2 := [][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}
	fmt.Println(input1, " <-- input,  output --> ", setZero(input1))
	fmt.Println(input2, " <-- input,  output --> ", setZero(input2))
}
