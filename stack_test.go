package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackImplementsLastInFirstOut(t *testing.T) {
	s := Stack{}

	s.Push(1)

	assert.Equal(t, []int{1}, s.items)

	s.Push(2)

	assert.Equal(t, []int{1, 2}, s.items)

	item, _ := s.Pop()

	assert.Equal(t, 2, item)
	assert.Equal(t, []int{1}, s.items)

	item, _ = s.Pop()

	assert.Equal(t, 1, item)
	assert.Equal(t, []int{}, s.items)
}

func TestStackPopErrorsWhenStackEmpty(t *testing.T) {
	s := Stack{}

	item, err := s.Pop()

	assert.Equal(t, -1, item)
	assert.Equal(t, errors.New("stack is empty"), err)
}
