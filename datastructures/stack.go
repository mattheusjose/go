package main

import (
	"fmt"
)

type Stack struct {
	data   []int
	length int
	top    int
}

func Push(stack *Stack, element int) {

	if IsFull(stack) {
		panic("Stack is full")
	}

	position := stack.top
	stack.data[position] = element
	stack.top += 1
}

func Pop(stack *Stack) int {

	if IsEmpty(stack) {
		panic("Stack is empty")
	}

	position := stack.top
	element := stack.data[position]
	stack.data[position] = 0
	stack.top -= 1
	return element
}

func IsFull(stack *Stack) bool {

	isFull := stack.top == stack.length-1
	return isFull
}

func IsEmpty(stack *Stack) bool {

	isEmpty := stack.top == 0
	return isEmpty
}

func PrintStack(stack *Stack) {

	for index := stack.top - 1; index >= 0; index-- {
		element := stack.data[index]
		fmt.Print(fmt.Sprintf("[%v]", element))
	}
}

func main() {

	length := 10
	stack := Stack{
		data:   make([]int, length),
		length: length,
		top:    0,
	}

	Push(&stack, 10)
	Push(&stack, 20)
	Push(&stack, 30)
	Pop(&stack)
	Pop(&stack)
	Push(&stack, 40)

	PrintStack(&stack)
}
