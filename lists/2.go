package main

import "fmt"

// Node comment
type Node struct {
	value int
	next  *Node
}

var root = new(Node)

func addNode(node *Node, value int) int {

	if root == nil {
		node = &Node{value, nil}
		root = node
		return 1
	}

	if node.next == nil {
		node.next = &Node{value, nil}
		return 1
	}

	for node.next != nil {
		node = node.next
	}

	node.next = &Node{value, nil}
	return 1
}

func printList(node *Node) {

	if node == nil {
		fmt.Println(" Empty List")
		return
	}

	for node != nil {
		fmt.Printf("%d -->", node.value)
		node = node.next
	}
	fmt.Println()
}

func middleElemFunc(node *Node) int {

	if node == nil {
		fmt.Println(" Empty List")
		return -1
	}

	var slow = node
	var fast = node

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow.value
}

func main() {
	root = nil
	addNode(root, 11)
	addNode(root, 22)
	addNode(root, 33)
	addNode(root, 44)
	addNode(root, 55)
	printList(root)
	fmt.Println("middleElemFunc value --->", middleElemFunc(root))
}
