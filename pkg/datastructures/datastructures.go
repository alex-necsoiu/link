package datastructures

// Stack implementation
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

// Queue implementation
type Queue struct {
	items []int
}

// Enqueue adds an item to the rear of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}

	front := q.items[0]
	q.items = q.items[1:]
	return front
}
