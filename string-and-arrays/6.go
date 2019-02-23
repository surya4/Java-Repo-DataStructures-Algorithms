package main

import "fmt"

// by counting values
func sortFunc1(arr [8]int) []int {
	count0 := 0
	count1 := 0
	count2 := 0
	var out []int

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count0++
		}
		if arr[i] == 1 {
			count1++
		}
		if arr[i] == 2 {
			count2++
		}
	}

	for i := 0; i < count0; i++ {
		out = append(out, 0)
	}
	for i := 0; i < count1; i++ {
		out = append(out, 1)
	}
	for i := 0; i < count2; i++ {
		out = append(out, 2)
	}
	return out
}

func sortFunc2(arr [12]int) [12]int {
	var low = 0
	var mid = 0
	var high = len(arr) - 1

	for mid <= high {
		switch arr[mid] {
		case 0:
			{
				arr[low] = arr[low] + arr[mid]
				arr[mid] = arr[low] - arr[mid]
				arr[low] = arr[low] - arr[mid]
				mid++
				low++
				break
			}

		case 1:
			{
				mid++
				break
			}

		case 2:
			{
				arr[high] = arr[high] + arr[mid]
				arr[mid] = arr[high] - arr[mid]
				arr[high] = arr[high] - arr[mid]
				high--
				break
			}
		}
	}
	return arr
}

func main() {
	input1 := [8]int{1, 0, 1, 2, 0, 2, 1, 0}
	input2 := [12]int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	fmt.Println(input1, " output 1 -->", sortFunc1(input1))
	fmt.Println(input2, " output 2 -->", sortFunc2(input2))
}
