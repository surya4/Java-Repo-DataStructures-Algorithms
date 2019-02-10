package main

import "fmt"

// Node comment
type Node struct {
	val int
	next  *Node
}

var root = new(Node)

func addNode(node *Node, data int) int {

	if root == nil {
		node = &Node{data, nil}
		root = node
		return 0
	}

	for node.next != nil {
		node = node.next
	}

	node.next = &Node{data, nil}
	return 1
}

func recursivelookUpNode(t *Node, data int) bool {
	if root == nil {
		t := &Node{data, nil}
		root = t
		return false
	}

	if data == t.val {
		return true
	}

	if t.next == nil {
		return false
	}

	return recursivelookUpNode(t.next, data)
}

func iterativelookUpNode(t *Node, data int) bool {
	if root == nil {
		t := &Node{data, nil}
		root = t
		return false
	}

	for t.next != nil {
		if data == t.val {
			return true
		}
		t = t.next
	}

	return false
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

func size(node *Node) int {
	if node == nil {
		fmt.Println(" Empty List")
		return 0
	}

	size := 0

	for node != nil {
		size++
		node = node.next
	}
	return size
}

func main() {
	root = nil
	addNode(root, 1)
	addNode(root, 2)
	addNode(root, 3)
	addNode(root, 4)
	printList(root)
	fmt.Println("Is 78 availilable in single LL -> ", recursivelookUpNode(root, 78))
	fmt.Println("Is 3 availilable in single LL -> ", recursivelookUpNode(root, 3))
	fmt.Println("Is 78 availilable in single LL -> ", iterativelookUpNode(root, 78))
	fmt.Println("Is 3 availilable in single LL -> ", iterativelookUpNode(root, 3))
	fmt.Println("List size -> ", size(root))
}
