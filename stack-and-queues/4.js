let Queue = {
  head : null,
  tail : null
}

function Node(data, head) {
  this.data = data;
  this.head = head;
}

function enqueue (el) {
 let node = new Node(el, null);
 if(Queue.tail == null){
  Queue.tail = node;   
  Queue.head = node;      
 } else {
  Queue.tail.head = node;
  Queue.tail = Queue.tail.head; 
 }
}

function dequeue () {
  if (Queue.head !== null) {
    var out = Queue.head;
    Queue.head = Queue.head.head; 
    return out.data;
  } else {
    return "empty Queue"
  }
}

function printList() {
  if(Queue.head == null) {
    process.stdout.write("empty Queue")
  } else {
    current = Queue;
    while(current.head) {
      current = current.head;
      process.stdout.write(String(current.data))
      process.stdout.write("--> ")
    }
    console.log("null")
  }
}

function peek ()  {
  if (Queue.head !== null) {
    return Queue.head.data;
  } else {
    return "empty Queue"
  }  
}

function isEmpty()  {
  return Queue.tail == null
}

enqueue(1);
enqueue(2);
enqueue(3);
enqueue(4);
enqueue(5);
enqueue(6);
enqueue(7);
enqueue(8);
console.log("final Queue ->")
printList()
console.log("dequeue ->", dequeue())
console.log("final Queue ->")
printList()
console.log("peek the top element in Queue", peek())
console.log("Is Queue Empty", isEmpty())
console.log("---------------")