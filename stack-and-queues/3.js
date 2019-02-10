let Stack = {
  head : null,
  tail : null
}

function Node(data, head) {
  this.data = data;
  this.head = head;
}

function push (el) {
  let node = new Node(el, null);
  if(Stack.tail == null){
   Stack.tail = node;   
   Stack.head = node;      
  } else {
    node.head = Stack.head;
    Stack.head  = node;
  }
 }

 function pop () {
  if (Stack.head == null) {
    return "empty Stack"
  } else {
    let out = Stack.head.data;
    Stack.head = Stack.head.head;
    return out;
}
 }


function printList() {
  if(Stack.head == null) {
    process.stdout.write("empty Stack")
  } else {
    current = Stack;
    while(current.head) {
      current = current.head;
      process.stdout.write(String(current.data))
      process.stdout.write("--> ")
    }
    console.log("null")
  }
}

function peek ()  {
  if (Stack.tail !== null) {
    return Stack.tail.data;
  } else {
    return "empty Stack"
  }  
}

function isEmpty()  {
  return Stack.head == null
}

push(1);
push(2);
push(3);
push(4);
push(5);
push(6);
push(7);
push(8);
console.log("final stack ->")
printList()
console.log("POP ->", pop())
console.log("final stack ->")
printList()
console.log("peek the top element in stack", peek())
console.log("Is Stack Empty", isEmpty())
console.log("---------------")