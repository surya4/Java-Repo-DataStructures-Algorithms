# Trees

* Solution 1. Implement binary tree
* Solution 2. Implement binary search tree
* Solution 3. Implement binary min-heap
* Solution 4. Implement binary max-heap
* Solution 5. Implement AVL tree (Red-Black Tree)
* Solution 6. height of tree
* Solution 7. Implement Red-Black tree
* Solution 8. Iterative traversal of Tree WithoutStack Without Recursion
* Solution 9. Find Closest Value In BST
* Solution 10. 
* Solution 11. 
* Solution 12. 
* Solution 13. 


<!-- func closestNodeInBSTree8(n int, node *Node8) int {
	nearestDiff := 10000000000

	curr := node

	for curr != nil {
		if curr.left == nil {
			fmt.Print(curr.data, " ")
			curr = curr.right
		} else {
			pre := curr.left
			for pre.right != nil && pre.right != curr {
				pre = pre.right
			}

			if pre.right == nil {
				pre.right = curr
				curr = curr.left
			} else {
				pre.right = nil
				fmt.Print(curr.data, " ")
				curr = curr.right
			}
		}
	}

	return n - nearestDiff
} -->