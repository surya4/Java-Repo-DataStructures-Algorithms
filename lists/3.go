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

func reverseLL(node *Node) int {

	if node == nil {
		fmt.Println(" Empty List")
		return -1
	}

	var prev *Node = nil
	var temp *Node = nil
	var current = node

	for current != nil {
		temp = current.next
		current.next = prev
		prev = current
		current = temp
	}
	node = prev

	fmt.Print("Reversed LL ---> ")
	for node != nil {
		fmt.Printf("%d -->", node.value)
		node = node.next
	}
	fmt.Println()
	return 1
}

func main() {
	root = nil
	addNode(root, 11)
	addNode(root, 22)
	addNode(root, 33)
	addNode(root, 44)
	addNode(root, 55)
	fmt.Print("Original LL ---> ")
	printList(root)
	reverseLL(root)
}
