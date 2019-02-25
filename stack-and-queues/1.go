package main

import "fmt"

var capacity = 10

var stack = make([]int, capacity)

var top = -1

func push(inp int) {
	top++
	stack[top] = inp
}

func pop() int {
	stack = stack[:top]
	top--
	return 2
}

func isEmpty() bool {
	return top == -1
}

func main() {
	fmt.Println("is stack empty", isEmpty())
	input := [10]int{51, 13, 41, 11, 21, 17, 81, 91, 4}
	for i := 0; i < len(input); i++ {
		push(input[i])
	}
	pop()
	fmt.Println("is stack empty", isEmpty())

}
