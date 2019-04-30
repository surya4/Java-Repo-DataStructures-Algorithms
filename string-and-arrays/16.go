package main

import "fmt"

// brute-force iteration
func twoSums1(arr []int, sum int) [2]int {
	out := [2]int{}
	if len(arr) == 0 {
		return out
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				// s = append(s, 0)
				out[0] = arr[i]
				out[1] = arr[j]
				return out
			}
		}
	}

	return out
}

// using hash
func twoSums2(arr []int, sum int) [2]int {
	out := [2]int{}
	hm := make(map[int]int)

	if len(arr) == 0 {
		return out
	}

	for i := 0; i < len(arr); i++ {
		if hm[sum-arr[i]] > 0 {
			out[1] = arr[i]
			out[0] = sum - arr[i]
			return out
		}
		hm[arr[i]] = 1
	}

	return out
}

func main() {
	input1 := []int{2, 7, 11, 15}
	input2 := []int{}
	sum1 := 9
	fmt.Println(input1, sum1, " <-- input,  output --> ", twoSums1(input1, sum1))
	fmt.Println(input2, sum1, " <-- input,  output --> ", twoSums1(input2, sum1))
	fmt.Println(input1, sum1, " <-- input,  output --> ", twoSums2(input1, sum1))
	fmt.Println(input2, sum1, " <-- input,  output --> ", twoSums2(input2, sum1))
}
