package main

import (
	"fmt"
	"math"
)

func subArrayFunc(arr []int) int {
	hm := make(map[int]int)

	// for i := 1; i < len(arr); i++ {
	// 	arr[i] += arr[i-1]
	// }

	fmt.Println("arr", arr)

	var sum = 0
	var returnData = 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]

		// fmt.Println(hm[sum])

		if sum == 0 {
			returnData = i + 1
		}

		if hm[sum] > -1 {
			lastIndexSeen := hm[sum]
			returnData = int(math.Max(float64(returnData), float64(i-lastIndexSeen)))
		} else {
			hm[sum] = i
		}

		// 		for(int p = 0; p < arr.length; p++){
		// 			sumTillNow += arr[p];
		// 			if(sumTillNow == 0) largestSubArrayCount = p+1;
		// 			if(sumMap.containsKey(sumTillNow)){
		// 					int lastIndexSeen = sumMap.get(sumTillNow);
		// 					largestSubArrayCount = Math.max(largestSubArrayCount, p-lastIndexSeen);
		// 			}else{
		// 					sumMap.put(sumTillNow, p);
		// 			}
		// 	}
		// 	return largestSubArrayCount;
		// }

		fmt.Println(returnData)

	}
	// fmt.Println(hm)
	// fmt.Println(arr)
	return returnData
}

func main() {
	// input := [15,  -2,  2,  -8,  1,  7,  10, 23]
	// input := []int{1, 2, 3, 3, -9, 6, 7, -8, 1, 9}
	input := []int{15, -2, 2, -8, 1, 7, 10, 23}
	fmt.Println(input, " output -->", subArrayFunc(input))

}
