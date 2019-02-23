package main

import "fmt"

func basicFunc(arr [10]int, sum int) [10]int {
	// for i:=1; i < len(arr); i++ {
	// 	temp :=  i+1
	// 	if( temp != arr[i]) {
	// 		return temp
	// 	}
	// }
	return arr
}

func main() {
	input := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum:= 15
	fmt.Println("output -->", basicFunc(input, sum))
}
