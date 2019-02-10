let tree = {
  root : null
}

function Node(data, left, right) {
  this.data = data;
  this.left = left;
  this.right = right;
}

function insert(data) {
  let node = new Node(data, null, null);
  if (tree.root == null) {
    tree.root = node;
  } else {
    let current = tree.root;
    let parent = current;
    while (true){ 
      parent = current; 
      if (data < current.data) { 
          current = current.left; 
          if (current == null){ 
              parent.left = node; 
              break; 
          } 
      } else { 
          current = current.right; 
          if (current == null){ 
              parent.right = node; 
              break; 
          } 
      } 
    }
  }
}

function isEmpty() {
  return arr.length == -1;
}

function treeHeight(node) {
  if (node == null) {
    return 0;
  } else {
    let leftH = treeHeight(node.left);
    let rightH = treeHeight(node.right);
    return Math.max((leftH, rightH) + 1);

  }
}

function avlMap(node) {
  let leftH = treeHeight(node.left);
  let rightH = treeHeight(node.right);
  let diff = leftH - rightH;
  if (diff < -1) {
    
  } else if (diff > 1) {
    
  } else {

  }
}

function leftRotate() {
  
}

function rightRotate() {
  
}

insert(5);
insert(4);
insert(8);
insert(6);
insert(1);
insert(14);
insert(2);
insert(7);

console.log("final avl tree", arr);
// remove(6);
console.log("final avl tree after remove 6", arr);
console.log("is avl tree empty", isEmpty());