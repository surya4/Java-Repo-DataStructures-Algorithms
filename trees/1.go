package main

import "fmt"

// Node comment
type Tree1 struct {
	left  *Tree1
	val   int
	right *Tree1
}

func insert1(node *Tree1, data int) *Tree1 {
	if node == nil {
		return &Tree1{nil, data, nil}
	}

	if data == node.val {
		return node
	}

	if data < node.val {
		node.left = insert1(node.left, data)
	} else {
		node.right = insert1(node.right, data)
	}
	return node
}

func inOrderTraverse1(node *Tree1) {
	if node == nil {
		return
	}
	inOrderTraverse1(node.left)
	fmt.Print(node.val, " ")
	inOrderTraverse1(node.right)
}

func preOrderTraverse1(node *Tree1) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	preOrderTraverse1(node.left)
	preOrderTraverse1(node.right)
}

func postOrderTraverse1(node *Tree1) {
	if node == nil {
		return
	}
	postOrderTraverse1(node.left)
	postOrderTraverse1(node.right)
	fmt.Print(node.val, " ")
}

func main() {
	tree1 := &Tree1{}
	tree1 = insert1(tree1, 1)
	tree1 = insert1(tree1, 2)
	tree1 = insert1(tree1, 3)
	tree1 = insert1(tree1, 4)
	tree1 = insert1(tree1, 5)
	tree1 = insert1(tree1, 5)
	tree1 = insert1(tree1, 7)
	tree1 = insert1(tree1, 8)
	tree1 = insert1(tree1, 9)

	fmt.Print("Inorder traversal: ")
	inOrderTraverse1(tree1)
	fmt.Println()

	fmt.Print("PreOrder traversal: ")
	preOrderTraverse1(tree1)
	fmt.Println()

	fmt.Print("PostOrder traversal: ")
	postOrderTraverse1(tree1)
	fmt.Println()
}
