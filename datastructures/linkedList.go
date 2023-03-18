package datastructures

import "fmt"

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(value string) {
	newNode := &Node{value: value, next: nil}

	if l.head == nil {
		l.head = newNode
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%s -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (l *LinkedList) Remove(value string) {
	if l.head == nil {
		return
	}
	if l.head.value == value {
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}
