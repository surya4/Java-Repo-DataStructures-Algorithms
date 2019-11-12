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
          return;
        }
        _add(node.left, data)
      } else {
        if (!node.right) {
          node.right = newNode;
          return;
        }
        _add(node.right, data)
      }
    }
    _add(this.root, data);

  }

  inorder() {
    let arr = [];
    const _traverse = (node) => {
      if (!node) {
        return
      }
      _traverse(node.left);
      arr.push(node.data);
      _traverse(node.right)
    }
    _traverse(this.root)
    console.log("Inorder traversal is ", arr)
  }

  height(node) {
    const _height = (node) => {
      if (!node) {
        return 0
      }
      let left = 1 + this.height(node.left)
      let right = 1 + this.height(node.right);
      return Math.max(left, right)
    }
    const height = _height(node);
    return height;

  }

  // longest path between two leaves in the tree
  diameter() {
    const _dia = (node) => {
      if (!node) {
        return;
      }

      const left_height = this.height(node.left);
      const right_height = this.height(node.right);
      return left_height + right_height;
    }
    console.log("Diameter of binary tree is ", _dia(this.root))
  }

  // most number of nodes a level
  width() {
    let result = -1;
    const _width = (node) => {
      if (!node) {
        return;
      }

      let queue = [];
      queue.push(node);

      while (queue.length) {
        let count = queue.length;
        result = Math.max(count, result);
        while (count > 0) {
          const front = queue.shift();
          if (front.left) {
            queue.push(front.left)
          }

          if (front.right) {
            queue.push(front.right)
          }
          count--;
        }
      }

      return result;
    }
    console.log("Width of binary tree is ", _width(this.root))
  }

  levelTraverse() {
    let arr = [];
    const _traverse = (node) => {
      if (!node) {
        return;
      }
      let queue = [];
      queue.push(node);

      while (queue.length) {
        const front = queue.shift();
        arr.push(front.data);

        if (front.left) {
          queue.push(front.left);
        }
        
        if (front.right) {
          queue.push(front.right);
        }
      }
    }

    _traverse(this.root);
    console.log("Level order traversal of binary tree is ", arr)
  }

  morrisTraverse() {
    let arr = [];
    const _traverse = (node) => {
      if (!node) {
        return;
      }
    }

    _traverse(this.root);
    console.log("Morris traversal of binary tree is ", arr)
  }
}


const tree = new BSTree();

// tree.add(2);
// tree.add(4);
// tree.add(1);
// tree.add(32);
// tree.add(12);
// tree.add(24);
// tree.add(11);
// tree.add(3);
// tree.add(13);
// tree.add(53);
// tree.add(23);
// tree.add(8);

// test width of tree
tree.add(10)
tree.add(5)
tree.add(15)
tree.add(2)
tree.add(7)
tree.add(12)
tree.add(30)
tree.add(35)

// tree.add(12);
// tree.add(11);
// tree.add(10);
// tree.add(9);
// tree.add(8);
// tree.add(7);
// tree.add(13);
// tree.add(14);
// tree.add(15);
// tree.add(16);
// tree.add(17);

tree.inorder();

tree.width();
tree.diameter();
tree.levelTraverse();
// tree.morrisTraverse();