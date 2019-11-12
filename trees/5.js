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
    this.root = null;
  }

  add (data) {

    if (!data) return;
    let newNode = new Node(data);
    if (this.root == null) {
      this.root = newNode;
      return
    }

    const _add = (node, data) => {
      if (data <= node.data) {
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
    _add(this.root, data);
  }

  inorder () {
    const arr = []
    const _inorder = (node) => {
      if (!node) return;
      _inorder(node.left)
      arr.push(node.data)
      _inorder(node.right)
    }
    _inorder(this.root);
    console.log("Inorder -->", arr);
  }

  // leafCountIter () {
  //   let count = 0;
  //   const _count = (node) => {
  //     if (!node) return
  //     if (!node.left && !node.right) {
  //       count++;
  //     }
      
  //     _count(node.left);
  //     _count(node.right);
  //   }

  //   _count(this.root);
  //   console.log("Leaf node count -> ", count)
  // }

  leafCountRecur() {
    const _count = (node) => {
      if (!node) return 0;
      if (!node.left && !node.right) return 1;
      return _count(node.left) + _count(node.right);
    }

    const count = _count(this.root);
    console.log("Leaf node count -> ", count)
  }

  nonLeafCountRecur() {
    const _count = (node) => {
      if (!node) return 0;
      if (!node.left && !node.right) return 0;
      return 1 + _count(node.left) + _count(node.right);
    }

    const count = _count(this.root);
    console.log("Non leaf node count -> ", count)
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
// tree.leafCountIter();
tree.leafCountRecur();
tree.nonLeafCountRecur();