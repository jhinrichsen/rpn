package main

import (
	"fmt"
	"os"
	"strconv"
)

// Stack implementation for integer values
type Stack struct {
	pointer int
	entries []int
}

// Depth returns number of entries in stack
func (stack *Stack) Depth() int {
	return stack.pointer
}

// Pop a value from the stack
func (stack *Stack) Pop() int {
	stack.pointer--
	n := stack.entries[stack.pointer]
	return n
}

// Push a value onto the stack
func (stack *Stack) Push(n int) {
	// Need to grow stack?
	if stack.pointer < len(stack.entries) {
		stack.entries[stack.pointer] = n
	} else {
		stack.entries = append(stack.entries, n)
	}
	stack.pointer++
}

func main() {
	stack := Stack{}
	for _, arg := range os.Args[1:] {
		switch arg {
		case "+":
			stack.Push(stack.Pop() + stack.Pop())
		case "-":
			n1 := stack.Pop()
			n2 := stack.Pop()
			stack.Push(n2 - n1)
		case "*":
			stack.Push(stack.Pop() * stack.Pop())
		case "/":
			n1 := stack.Pop()
			n2 := stack.Pop()
			stack.Push(n2 / n1)
		default:
			n, _ := strconv.Atoi(arg)
			stack.Push(n)
		}
	}
	if stack.Depth() > 0 {
		fmt.Println(stack.Pop())
	}
	if stack.Depth() > 0 {
		fmt.Printf("warning: want depth=0 but got %v\n", stack.Depth())
	}
}
