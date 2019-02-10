package main

import "fmt"

func basicFunc(s string) string {
	return s
}

func main() {
	input := "hello world"
	fmt.Println(input, " output -->", basicFunc(input))
}
