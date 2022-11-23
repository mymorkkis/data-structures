package main

import "errors"

// A Stack implements last in first out storing of its items.
type Stack[T comparable] struct {
	Length int
	head   *node[T]
}

// Push adds the item to the Stack.
func (s *Stack[T]) Push(item T) {
	n := node[T]{value: item, nextNode: s.head}
	s.head = &n
	s.Length++
}

// Pop removes the last item added to the Stack and returns an error when the Stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.Length == 0 {
		var emptyValue T
		return emptyValue, errors.New("stack is empty")
	}

	head := s.head
	s.head = head.nextNode
	s.Length--

	return head.value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Length == 0 {
		var emptyValue T
		return emptyValue, errors.New("stack is empty")
	}

	return s.head.value, nil
}
