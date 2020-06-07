const LL1 = function () {
  this.head = null

  this.push = (el) => {
    const node = {
      value: el,
      next: null,
      prev: null
    }
    if (!this.head) {
      this.head = node
    } else {
      current = this.head
      while (current.next) {
        current = current.next
      }
      current.next = node
    }
  }

  this.deleteFromHead = () => {
    let current = this.head
    current = current.next
    return current
  }

  this.deleteFromTail = () => {
    let current = this.head
    while (current.next.next) {
      current = current.next
    }
    current.next = null
  }

  this.deleteByIndex = (index) => {
    let count = 1
    let current = this.head
    while (current.next) {
      ++count
      process.stdout.write(String(current.value))
      process.stdout.write('--> ')
      if (index == count && current.next.next) {
        current = current.next.next
      } else {
        current = current.next
      }
    }
    console.log('null')
  }

  this.deleteByValue = (el) => {
    let current = this.head
    while (current.next) {
      if (current.next.value == el && current.next.next) {
        current = current.next.next
      } else {
        current = current.next
      }
    }
  }

  this.printList = () => {
    if (!this.head) {
      process.stdout.write('empty list')
    } else {
      current = this.head
      while (current.next) {
        process.stdout.write(String(current.value))
        process.stdout.write('--> ')
        current = current.next
      }
      console.log('null')
    }
  }

  this.printListByValue = (el) => {
    if (!this.head) {
      process.stdout.write('empty list')
      // return
    } else {
      current = this.head
      while (current.next) {
        if (current.value == el) {
          break
        } else {
          current = current.next
        }
      }
      while (current.next) {
        process.stdout.write(String(current.value))
        process.stdout.write('--> ')
        current = current.next
      }
      console.log('null')
    }
  }
}

const ll1 = new LL1()

ll1.push(1)
ll1.push(2)
ll1.push(3)
ll1.push(4)
ll1.push(5)
ll1.push(6)
ll1.push(7)
ll1.push(8)
console.log('final linkedlist', ll1.printList())
const val = ll1.deleteFromHead()
console.log('current head after delte head node', val.value)

console.log('list after deleteFromHead -> ')
ll1.printListByValue(val.value)
console.log('list after deleteFromTail -> ')
ll1.deleteFromTail()
console.log('list after deleteByIndex -> ')
ll1.deleteByIndex(2)
console.log('list after deleteByValue -> ')
ll1.deleteByValue(6)
console.log('---------------')
