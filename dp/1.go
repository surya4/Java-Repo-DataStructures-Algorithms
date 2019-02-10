package main

import (
	"fmt"
)

func fibonacci(max int, fibonacciList []int) []int {

	if max > 1 {
		for i := 2; i < max; i++ {
			fibonacciList = append(fibonacciList, fibonacciList[i-1]+fibonacciList[i-2])
		}
	}
	return fibonacciList
}

func main() {
	max := 16
	fibonacciList := make([]int, 0)
	fibonacciList = append(fibonacciList, 0, 1)
	fmt.Println(" output -->", fibonacci(max, fibonacciList))
}
