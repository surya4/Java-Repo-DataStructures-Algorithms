package main

import "fmt"

type Node10 struct {
	data int
	next *Node10
}

var root10 = new(Node10)

func addNode10(node *Node10, data []int) {
	if len(data) == 0 {
		fmt.Println(" Empty List")
		return
	}

	if root10 == nil {
		node = &Node10{data[0], nil}
		root10 = node
	}
	for i := 1; i < len(data); i++ {
		node.next = &Node10{data[i], nil}
		node = node.next
	}
	fmt.Print("Before ---> ")
	printList10(root10)
}

func printList10(node *Node10) {
	if node == nil {
		fmt.Println(" Empty List")
		return
	}
	for node != nil {
		fmt.Printf(" %d -->", node.data)
		node = node.next
	}
	fmt.Println()
}

// brute-force (to-do)
// func removeDuplicates1(node *Node10) {
// 	if node == nil {
// 		fmt.Println(" Empty List")
// 		return
// 	}

// 	prev := &Node10{0, nil}
// 	curr := node
// 	for curr.next != nil {
// 		temp := node
// 		for temp != curr {
// 			if curr.data == temp.data {
// 				prev.next = temp
// 				curr = temp
// 				// fmt.Println("node.data-func ----> ", prev.next, temp, curr)
// 				break
// 			}
// 			temp = temp.next
// 			// fmt.Println("node.data-func ----> ", temp, curr)
// 		}
// 		if temp == curr {
// 			prev = curr
// 			curr = curr.next
// 		}
// 		fmt.Println()
// 	}

// 	fmt.Print("Brute force ----> ")
// 	printList10(root10)
// 	fmt.Println()
// }

// using hash
func removeDuplicates2(node *Node10) {
	if node == nil {
		fmt.Println(" Empty List")
		return
	}
	hm := make(map[int]int)
	temp := &Node10{0, nil}
	for node != nil {
		if hm[node.data] > 0 {
			temp.next = node.next
		} else {
			hm[node.data] = hm[node.data] + 1
			temp = node
		}
		node = node.next
	}
	fmt.Print("Hash-func ----> ")
	printList10(root10)
	fmt.Println()
}

func main() {
	root10 = nil
	input1 := []int{}
	input2 := []int{1, 2, 3, 4, 4, 5, 6}
	input3 := []int{1, 2, 3, 4, 5, 6, 6}
	input4 := []int{1, 1, 3, 4, 5, 6, 7}
	addNode10(root10, input1)
	removeDuplicates1(root10)
	// removeDuplicates2(root10)

	addNode10(root10, input2)
	removeDuplicates1(root10)
	// removeDuplicates2(root10)

	addNode10(root10, input3)
	removeDuplicates1(root10)
	// removeDuplicates2(root10)

	addNode10(root10, input4)
	removeDuplicates1(root10)
	// removeDuplicates2(root10)
}
