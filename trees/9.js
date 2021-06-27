const root = Symbol('root')

class Node {
  constructor (data) {
    this.data = data
    this.left = null
    this.right = null
  }
}

class BSTree {
  constructor () {
    this.root = null
  }

  add (data) {
    if (!data) {
      return
    }

    const newNode = new Node(data)

    if (!this.root) {
      this.root = newNode
      return
    }

    const _add = (node, data) => {
      if (data <= node.data) {
        if (!node.left) {
          node.left = newNode
          return
        }
        _add(node.left, data)
      } else {
        if (!node.right) {
          node.right = newNode
          return
        }
        _add(node.right, data)
      }
    }
    _add(this.root, data)
  }

  inorder () {
    const arr = []
    const _print = (node) => {
      if (!node) {
        return
      }
      _print(node.left)
      arr.push(node.data)
      _print(node.right)
    }
    _print(this.root)

    console.log('Inorder of tree is ', arr)
  }

  // Avg Time: O(log(N)), Worst Time: O(N)
  // Avg Space: O(log(N)) or O(d) where d is depth of tree, Worst Space: O(N)
  closestRecursive (target) {
    let diff = Infinity
    let nearest = Infinity

    const _closest = (node, data) => {
      if (node === null) return
      if (Math.abs(node.data - target) <= diff) {
        nearest = node.data
        diff = Math.abs(nearest - target)
      }

      if (target <= node.data) _closest(node.left, target)
      else if (target > node.data) _closest(node.right, target)
      return { diff, nearest }
    }

    const value = _closest(this.root, target)
    console.log('Closest value is', value.nearest)
    console.log('Closest number is', value.diff)
  }

  // Avg Time: O(log(N)), Worst Time: O(N)
  // Avg Space: O(1), Worst Space: O(1) - iterative
  closestIterative (target) {
    let diff = Infinity
    let nearest = Infinity

    const _closest = (node, data) => {
      if (!node) return
      let curent = node
      while (curent !== null) {
        if (Math.abs(curent.data - target) <= diff) {
          nearest = curent.data
          diff = Math.abs(nearest - target)
        }
        if (target < curent.data) curent = curent.left
        else if (target > curent.data) curent = curent.right
        else break
      }
      return { diff, nearest }
    }

    const value = _closest(this.root, target)
    console.log('Closest value is', value.nearest)
    console.log('Closest number is', value.diff)
  }

  // todo
  closest3rdNode (data) {
    let diff = Infinity
    let val = Infinity

    const _closest = (node, data) => {
      if (!node) {
        return
      }

      if (Math.abs(node.data - data) <= diff) {
        val = node.data
        diff = Math.abs(node.data - data)
      }

      if (data <= node.data) {
        _closest(node.left, data)
      } else {
        _closest(node.right, data)
      }

      return { diff, val }
    }

    const value = _closest(this.root, data)
    // console.log("Third closest value is", value.val);
    // console.log("Third closest number is", value.diff);
  }

  kDistantNodes (k) {
    const arr = []
    const _distant = (node, dist) => {
      if (!node) {
        return
      }

      if (dist == 0) {
        arr.push(node.data)
      } else {
        _distant(node.left, dist - 1)
        _distant(node.right, dist - 1)
      }
    }

    _distant(this.root, k)
    console.log(k, 'distant nodes are', arr)
  }
}

const tree = new BSTree()

tree.add(2)
tree.add(4)
tree.add(1)
tree.add(32)
tree.add(12)
tree.add(24)
tree.add(11)
tree.add(3)
tree.add(13)
tree.add(53)
tree.add(23)
tree.add(8)

tree.inorder()
tree.closestRecursive(19)
tree.closestIterative(19)
tree.closest3rdNode(19)
tree.kDistantNodes(2)
