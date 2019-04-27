package main

import "fmt"

func angramCheck(s1, s2 string) bool {
	hm := make(map[int]int)
	for _, r := range s1 {
		hm[int(r)] = hm[int(r)] + 1
	}

	for _, r := range s2 {
		if hm[int(r)] > 0 {
			hm[int(r)] = hm[int(r)] - 1
		} else {
			hm[int(r)] = hm[int(r)] + 1
		}
	}

	for _, v := range hm {
		if v > 0 {
			return false
		}
	}

	return true
}

func main() {
	input1 := "hello world"
	input2 := "world helloo"
	input3 := "world hello"
	fmt.Println(input1, input2, " <-- output --> ", angramCheck(input1, input2))
	fmt.Println(input1, input3, " <-- output --> ", angramCheck(input1, input3))
}
