package main

import (
	"fmt"
)

// recursion
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// iteration
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
	fmt.Println(" output 1 -->", fibonacci(max, fibonacciList))
	fmt.Println(" output 2 -->", fib(max))
}
