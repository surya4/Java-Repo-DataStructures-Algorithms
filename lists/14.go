package main

import "fmt"

type Node14 struct {
	data int
	next *Node14
}

var root14 = new(Node14)

func addNodes14(node *Node14, newNode, nextNode int) {
	fmt.Println("node -> ", node)
	curr := &Node14{newNode, nil}

	if node == nil {
		root14 = curr
		return
	}

	// node.next = &Node14{newNode, nil}

	temp := node

	// node.next = curr

	// head := node
	// fmt.Println("-> ", head, newNode, nextNode)

	if nextNode > 0 {
		for node != nil {
			node = node.next
			fmt.Print("inside -> ", node)
			if node.data == nextNode {
				break
			}
		}
	} else {
		node.next = &Node14{newNode, nil}
	}

	// node.next = &Node14{newNode, nil}

	// fmt.Println("-> ", head, newNode, nextNode)

	// temp := &Node14{newNode, nil}
	// temp.next = head

	// fmt.Println("temp-> ", node, temp, head, newNode, nextNode)

	for temp != nil {
		fmt.Println("nodeeee-> ", temp)
		temp = temp.next
	}
}

func printList14(node *Node14) {
	fmt.Println("printList14 node", node)
	if node == nil {
		fmt.Println("Empty List")
		return
	}
	for node != nil {
		fmt.Print("-> ", node.data)
		node = node.next
	}
	fmt.Println()
}

func checkIfLoopExists(node *Node14) bool {
	fmt.Println("node", node)
	if node == nil {
		return false
	}

	var count = 0

	var fast = node
	var slow = node

	for fast != nil || fast.next != nil || slow != nil {
		fmt.Println("fast, slow ---> ", fast, slow)
		slow = slow.next
		fast = fast.next.next
		count++

		if fast == slow {
			fmt.Println("fast, slow ---> ", fast, slow)
			return true
		}

		if count == 50 {
			break
		}
	}
	return false
}

func beginningOfCirculareLoop(node *Node14) {
	// head := node
	var fast = node
	var slow = node

	for fast != nil || fast.next != nil || slow != nil {
		fmt.Println("fast, slow ---> ", fast, slow)
		slow = slow.next
		fast = fast.next.next
		// count++

		if fast == slow {
			fmt.Println("fast, slow ---> ", fast, slow)
			// return true
			break
		}

		// if count == 50 {
		// 	break
		// }
	}

	// if checkIfLoopExists(head) == true {

	// } else {
	// 	fmt.Print("List ois not a loop ")
	// }
	// fmt.Print("Final ")
	// fmt.Println("detectLoop ---> ", checkIfLoopExists(head))
	// printList14(head)
	// return s
}

func main() {
	root14 = nil
	addNodes14(1, -1)
	fmt.Println("root14 -> ", root14)
	addNodes14(2, -1)
	fmt.Println("-> ", root14)
	// addNodes14(root14, 3, 2)
	// fmt.Println("-> ", newNode, nextNode)
	// addNodes14(root14, 4, 3)
	// addNodes14(root14, 5, 4)
	// addNodes14(root14, 6, 5)
	// addNodes14(root14, 7, 6)
	// printList14(root14)

	// input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 3}
	// root14 = &Node14{data: 1, next: root14.next}
	// root14.next = &Node14{data: 2, next: root14.next.next}
	// root14.next.next = &Node14{data: 3, next: root14.next.next}
	// root14.next.next.next = &Node14{data: 4, next: root14.next.next.next}
	// root14.next.next.next.next = &Node14{data: 5, next: root14.next.next.next.next}
	// root14.next.next.next.next.next = &Node14{data: 6, next: root14.next.next.next.next.next}
	// root14.next.next.next.next.next.next = &Node14{data: 7, next: root14.next.next.next.next.next.next}
	// root14.next.next.next.next.next.next.next = &Node14{data: 8, next: root14.next.next.next.next.next.next.next}
	// root14.next.next.next.next.next.next.next.next = &Node14{data: 9, next: root14.next.next.next.next.next.next.next.next}
	// addNodes14(input)
	// res, err := checkIfLoopExists(root14)
	// fmt.Println("checkIfLoopExists ", checkIfLoopExists(root14))

	// if err != nil {
	// 	fmt.Print("err ", err)
	// } else {
	// 	fmt.Print("res ", res)
	// }
	// beginningOfCirculareLoop(root14)
}
