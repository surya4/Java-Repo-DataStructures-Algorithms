package main

import (
	"fmt"
	"strings"
)

func isRotationSubstrings(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var str = s1 + s1
	return strings.Contains(str, s2)
}

func main() {
	input1 := "hello world"
	input2 := " worldhello"
	fmt.Println(input1, input2, " <-- output --> ", isRotationSubstrings(input1, input2))
}
