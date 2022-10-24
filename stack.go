package main

import "errors"

// A Stack implements last in first out storing of its items.
type Stack struct {
	items []int
}

// Push adds the item to the Stack.
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes the last item added to the Stack and returns an error when the Stack is empty.
func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("stack is empty")
	}
	lastIdx := len(s.items) - 1
	poppedItem := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return poppedItem, nil
}
