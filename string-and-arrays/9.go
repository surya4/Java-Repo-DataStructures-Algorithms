package main

import "fmt"

var top = -1
var stack = make([]string, 10)

func push(val string) {
	top++
	stack[top] = val
}

func pop() string {
	var topValue = stack[top]
	stack = stack[:top]
	top--
	return topValue
}

func isEmpty(len int) bool {
	return top == -1
}

func peek() string {
	if top != -1 {
		return stack[top]
	}
	return "Empty"
}

func parethesisFunc(str string) bool {
	isBalanced := false
	topElem
	for i, r := range str {
		// fmt.Println("r", i, string(r))
		switch string(r) {
		case "{":
			push(string(r))

		case "(":
			push(string(r))

		case "[":
			push(string(r))

		case "}":
			pop()
			topElem = peek()
			fmt.Println("r", topElem, stack[top])
			// if topElem ==  {

			// }

		case ")":
			pop()
			topElem = peek()
			fmt.Println("r", topElem, stack[top])

		case "]":
			pop()
			topElem = peek()
			fmt.Println("r", topElem, stack[top])
		}

		// fmt.Println(i, string(r))
	}

	fmt.Println("stack", stack)

	return isBalanced
}

func main() {
	input1 := "{([])}"
	input2 := "([]"
	fmt.Println(input1, " output -->", parethesisFunc(input1))
	fmt.Println(input2, " output -->", parethesisFunc(input2))
}
