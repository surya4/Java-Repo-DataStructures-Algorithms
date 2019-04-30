package main

import "fmt"

type Node12 struct {
	data int
	next *Node12
}

var root12 = new(Node12)

func addNode12(node *Node12, data []int) {
	if len(data) == 0 {
		fmt.Println("EmptyList")
	}

	if root12 == nil {
		node = &Node12{data[0], nil}
		root12 = node
	}

	for i := 1; i < len(data); i++ {
		node.next = &Node12{data[i], nil}
		node = node.next
	}
	fmt.Print("List ---> ")
	printList12(root12)
}

func printList12(node *Node12) {
	if node == nil {
		fmt.Println("Empty List")
		return
	}
	for node != nil {
		fmt.Printf(" %d -->", node.data)
		node = node.next
	}
	fmt.Println()
}

func middleOfList(node *Node12) {
	if node == nil {
		fmt.Println("Empty List")
		return
	}

	temp1 := node
	temp2 := node
	prev := &Node12{0, nil}

	for temp1 != nil && temp1.next != nil {
		temp1 = temp1.next.next
		prev = temp2
		temp2 = temp2.next
	}

	prev.next = temp2.next
	printList12(node)
}

func main() {
	root12 = nil
	input := []int{1, 2, 3, 4, 5, 6, 7}
	addNode12(root12, input)
	middleOfList(root12)
}
