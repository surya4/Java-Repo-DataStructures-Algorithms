package main

import "fmt"

type Node11 struct {
	data int
	next *Node11
}

var root11 = new(Node11)

func addNode11(node *Node11, data []int) {
	if len(data) == 0 {
		fmt.Println(" Empty List")
		return
	}
	if root11 == nil {
		node = &Node11{data[0], nil}
		root11 = node
	}
	for i := 1; i < len(data); i++ {
		node.next = &Node11{data[i], nil}
		node = node.next
	}
	fmt.Print("List ---> ")
	printList11(root11)
}

func printList11(node *Node11) {
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

// using count and length of list
func getSubList1(node *Node11, pos int) {
	if node == nil || pos < 0 {
		fmt.Println(" Empty List")
		return
	}

	list_len := 0

	temp1 := node
	temp2 := node

	for temp1.next != nil {
		list_len += 1
		temp1 = temp1.next
	}

	for i := 0; i < (list_len - pos); i++ {
		temp2 = temp2.next
	}

	fmt.Println(pos, "th from last", "is: ", temp2.data)
}

func getSubList2(node *Node11, pos int) {
	if node == nil || pos < 0 {
		fmt.Println(" Empty List")
		return
	}

	temp1 := node
	temp2 := node

	for i := 0; i < pos; i++ {
		temp1 = temp1.next
	}

	for temp1.next != nil {
		temp1 = temp1.next
		temp2 = temp2.next
	}

	fmt.Println(pos, "th from last", "is: ", temp2.data)
}

func main() {
	root11 = nil
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	addNode11(root11, input)
	getSubList1(root11, 4)
	getSubList2(root11, 4)
	getSubList1(root11, 0)
	getSubList2(root11, 0)
}
