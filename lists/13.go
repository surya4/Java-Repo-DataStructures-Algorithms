package main

import "fmt"

type Node13 struct {
	data int
	next *Node13
}

func addNode13(data []int) *Node13 {
	if len(data) == 0 {
		fmt.Println("Empty List")
	}

	node := &Node13{data[0], nil}
	head := node

	for i := 1; i < len(data); i++ {
		node.next = &Node13{data[i], nil}
		node = node.next
	}
	fmt.Print("List ---> ")
	printList13(head)
	return head
}

func printList13(node *Node13) {
	if node == nil {
		fmt.Println("Empty List")
	}

	for node != nil {
		fmt.Printf("%d -> ", node.data)
		node = node.next
	}

	fmt.Println()
}

func addLists(node1, node2 *Node13) {
	if node1 == nil || node2 == nil {
		fmt.Println("Empty List")
	}

	node := &Node13{0, nil}
	temp := node
	carry := 0
	val := 0

	for node1 != nil || node2 != nil || carry > 0 {
		val = carry
		if node1 != nil {
			val += node1.data
			node1 = node1.next
		}
		if node2 != nil {
			val += node2.data
			node2 = node2.next
		}
		temp.next = &Node13{val % 10, nil}
		carry = val / 10
		temp = temp.next
	}
	node = node.next
	fmt.Print("Final List ---> ")
	printList13(node)
}

func main() {
	input1 := []int{7, 2, 3, 4, 5, 6, 7, 9}
	input2 := []int{6, 1, 2, 3, 7, 5, 6}
	addLists(addNode13(input1), addNode13(input2))
}
