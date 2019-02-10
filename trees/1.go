package main

import "fmt"

// Node comment
type Tree struct {
	left  *Tree
	val  int
	right *Tree
}

func insert(node *Tree, data int) *Tree {
	if node == nil {
		return &Tree{nil, data, nil}
	} 

	if data == node.val {
		return node
	} 

	if data < node.val {
		node.left = insert(node.left, data)
	} else {
		node.right = insert(node.right, data)
	}
	return node
}

func inOrderTraverse(node *Tree) {
	if node == nil { 
		return
	}
	inOrderTraverse(node.left)
	fmt.Print(node.val, " ")
	inOrderTraverse(node.right)
}

func preOrderTraverse(node *Tree) {
	if node == nil {
		return
	}
		fmt.Print(node.val, " ")
		preOrderTraverse(node.left)
		preOrderTraverse(node.right)
}

func postOrderTraverse(node *Tree) {
	if node == nil { 
		return
	}
	postOrderTraverse(node.left)
	postOrderTraverse(node.right)
	fmt.Print(node.val, " ")
}

func main() {
	tree := &Tree{}
	tree = insert(tree, 1)
	tree = insert(tree, 2)
	tree = insert(tree, 3)
	tree = insert(tree, 4)
	tree = insert(tree, 5)
	tree = insert(tree, 5)
	tree = insert(tree, 7)
	tree = insert(tree, 8)
	tree = insert(tree, 9)

fmt.Print("Inorder traversal: "); 
inOrderTraverse(tree);
fmt.Println();

fmt.Print("PreOrder traversal: "); 
preOrderTraverse(tree); 
fmt.Println();

fmt.Print("PostOrder traversal: "); 
postOrderTraverse(tree); 
fmt.Println();
}