package main

import "fmt"

// Node comment
type Node struct {
	data  int
	left  *Node
	right *Node
}

type bSTree struct {
	root *Node
}

func (tree *bSTree) insert(data int) {
	fmt.Println("data: ", data); 

	node := &Node{data, nil, nil}
	fmt.Println("node: ", node); 
	// if tree == nil {
	// 	tree.root = node
	// } else {
	// 	current := tree.root;
	// 	parent := current;
		// for true {
		// 	parent = current;
		// 	fmt.Print("parent: ", parent); 

		// 	if (data < current.data) {
		// 		current = current.left;
		// 		if (current == nil) {
		// 			parent.left = node
		// 			break;
		// 		}
		// 	} else {
		// 		current = current.right;
		// 		if (current == nil) {
		// 			parent.left = node
		// 			break;
		// 		}
		// 	}
		// }
	// }
}

func inOrderTraverse(root *Node) {
	if root != nil {
		inOrderTraverse(root.left)
		fmt.Print(root.data, " ")
		inOrderTraverse(root.right)
	}
}

func preOrderTraverse(root *Node) {
	if root != nil {
		fmt.Print(root.data, " ")
		preOrderTraverse(root.left)
		preOrderTraverse(root.right)
	}
}

func postOrderTraverse(root *Node) {
	if root != nil {
		postOrderTraverse(root.left)
		postOrderTraverse(root.right)
		fmt.Print(root.data, " ")
	}
}

func main() {
	tree := &bSTree{}
	tree.insert(1)
	tree.insert(2)
	tree.insert(3)
	tree.insert(4)
	tree.insert(5)
	tree.insert(5)
	tree.insert(7)
	tree.insert(8)
	tree.insert(9)

// root := &Node{data: 1, left: nil, right: nil}

// fmt.Print("Inorder traversal: "); 
// inOrderTraverse(root);
// fmt.Println();

// fmt.Print("PreOrder traversal: "); 
// preOrderTraverse(root); 
// fmt.Println();

// fmt.Print("PostOrder traversal: "); 
// postOrderTraverse(root); 
// fmt.Println();
}