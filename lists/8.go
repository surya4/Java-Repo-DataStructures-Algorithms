package main

import "fmt"

// Node comment
type Node struct {
	val int
	prev  *Node
	next  *Node
}

var root = new(Node)

func addNode(node *Node, data int) int {

	if root == nil {
		node = &Node{data, nil, nil}
		root = node
		return 0
	}

	if node.next == nil {
		// temp := node
		// node.next = &Node{data, temp, nil}
		// temp := node
		node.next = &Node{data, node, nil}
		return -2
	}

	return addNode(node.next, data)
}

// func deleteFromHead() {
// fmt.Printf(" %d -->", node, Node)
// node = Node.next
// }

// func deleteFromTail(data int) int {

// }
// func deleteByIndex(data int) int {

// }
// func deleteByval(data int) int {

// }

func printList(node *Node) {
	if node == nil {
		fmt.Println(" Empty List")
		return
	}

	for node != nil {
		fmt.Printf(" %d -->", node.val)
		node = node.next
	}
	fmt.Println()
}

func traverseBack(node *Node) {
	if node == nil {
		fmt.Println(" Empty List")
		return
	}

	temp := node

	for node != nil {
		temp = node
		node = node.next
	}

	for temp != nil {
		fmt.Printf(" %d -->", temp.val)
		temp = temp.prev
	}

	fmt.Println()
}

func main() {
	root = nil
	addNode(root, 1)
	addNode(root, 2)
	addNode(root, 3)
	addNode(root, 4)
	printList(root)
	traverseBack(root)
}