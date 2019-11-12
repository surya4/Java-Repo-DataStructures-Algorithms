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
    this.right = null;
  }
  add(data) {
    if (!data) {
      return
    }
    const newNode = new Node(data);

    if (!this.root) {
      this.root = newNode;
      return;
    }

    const _add = (node, data) => {
      if (!node) {
        return
      }
      if (data <= node.data) {
        if (!node.left) {
          node.left = newNode;
          return
        }
        _add(node.left, data);
      } else {
        if (!node.right) {
          node.right = newNode;
          return
        }
        _add(node.right, data);
      }
    }
    _add(this.root, data);
  }

  inorder() {
    let arr = [];
    const _print = (node) => {
      if (!node) {
        return
      }
      _print(node.left);
      arr.push(node.data);
      _print(node.right);
    }
    _print(this.root);
    console.log("Inorder traversal -->", arr);
  }

  leftView() {
    let arr = [];
    if (!this.root) {
      return
    }
    const _view = (node) => {
      if (!node) {
        return
      }
      arr.push(node.data);
      _view(node.left);
      _view(node.right);
    }
    _view(this.root);
    console.log("Left view of tree", arr);
  }

  rightView() {
    let arr = [];
    let curr = 0;
    if (!this.root) {
      return
    }
    const _view = (node, next) => {
      if (!node) {
        return
      }
      console.log("abra--->", { data: node.data, curr, next})
      if (curr < next) {
        curr = next;
        arr.push(node.data);
      }
      // arr.push(node.data);
      _view(node.left, next+1);
      _view(node.right, next+1);
    }
    _view(this.root, 1);
    console.log("Right view of tree", arr);
  }

  ancestorNodes(child) {
    let arr = [];
    const _traverse = (node, child) => {
      if (!node) {
        return
      }

      if (node.data == child) return true;

      if (_traverse(node.left, child) || _traverse(node.right, child)) {
        arr.push(node.data)
        return true;
      }
      return false
    }
    _traverse(this.root, child);
    console.log("Ancestors of node in tree", arr);
  }

  // least common ancestor for BT
  lcaNodes1(n1, n2) {
    const _traverse = (node, n1, n2) => {
      if (!node || !n1 || !n2) {
        return
      }

      if (node.data == n1 || node.data == n2) {
        return node;
      }

      const left = _traverse(node.left, n1, n2);
      const right = _traverse(node.right, n1, n2);

      if (left && right) {
        return node
      }

      if (left) {
        return left;
      } else {
        return right;
      }
    }
    const val = _traverse(this.root, n1, n2);
    console.log("Least common ancestors of node in binary tree", val.data);
  }

    // least common ancestor for BST
    lcaNodes2(n1, n2) {
      const _traverse = (node, n1, n2) => {
        if (!node || !n1 || !n2) {
          return
        }
  
        if (node.data > n1 && node.data > n2) {
          return _traverse(node.left, n1, n2);        
        }
  
        if (node.data < n1 && node.data < n2) {
          return _traverse(node.right, n1, n2);        
        }
  
        return node.data;
      }
      const val = _traverse(this.root, n1, n2);
      console.log("Least common ancestors of node in BST", val);
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

tree.leftView();
tree.rightView();

tree.ancestorNodes(13);
tree.lcaNodes1(13, 53);
tree.lcaNodes2(13, 53);