package main

import "fmt"

// Node8 comment
type Node8 struct {
	data  int
	left  *Node8
	right *Node8
}

type bSTree8 struct {
	root *Node8
}

func (tree *bSTree8) insertNode8(data int, node *Node8) {
	temp := &Node8{data, nil, nil}

	if tree.root == nil {
		tree.root = temp
		return
	}

	curr := node
	par := curr
	if data < node.data {
		curr = curr.left
		if curr == nil {
			par.left = temp
			return
		}
	} else {
		curr = curr.right
		if curr == nil {
			par.right = temp
			return
		}
	}

	tree.insertNode8(data, curr)
}

func preOrderPrintNode8(node *Node8) {
	if node != nil {
		fmt.Print(node.data, " -> ")
		preOrderPrintNode8(node.left)
		preOrderPrintNode8(node.right)
	}
	return
}

func postOrderPrintNode8(node *Node8) {
	if node != nil {
		postOrderPrintNode8(node.left)
		postOrderPrintNode8(node.right)
		fmt.Print(node.data, " -> ")
	}
	return
}

func inOrderPrintNode8(node *Node8) {
	if node != nil {
		inOrderPrintNode8(node.left)
		fmt.Print(node.data, " -> ")
		inOrderPrintNode8(node.right)
	}
	return
}

func preIterativeTraversal8(node *Node8) {
	for node != nil {
		if node.left == nil {
			fmt.Print(node.data, " -> ")
			node = node.right
		} else {
			curr := node.left
			for curr.right != nil && curr.right != node {
				curr = curr.right
			}

			if curr.right == nil {
				fmt.Print(node.data, " -> ")
				curr.right = node
				node = node.left
			} else {
				curr.right = nil
				node = node.right
			}
		}
	}
}

// func postIterativeTraversal8(node *Node8) {
// 	for node != nil {
// 		if node.left == nil {
// 			fmt.Print(node.data, " -> ")
// 			node = node.right
// 		} else {
// 			curr := node.left
// 			for curr.right != nil && curr.right != node {
// 				curr = curr.right
// 			}

// 			fmt.Println(node.data, curr, " -> ")

// 			if curr.right == nil {
// 				curr.right = node
// 				node = node.left
// 				// fmt.Print(node.data, " -> ")
// 			} else {
// 				curr.right = nil
// 				// fmt.Print(node.data, " -> ")
// 				node = node.right
// 			}
// 			// fmt.Print(node.data, " -> ")
// 		}
// 	}
// }

func inIterativeTraversal8(node *Node8) {
	for node != nil {
		if node.left == nil {
			fmt.Print(node.data, " -> ")
			node = node.right
		} else {
			curr := node.left
			for curr.right != nil && curr.right != node {
				curr = curr.right
			}

			if curr.right == nil {
				curr.right = node
				node = node.left
			} else {
				curr.right = nil
				fmt.Print(node.data, " -> ")
				node = node.right
			}
		}
	}
}

func main() {
	tree := &bSTree8{}
	input1 := []int{9, 4, 17, 3, 6, 5, 7, 22, 20}
	for i := 0; i < len(input1); i++ {
		tree.insertNode8(input1[i], tree.root)
	}
	fmt.Print("preOrder Recursive ")
	preOrderPrintNode8(tree.root)
	fmt.Println()
	fmt.Print("postOrder Recursive ")
	postOrderPrintNode8(tree.root)
	fmt.Println()
	fmt.Print("inOrder Recursive ")
	inOrderPrintNode8(tree.root)
	fmt.Println()

	fmt.Println()
	fmt.Print("preOrder Iterative ")
	preIterativeTraversal8(tree.root)
	fmt.Println()
	// fmt.Print("postOrder Iterative ")
	// postIterativeTraversal8(tree.root)
	// fmt.Println()
	fmt.Print("inOrder Iterative ")
	inIterativeTraversal8(tree.root)
	fmt.Println()
}
