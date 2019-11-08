const head = Symbol("head");

class Node {
  constructor(data) {
    this.data = data;
    this.next = null;
  }
}

class LinkedList {
  constructor() {
    this[head] = null;
  }
  add(data) {
    const newNode = new Node(data);
    if (this[head] === null) {
      this[head] = newNode;
    } else {
      let current = this[head];
      while (current.next != null) {
        current = current.next;
      }
      current.next = newNode;
    }
  }
  print() {
    if (this[head] === null) {
      console.log("Empty List");
      return
    }
    let current = this[head];
      while (current != null) {
        console.log(current.data, "->");
        current = current.next;
      }

  }
  reverseList() {
    if (this[head] === null) {
      console.log("Empty List");
      return
    }
    let prev = null;
    let temp = null;
    let current = this[head];

    while (current != null) {
      temp = current.next;
      current.next = prev;
      prev = current;
      current = temp;
    }

    let newHead = prev;
    while (newHead != null) {
      console.log(newHead.data, "->");
      newHead = newHead.next;
    }

  }

}

const list = new LinkedList();

list.add(1);
list.add(2);
list.add(3);
list.add(4);
list.print();
list.reverseList();
