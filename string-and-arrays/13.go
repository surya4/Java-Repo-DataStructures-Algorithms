package main

import (
	"fmt"
)

func imageRotate(matrix [][]int, n int) [][]int {
	// for i := range s {
	// 	s = strings.Replace(s, " ", "%20", i)
	// }

	return matrix
}

func main() {
	input1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	input2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(input1, " <-- input,  output --> ", imageRotate(input1), 4)
	fmt.Println(input2, " <-- input,  output --> ", imageRotate(input2), 3)
}
