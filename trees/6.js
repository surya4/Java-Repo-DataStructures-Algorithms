let tree = {
  root : null
}

function Node(data, left, right) {
  this.data = data;
  this.left = left;
  this.right = right;
}

tree.root = new Node(23, null, null);
tree.root.left = new Node(16, null, null);
tree.root.right = new Node(45, null, null);
tree.root.left.left = new Node(3, null, null);
tree.root.left.right = new Node(22, null, null);
tree.root.right.left = new Node(37, null, null);
tree.root.right.right = new Node(99, null, null);

function treeHeight(node) {
  if (node == null) {
    return 0;
  } else {
    let leftH = treeHeight(node.left);
    let rightH = treeHeight(node.right);
    return Math.max((leftH, rightH) + 1);

  }
}

console.log("Heigth opf Tree", treeHeight(tree.root));