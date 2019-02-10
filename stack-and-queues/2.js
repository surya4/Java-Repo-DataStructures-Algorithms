// Implement by not using inbuilt JS functions
let Queue1 = function () {
  this.array = [];
  let top = -1;

  this.isEmpty = () => {
    return top == -1;
  }

  this.enqueue = (el) => {
    this.array[++top] = el;
  }

  this.dequeue = () => {
    if (top != -1) {
      this.array.splice(top--);
      return this.array[0];
    } else {
      return "empty Queue";
    }
  }

  this.peek = () => {
    if (top != -1) {
      return this.array[top];
    } else {
      return "empty Queue";
    }
  }
}

let myQueue = new Queue1();

myQueue.enqueue(1);
myQueue.enqueue(2);
myQueue.enqueue(3);
myQueue.enqueue(4);
myQueue.enqueue(5);
myQueue.enqueue(6);
myQueue.enqueue(7);
myQueue.enqueue(8);
console.log("final Queue", myQueue.array)
console.log( myQueue.dequeue(), "final Queue after dequeue -> ", myQueue.array)
console.log("peek the top element in Queue", myQueue.peek())
console.log("---------------")

// Implement by using inbuilt JS functions
let Queue2 = [];

Queue2.enqueue(1);
Queue2.enqueue(2);
Queue2.enqueue(3);
Queue2.enqueue(4);
Queue2.enqueue(5);
Queue2.enqueue(6);
Queue2.enqueue(7);
Queue2.enqueue(8);
console.log("final Queue", Queue2)
console.log( Queue2.shift(), "final Queue after dequeue -> ", Queue2)
console.log("peek the top element in Queue", Queue2[Queue2.length-1])
console.log("---------------")