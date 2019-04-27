package main

import (
	"fmt"
)

// using array brute-force O(N^2)
func checkUniqueFunc1(str string) bool {
	// convert string to array of chars
	arr := []rune(str)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				fmt.Println("matched", string(arr[i]), i, j)
				return true
			}
		}
	}
	return false
}

// using single array O(n)
func checkUniqueFunc2(str string) bool {
	var arr [255]int
	for i, r := range str {
		if arr[r] > 0 {
			fmt.Println("matched", string(arr[i]), i, r, arr[r])
			return true
		}
		arr[r] = int(r)
	}
	return false
}

// using hash count array O(n)
func checkUniqueFunc3(str string) bool {
	hm := make(map[int]int)
	// fmt.Println("matched", hm[104])
	for _, r := range str {
		if hm[int(r)] > 0 {
			fmt.Println("matched", string(hm[int(r)]), r)
			return true
		}
		hm[int(r)] = hm[int(r)] + 1
	}
	fmt.Println("hm", hm)
	return false
}

// using string - second quest O(N^2)
func checkUniqueFunc4(str string) bool {
	for i, r1 := range str {
		var sub = str[i+1 : len(str)]
		for j, r2 := range sub {
			if r1 == r2 {
				fmt.Println("matched", string(r1), i, j)
				return true
			}
		}
	}
	return false
}

func main() {
	input := "hello world"
	fmt.Println(input, " output -->", checkUniqueFunc1(input))
	fmt.Println(input, " output -->", checkUniqueFunc2(input))
	fmt.Println(input, " output -->", checkUniqueFunc3(input))
	fmt.Println(input, " output -->", checkUniqueFunc4(input))
}
