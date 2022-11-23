package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackImplementsLastInFirstOut(t *testing.T) {
	s := Stack[int]{}

	s.Push(1)
	assert.Equal(t, 1, s.Length)

	head, _ := s.Peek()
	assert.Equal(t, 1, head)

	s.Push(2)
	assert.Equal(t, 2, s.Length)

	head, _ = s.Peek()
	assert.Equal(t, 2, head)

	item, _ := s.Pop()
	assert.Equal(t, 2, item)
	assert.Equal(t, 1, s.Length)

	head, _ = s.Peek()
	assert.Equal(t, 1, head)

	item, _ = s.Pop()
	assert.Equal(t, 1, item)
	assert.Equal(t, 0, s.Length)
}

func TestStackPopErrorsWhenStackEmpty(t *testing.T) {
	s := Stack[int]{}

	item, err := s.Pop()

	assert.Empty(t, item)
	assert.Equal(t, errors.New("stack is empty"), err)
}

func TestStackPeekErrorsWhenStackEmpty(t *testing.T) {
	s := Stack[int]{}

	item, err := s.Peek()

	assert.Empty(t, item)
	assert.Equal(t, errors.New("stack is empty"), err)
}
