const root = Symbol('root');

class Node {
  constructor(data) {
    this.data = data;
    this.left = null;
    this.right = null;
  }
}

class bTree {
  constructor() {
    this[root] = null;
  }

  inOrderTraverse(node) {
    if (node == null) return;
    this.inOrderTraverse(node.left); 
    process.stdout.write(node.data + " ")
    this.inOrderTraverse(node.right); 
  }

  preOrderTraverse(node) {
    if (node == null) return;
    process.stdout.write(node.data + " ")
    this.preOrderTraverse(node.left);  
    this.preOrderTraverse(node.right); 
  }

  postOrderTraverse(node) {
    if (node == null) return;
    this.postOrderTraverse(node.left);  
    this.postOrderTraverse(node.right);
    process.stdout.write(node.data + " ")
  }

}

let btree = new bTree();

let tree = new Node(23, null, null);
tree.left = new Node(16, null, null);
tree.right = new Node(45, null, null);
tree.left.left = new Node(3, null, null);
tree.left.right = new Node(22, null, null);
tree.right.left = new Node(37, null, null);
tree.right.right = new Node(99, null, null);

process.stdout.write("Inorder traversal: "); 
btree.inOrderTraverse(tree); 
console.log();

process.stdout.write("PreOrder traversal: "); 
btree.preOrderTraverse(tree); 
console.log();

process.stdout.write("PostOrder traversal: "); 
btree.postOrderTraverse(tree); 
console.log();