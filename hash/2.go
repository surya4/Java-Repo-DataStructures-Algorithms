package main

import "fmt"

func sumArrFunc(s string) string {
	return s
}

func main() {
	input := "hello world"
	fmt.Println(input, " output -->", sumArrFunc(input))
}
