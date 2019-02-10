package main

import (
	"fmt"
	"sort"
)

func sumOfArray(arr []int, sum int) []int {
	sort.Ints(arr)
	return arr
	// return make([]int, 2)
}

func main() {
	arr := []int{1, 4, 45, 6, 10, 8}
	sum := 16
	fmt.Println(" output -->", sumOfArray(arr, sum))
}
