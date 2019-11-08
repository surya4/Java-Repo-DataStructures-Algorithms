let root = Symbol('root');

class Node {
  constructor(data) {
    this.data = data;
    this.left = null;
    this.right = null;
  }
}

class BSTree {
  constructor() {
    this[root] = null;
  }

  addIter(data) {
    if (data == null) return
    let newNode = new Node(data);
    if (this.root == null) {
      this.root = newNode;
    } else {
     let current = this.root;
     let parent = null;
     while (current != null) {
      parent = current;
       if (data < current.data) {
        current = current.left;
        if (current == null) {
          parent.left = newNode;
          break
        }
       } else {
        current = current.right;
        if (current == null) {
          parent.right = newNode;
          break
        }
       }
     } 
    }
  }

  addRecur(node, data) {
    if (data == null) return

    let newNode = new Node(data);
    if (this.root == null) {
      this.root = newNode;
      return
    }

    if (data <= node.data) {
      if (node.left == null) {
        node.left = newNode;
        return
      }
      this.addRecur(node.left, data);
    } else {
      if (node.right == null) {
        node.right = newNode;
        return
      }
      this.addRecur(node.right, data);
    }

  }

  // delIter(data) {
  //   if (data == null) return
  //   let newNode = new Node(data);
  //   if (this.root == null) {
  //     this.root = newNode;
  //   } else {

  //   }
  // }

  // delRecur(node, data) {
  //   if (data == null) return

  //   let newNode = new Node(data);
  //   if (this.root == null) {
  //     this.root = newNode;
  //     return
  //   }

  // }

  preOrderRecur() {
    const arr = [];
    const preOrder = (node) => {
      if (!node) return
      arr.push(node.data);
      preOrder(node.left);
      preOrder(node.right);
    }
    preOrder(this.root);
    console.log("Preorder traversal: ", arr)
  }

  postOrderRecur() {
    const arr = [];
    const postOrder = (node) => {
      if (!node) return
      postOrder(node.left);
      postOrder(node.right);
      arr.push(node.data);
    }
    postOrder(this.root);
    console.log("Postorder traversal: ", arr)
  }

  inOrderRecur() {
    const arr = [];
    const inOrder = (node) => {
      if (!node) return
      inOrder(node.left);
      arr.push(node.data);
      inOrder(node.right);
    }
    inOrder(this.root);
    console.log("Inorder traversal: ", arr)
    return arr;
  }

  // inOrderIter(node) {
  //   if (node == null) return
  //   process.stdout.write(node.data + " -> ")
  // }

  // preOrderIter(node) {
  //   if (node == null) return
  //   process.stdout.write(node.data + " -> ")
  // }

  // postOrderIter(node) {
  //   if (node == null) return
  //   process.stdout.write(node.data + " -> ")
  // }
}

const tree1 = new BSTree();
const tree2 = new BSTree();

tree1.addIter(23); 
tree1.addIter(45); 
tree1.addIter(16);
tree1.addIter(37);
tree1.addIter(3);
tree1.addIter(99);
tree1.addIter(22);
tree1.addIter(5);
tree1.addIter(12);
tree1.addIter(67);
tree1.addIter(34);

tree2.addRecur(tree2.root,23); 
tree2.addRecur(tree2.root,45); 
tree2.addRecur(tree2.root,16);
tree2.addRecur(tree2.root,37);
tree2.addRecur(tree2.root,2);
tree2.addRecur(tree2.root,99);
tree2.addRecur(tree2.root,22);
tree2.addRecur(tree2.root,5);
tree2.addRecur(tree2.root,12);
tree2.addRecur(tree2.root,67);
tree2.addRecur(tree2.root,34);

// tree1.delIter(34);
// tree2.delRecur(tree2.root,34);

console.log("-- Iterative --")
tree1.inOrderRecur(tree1.root);
tree1.preOrderRecur(tree1.root);
tree1.postOrderRecur(tree1.root);

console.log("-- Recursive --")
tree2.inOrderRecur(tree2.root);
tree2.preOrderRecur(tree2.root);
tree2.postOrderRecur(tree2.root);