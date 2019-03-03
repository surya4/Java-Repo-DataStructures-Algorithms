package main

import "fmt"

func zigZagFunc(arr []int) []int {
	output := make([]int, len(arr))
	copy(output, arr)

	var conditionFlag = true
	for i := 1; i < len(output); i++ {
		if conditionFlag == true {
			if output[i] < output[i-1] {
				output[i] = output[i] + output[i-1]
				output[i-1] = output[i] - output[i-1]
				output[i] = output[i] - output[i-1]
			}
		} else {
			if output[i] > output[i-1] {
				output[i] = output[i] + output[i-1]
				output[i-1] = output[i] - output[i-1]
				output[i] = output[i] - output[i-1]
			}
		}
		fmt.Println(" arr  conditionFlag-->", i, conditionFlag, output)
		conditionFlag = !conditionFlag
	}
	return output
}

func main() {
	input := []int{4, 3, 7, 8, 6, 2, 1}
	input1 := []int{1, 1, 1, 2, 2, 2, 3, 3}
	fmt.Println(input, " output -->", zigZagFunc(input))
	fmt.Println(input1, " output -->", zigZagFunc(input1))
}
