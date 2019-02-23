package main

import (
	"fmt"
	"math"
)

func contigiousSubarraySum(arr [8]int) int {
	sum := math.MinInt64
	tempSum := 0

	for i := 0; i < len(arr); i++ {
		tempSum += arr[i]

		if tempSum < 0 {
			tempSum = 0
		}

		if tempSum > sum {
			sum = tempSum
		}
	}

	return sum
}

func main() {
	input := [8]int{-2, -5, 6, -2, -3, 1, 5, -6}
	fmt.Println("Sorted output -->", contigiousSubarraySum(input))
}
