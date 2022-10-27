package main

import "errors"

// A Stack implements last in first out storing of its items.
type Stack[T any] struct {
	items []T
}

// Push adds the item to the Stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes the last item added to the Stack and returns an error when the Stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var emptyValue T
		return emptyValue, errors.New("stack is empty")
	}
	lastIdx := len(s.items) - 1
	poppedItem := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return poppedItem, nil
}
