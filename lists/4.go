package main

import "fmt"

// Node comment
type Node struct {
	value int
	next  *Node
}

var root = new(Node)

func addNode(head *Node, value int) int {

	if root == nil {
		head = &Node{value, nil}
		root = head
		return 1
	}

	if head.next == nil {
		head.next = &Node{value, nil}
		return 1
	}

	for head.next != nil {
		head = head.next
	}

	head.next = &Node{value, nil}
	return 1
}

func printList(head *Node) {

	if head == nil {
		fmt.Println(" Empty List")
		return
	}

	for head != nil {
		fmt.Printf("%d -->", head.value)
		head = head.next
	}
	fmt.Println()
}

func rotateLL(head *Node, k int) int {

	// check if list is empty or k < 0
	if head == nil || k < 0 {
		fmt.Println("Empty List or wrong K")
		return -1
	}

	// keep parent node head in a a different node to create subsets
	var current = head
	// iterate the temp node till reaches given
	count := 1
	for count < k && current != nil {
		count++
		current = current.next
	}

	// store next node of temp in a diffetent node
	var newHead = current.next
	// keep new team in again a differnt node so iterate it and add null after end of list
	var temp = newHead
	// add null and close cylcic behaviour of list
	current.next = nil

	// iterate till end to add the first half od list
	for temp.next != nil {
		temp = temp.next
	}

	// add the first half of list
	temp.next = head

	// traverse and print it
	for newHead != nil {
		fmt.Print(newHead.value, "-->")
		newHead = newHead.next
	}

	fmt.Println()
	return 1
}


func main() 
{
	root = nil
	addNode(root, 11)
	addNode(root, 22)
	addNode(root, 33)
	addNode(root, 44)
	addNode(root, 55)
	addNode(root, 66)
	addNode(root, 77)
	addNode(root, 88)
	fmt.Print("Original LL ---> ")
	printList(root)
	rotateLL(root, 4)
}
