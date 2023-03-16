package main

import (
	"fmt"
)

func main() {
	// Stack usage example
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop()) // Prints 3
	fmt.Println(stack.Pop()) // Prints 2

	// Queue usage example
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue()) // Prints 1
	fmt.Println(queue.Dequeue()) // Prints 2
}
