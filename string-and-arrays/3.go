package main

import "fmt"

func basicFunc(arr [8]int) int {
	for i:=1; i < len(arr); i++ {
		temp :=  i+1
		if( temp != arr[i]) {
			return temp
		}
	}
	return -1
}

func main() {
	input := [8]int{1,2,3,4,5,6, 7, 8}
	fmt.Println(input, " output -->", basicFunc(input))
}
