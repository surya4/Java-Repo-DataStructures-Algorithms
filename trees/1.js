let BTree = {
  root : null
}

function Node(data, left, right) {
  this.data = data;
  this.left = left;
  this.right = right;
}

function inOrderTraverse(node) {
  if (node == null) return;
  inOrderTraverse(node.left); 
  process.stdout.write(node.data + " ")
  inOrderTraverse(node.right); 
}

function preOrderTraverse(node) {
  if (node == null) return;
  process.stdout.write(node.data + " ")
  preOrderTraverse(node.left);  
  preOrderTraverse(node.right); 
}

function postOrderTraverse(node) {
  if (node == null) return;
  postOrderTraverse(node.left);  
  postOrderTraverse(node.right);
  process.stdout.write(node.data + " ")
}

BTree.root = new Node(23, null, null);
BTree.root.left = new Node(16, null, null);
BTree.root.right = new Node(45, null, null);
BTree.root.left.left = new Node(3, null, null);
BTree.root.left.right = new Node(22, null, null);
BTree.root.right.left = new Node(37, null, null);
BTree.root.right.right = new Node(99, null, null);

process.stdout.write("Inorder traversal: "); 
inOrderTraverse(BTree.root); 
console.log();

process.stdout.write("PreOrder traversal: "); 
preOrderTraverse(BTree.root); 
console.log();

process.stdout.write("PostOrder traversal: "); 
postOrderTraverse(BTree.root); 
console.log();