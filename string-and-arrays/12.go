package main

import (
	"fmt"
	"strings"
)

func charReplace(s string) string {
	for i := range s {
		s = strings.Replace(s, " ", "%20", i)
	}

	return s
}

func main() {
	input := "hello world "
	fmt.Println(input, " <-- output --> ", charReplace(input))
}
