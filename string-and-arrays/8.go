package main

import "fmt"

func zigZagFunc(arr []int) []int {
	var conditionFlag = true
	for i := 1; i < len(arr); i++ {
		if conditionFlag == true {
			if arr[i] < arr[i-1] {
				arr[i] = arr[i] + arr[i-1]
				arr[i-1] = arr[i] - arr[i-1]
				arr[i] = arr[i] - arr[i-1]
			}
		} else {
			if arr[i] > arr[i-1] {
				arr[i] = arr[i] + arr[i-1]
				arr[i-1] = arr[i] - arr[i-1]
				arr[i] = arr[i] - arr[i-1]
			}
		}
		fmt.Println(" arr  conditionFlag-->", conditionFlag, arr)
		conditionFlag = !conditionFlag
	}
	return arr
}

func main() {
	input := []int{4, 3, 7, 8, 6, 2, 1}
	fmt.Println(input, " output -->", zigZagFunc(input))
}
