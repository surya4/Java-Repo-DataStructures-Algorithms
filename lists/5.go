package main

import "fmt"

// Node comment
type Node struct {
	value int
	next  *Node
}

var root = new(Node)

func detectLoop(node *Node) bool {

	if node == nil {
		fmt.Println(" Empty List")
		return false
	}

	var hare = node
	var tort = node
	var count = 0
	for hare != nil || hare.next != nil || tort != nil {
		fmt.Println("%d -->", count)
		tort = tort.next
		hare = hare.next.next
		count++

		if count > 10 {
			break
		}

		if hare == tort {
			return true
		}
	}
	return false
}

func main() {
	root = &Node{value: 1, next: root}
	root.next = &Node{value: 2, next: root.next}
	root.next.next = &Node{value: 3, next: root.next.next}
	root.next.next.next = &Node{value: 4, next: root.next.next.next}
	root.next.next.next.next = &Node{value: 5, next: root.next.next}
	fmt.Println("detectLoop ---> ", detectLoop(root))
}
