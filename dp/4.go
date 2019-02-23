package main

import (
	"fmt"
	"sort"
)

func maxSumFunc(arr []int) int {
	var arryLen = len(arr)
	sum := make([]int, arryLen)
	pos := make([]int, arryLen)

	for i := 0; i < len(arr); i++ {
		sum[i] = arr[i]
		pos[i] = i
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && sum[i] < sum[j]+arr[i] {
				sum[i] = sum[j] + arr[i]
				pos[i] = j
			}
		}
	}

	// elemenst part of max sum
	maxSumElementsFunc(sum, pos, arr)

	var tempSum = arr[0]
	for i := 1; i < len(sum); i++ {
		if sum[i] > tempSum {
			tempSum = sum[i]
		}
	}
	return tempSum
}

func maxSumElementsFunc(sum, pos, arr []int) {
	var maxPos = 0
	var tempMaxSum = sum[0]
	var arryLen = len(sum)

	posArray := make([]int, arryLen)

	for i := 0; i < len(sum); i++ {
		posArray[i] = -1
	}

	for i := 1; i < len(sum); i++ {
		if sum[i] > tempMaxSum {
			tempMaxSum = sum[i]
			maxPos = i
		}
	}

	posArray = append(posArray, maxPos)

	for i := maxPos - 1; i >= 0; i-- {
		if pos[i] <= pos[maxPos] && arr[i] <= arr[maxPos] {
			posArray = append(posArray, i)
		}
	}

	sort.Ints(posArray)

	for i := 0; i < len(posArray); i++ {
		if posArray[i] > -1 {
			fmt.Print(arr[posArray[i]], " ")
		}
	}
	fmt.Println()
}

func main() {
	input := []int{1, 101, 2, 3, 100, 4, 5}
	input2 := []int{4, 6, 1, 3, 8, 4, 6}
	fmt.Println(input, " output -->", maxSumFunc(input))
	fmt.Println(input2, " output -->", maxSumFunc(input2))
}
