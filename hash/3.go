package main

import (
	"fmt"
	"math"
)

func subArrayFunc(arr []int) int {
	hm := make(map[int]int)
	max_el := 0

	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		if arr[i] == 0 {
			max_el = math.Max(max_el, 1)
		}

		if sum == 0 {
			max_el = math.Max(max_el, i+1)
		}

		hm[sum] := i


		// hm[sum] = hm[sum] + 1
		// old := hm[sum]
		// fmt.Println(sum, old)

		// if old == 0 {
		// 	hm[sum] = i
		// 	sum += arr[i]
		// } else {
		// 	fmt.Println("here i am")
		// }

		// if val, ok := hm[sum]; ok {
		// 	//do something here
		// 	fmt.Println("here i am", val, ok)
		// }

		// hm[sum] = i

		// fmt.Println(hm[sum])
		// if hm[sum] != 0 {
		// 	fmt.Println(hm[sum])
		// 	// hm[sum] = hm[sum]++
		// } else {
		// 	hm[sum] = 1
		// }

		// for k, _ := range hm {
		// 	// fmt.Println(k, hm[k])
		// 	hm[k+arr[i]] = hm[k+arr[i]] + 1
		// 	fmt.Println(k)
		// }

		// if arr[i] == 0 || sum == 0 || hm[sum] != 0 {
		// 	state = true
		// 	break
		// 	// return state
		// 	// fmt.Println("here")
		// }
		// hm.put(sum, i)
		fmt.Println(hm)
	}
	// fmt.Println(hm)
	// fmt.Println(arr)
	return max_el
}

func main() {
	// input := [15,  -2,  2,  -8,  1,  7,  10, 23]
	input := []int{15, -2, 2, -8, 1, 7, 10, 23}
	fmt.Println(input, " output -->", subArrayFunc(input))

}
