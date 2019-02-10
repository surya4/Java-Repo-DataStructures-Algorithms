package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("input --->", runes[i], runes[j])
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "hel_lo worls"
	fmt.Println(input, " reverse output -->", reverseString(input))
}
