const root = Symbol('root');

class Node {
  constructor(data) {
    this.data = data;
    this.right = null;
    this.left = null
  }
}

class BSTree {
  constructor() {
    this.root = null;
  }

  add(data) {
    if (!data) return
    let newNode = new Node(data);

    if (this.root == null) {
      this.root = newNode;
      return
    }

    const _add = (node, data) => {
      if (!node) return
      if (data < node.data) {
        if (!node.left) {
          node.left = newNode;
          return
        }
        _add(node.left, data)
      } else {
        if (!node.right) {
          node.right = newNode;
          return
        }
        _add(node.right, data)
      }
    }
    _add(this.root, data)
  }

  inorder() {
    let arr = [];
    const _print = (node) => {
      if (!node) return
      _print(node.left);
      arr.push(node.data);
      _print(node.right);
    }
    _print(this.root);
    console.log("Inorder is -> ", arr)
  }

  height() {
    const _height = (node) => {
      if (!node) return 0;
      const left_height = 1 + _height(node.left);
      const right_height = 1 + _height(node.right);
      return Math.max(left_height, right_height);
    }
    const h = _height(this.root);
    console.log("Max height of tree is -> ", h)
  }

  minHeight() {
    const _height = (node) => {
      if (!node) return 0;
      const left_height = 1 + _height(node.left);
      const right_height = 1 + _height(node.right);
      return Math.min(left_height, right_height);
    }
    const h = _height(this.root);
    console.log("Min height of tree is -> ", h)
  }

  isHeightBalanced() {
    const _isBalanced = (node, isBalanced) => {
      if (!node) return 0;
      const left_height = 1 + _isBalanced(node.left, isBalanced);
      const right_height = 1 + _isBalanced(node.right, isBalanced);

      if (Math.abs(left_height - right_height) >= 1) {
        isBalanced = false
      }

      return isBalanced;

    }
    const bal = _isBalanced(this.root, true);
    console.log("Is tree height balanced -> ", bal)
  }
}

const tree = new BSTree();

tree.add(2);
tree.add(4);
tree.add(1);
tree.add(32);
tree.add(12);
tree.add(24);
tree.add(11);
tree.add(3);
tree.add(13);
tree.add(53);
tree.add(23);
tree.add(8);

tree.inorder();
tree.height();
tree.minHeight();
tree.isHeightBalanced();




























































// let tree = {
//   root : null
// }

// function Node(data, left, right) {
//   this.data = data;
//   this.left = left;
//   this.right = right;
// }

// tree.root = new Node(23, null, null);
// tree.root.left = new Node(16, null, null);
// tree.root.right = new Node(45, null, null);
// tree.root.left.left = new Node(3, null, null);
// tree.root.left.right = new Node(22, null, null);
// tree.root.right.left = new Node(37, null, null);
// tree.root.right.right = new Node(99, null, null);

// function treeHeight(node) {
//   if (node == null) {
//     return 0;
//   } else {
//     let leftH = treeHeight(node.left);
//     let rightH = treeHeight(node.right);
//     return Math.max((leftH, rightH) + 1);

//   }
// }

// console.log("Heigth opf Tree", treeHeight(tree.root));