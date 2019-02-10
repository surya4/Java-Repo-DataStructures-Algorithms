// Implement by not using inbuilt JS functions
let Stack1 = function () {
  this.array = [];
  let top = -1;

  this.isEmpty = () => {
    return top == -1;
  }

  this.push = (el) => {
    this.array[++top] = el;
  }

  this.pop = () => {
    let output = this.array[top];
    this.array.splice(top--);
    return output;
  }

  this.peek = () => {
    if (top != -1) {
      return this.array[top];
    } else {
      return "empty stack";
    }    
  }
}

let stack1 = new Stack1();

stack1.push(1);
stack1.push(2);
stack1.push(3);
stack1.push(4);
stack1.push(5);
stack1.push(6);
stack1.push(7);
stack1.push(8);
console.log("final stack", stack1.array)
console.log( stack1.pop(), "final stack after pop -> ", stack1.array)
console.log("peek the top element in stack", stack1.peek())
console.log("---------------")

// Implement by using inbuilt JS functions
let Stack2 = [];

Stack2.push(1);
Stack2.push(2);
Stack2.push(3);
Stack2.push(4);
Stack2.push(5);
Stack2.push(6);
Stack2.push(7);
Stack2.push(8);
console.log("final stack", Stack2)
console.log( Stack2.pop(), "final stack after pop -> ", Stack2)
console.log("peek the top element in stack", Stack2[Stack2.length-1])