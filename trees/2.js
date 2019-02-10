let BSTree = {
  root : null
}

function Node(data, left, right) {
  this.data = data;
  this.left = left;
  this.right = right;
}

function insert(data) {
  let node = new Node(data, null, null);
  if (BSTree.root == null) {
    BSTree.root = node;
  } else {
    let current = BSTree.root;
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

// function isTreeBinary() {
  
// }

insert(23); 
insert(45); 
insert(16); 
insert(37); 
insert(3); 
insert(99); 
insert(22);

process.stdout.write("Inorder traversal: "); 
inOrderTraverse(BSTree.root); 
console.log();

process.stdout.write("PreOrder traversal: "); 
preOrderTraverse(BSTree.root); 
console.log();

process.stdout.write("PostOrder traversal: "); 
postOrderTraverse(BSTree.root); 
console.log();