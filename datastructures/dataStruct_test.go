package datastructures

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	printStack(*stack)
	if value := stack.Pop(); value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}

	printStack(*stack)
	if value := stack.Pop(); value != 2 {
		t.Errorf("Expected 2, got %d", value)

	}

	stack.Push(4)
	printStack(*stack)

	if value := stack.Pop(); value != 4 {
		t.Errorf("Expected 4, got %d", value)
		printStack(*stack)

	}
	printStack(*stack)

	if value := stack.Pop(); value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	printStack(*stack)
	// Test empty stack
	if value := stack.Pop(); value != -1 {
		t.Errorf("Expected -1 (empty stack), got %d", value)
	}
}
func printStack(str Stack) {
	fmt.Print("\n")
	for i := range str.items {
		fmt.Print("Stack Item : ", str.items[i], "\n")
	}
}

func TestQueue(t *testing.T) {
	queue := &Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	printQueue(*queue)
	if value := queue.Dequeue(); value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}
	printQueue(*queue)

	if value := queue.Dequeue(); value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}
	printQueue(*queue)

	queue.Enqueue(4)

	if value := queue.Dequeue(); value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
	printQueue(*queue)

	if value := queue.Dequeue(); value != 4 {
		t.Errorf("Expected 4, got %d", value)
	}
	printQueue(*queue)

	// Test empty queue
	if value := queue.Dequeue(); value != -1 {
		t.Errorf("Expected -1 (empty queue), got %d", value)
	}
}
func printQueue(que Queue) {
	fmt.Print("\n")
	for i := range que.items {
		fmt.Print("Queue Item : ", que.items[i], "\n")
	}
}
