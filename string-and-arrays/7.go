package main

import "fmt"

func equilibFunc(arr []int) int {
	var sum = 0
	var sum2 = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	for i := 0; i < len(arr); i++ {
		sum2 += arr[i]
		sum -= arr[i]
		fmt.Println(" sum2 -->", sum, sum2)
		if sum == sum2 {
			return arr[i]
		}
	}
	return -1
}

func main() {
	input := []int{1, 2, 5, 4, 3, 1}
	input1 := []int{1, 2, 3, 4, 2, 1}
	fmt.Println(input, " output -->", equilibFunc(input))
	fmt.Println(input1, " output -->", equilibFunc(input1))

}
