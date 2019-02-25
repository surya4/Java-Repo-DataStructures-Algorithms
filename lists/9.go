package main

import "fmt"

type Node struct {
	val  int
	prev *Node
	next *Node
}

var root = new(Node)

func addNode(node *Node, data int) int {
	if root == nil {
		node = &Node{data, nil, nil}
		root = node
		return 0
	}

	if node.next == nil {
		node.next = &Node{data, node, nil}
		return 1
	}
	return addNode(node.next, data)
}

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

func main() {
	root = nil
	addNode(root, 1)
	addNode(root, 2)
	addNode(root, 3)
	addNode(root, 4)
	printList(root)
}
