package main

import (
	"fmt"

	str "github.com/alex-necsoiu/link/datastructures"
)

func main() {

	fmt.Println("I tu padre que tal?")
	// firststack := d.Stack
	// Stack usage example
	stack := &str.Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop()) // Prints 3
	fmt.Println(stack.Pop()) // Prints 2

	// Queue usage example
	queue := &str.Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue()) // Prints 1
	fmt.Println(queue.Dequeue()) // Prints 2
}
