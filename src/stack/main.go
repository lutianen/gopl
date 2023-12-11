package main

import "fmt"

func main() {
	stack := make([]int, 0)

	for i := 0; i < 10; i++ {
		stack = push(stack, i)
	}
	displayStack(stack)
	fmt.Println(top(stack))

	stack = pop(stack)
	displayStack(stack)

	stack = remove(stack, 0)
	displayStack(stack)

	stack = remove(stack, 10)
	displayStack(stack)

	stack = remove(stack, len(stack)-1)
	displayStack(stack)
}

func push(stack []int, value int) []int {
	stack = append(stack, value)
	return stack
}

func top(stack []int) int {
	return stack[len(stack)-1]
}

func pop(stack []int) []int {
	stack = stack[:len(stack)-1]
	return stack
}

func remove(stack []int, index int) []int {
	if index >= len(stack) {
		return stack
	}

	copy(stack[index:], stack[index+1:])
	return stack[:len(stack)-1]
}

func displayStack(stack []int) {
	fmt.Printf("[ ")
	for _, val := range stack {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("]\n")
}
