class BinaryTree {
  constructor (data) {
    this.data = data
    this.left = null
    this.right = null
  }
}

function branchSums (root) {
  if (root === null) return
  return branchSums2(root, 0, [])
}

function branchSums2 (root, sum, output) {
  if (root === null) return
  sum += root.data

  if (root.left === null && root.right === null) {
    output.push(sum)
  }

  branchSums2(root.left, sum, output)
  branchSums2(root.right, sum, output)
  return output
}

const tree = new BinaryTree(23, null, null)
tree.left = new BinaryTree(16, null, null)
tree.right = new BinaryTree(45, null, null)
tree.left.left = new BinaryTree(3, null, null)
tree.left.right = new BinaryTree(22, null, null)
tree.right.left = new BinaryTree(37, null, null)
tree.right.right = new BinaryTree(99, null, null)

console.log('Branch sum is: ', branchSums(tree))
