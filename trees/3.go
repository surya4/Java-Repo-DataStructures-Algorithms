package main

import (
	"fmt"
	"math"
)

var arr []int

func insertData(array []int) {
	for i := 0; i < len(array); i++ {
		// fmt.Println("parent", array)
		heapifyUp(array[i])
	}
}

func heapifyUp(pos int) {
	if pos > -1 {

		parent := int(math.Floor(float64((pos+1)/2)) - 1)

		fmt.Println("parent", pos, parent)
		if parent > -1 && arr[pos] < arr[parent] {
			arr[parent] = arr[parent] + arr[pos]
			arr[pos] = arr[parent] - arr[pos]
			arr[parent] = arr[parent] - arr[pos]
		}
		fmt.Println("arr", arr)
		heapifyUp(parent)
	}
	return
}

// func heapifydown(pos int) {

// 	fmt.Println("pos", pos)

// }

func main() {
	input := []int{11, 12, 13, 41, 5, 16, 17, 81, 19, 110}
	// fmt.Println("input", insertData)
	insertData(input)
	fmt.Println("out ->", arr)
	// heapifydown(input)
}
