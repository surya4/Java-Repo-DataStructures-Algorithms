package main

import "fmt"

// Node2 comment
type Node2 struct {
	data  int
	left  *Node2
	right *Node2
}

type bSTree2 struct {
	root *Node2
}

func (tree *bSTree2) insert(data int) {
	fmt.Println("data: ", data)

	node := &Node2{data, nil, nil}
	fmt.Println("node: ", node)
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

func (tree *bSTree2) insertNode2(data int, node *Node2) {
	temp := &Node2{data, nil, nil}

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

	tree.insertNode2(data, curr)
}

func preOrderPrintNode2(node *Node2) {
	if node != nil {
		fmt.Print(node.data, " -> ")
		preOrderPrintNode2(node.left)
		preOrderPrintNode2(node.right)
	}
	return
}

func postOrderPrintNode2(node *Node2) {
	if node != nil {
		postOrderPrintNode2(node.left)
		postOrderPrintNode2(node.right)
		fmt.Print(node.data, " -> ")
	}
	return
}

func inOrderPrintNode2(node *Node2) {
	if node != nil {
		inOrderPrintNode2(node.left)
		fmt.Print(node.data, " -> ")
		inOrderPrintNode2(node.right)
	}
	return
}

func main() {
	tree := &bSTree2{}
	input1 := []int{9, 4, 17, 3, 6, 5, 7, 22, 20}
	for i := 0; i < len(input1); i++ {
		tree.insertNode2(input1[i], tree.root)
	}
	fmt.Print("preOrder ")
	preOrderPrintNode2(tree.root)
	fmt.Println()
	fmt.Print("postOrder ")
	postOrderPrintNode2(tree.root)
	fmt.Println()
	fmt.Print("inOrder ")
	inOrderPrintNode2(tree.root)
	fmt.Println()
}
