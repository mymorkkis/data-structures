package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	stack := Stack{}

	stack.Push(1)

	assert.Equal(t, []int{1}, stack.items)

	stack.Push(2)

	assert.Equal(t, []int{1, 2}, stack.items)
}

func TestStackPop(t *testing.T) {
	stack := Stack{items: []int{1, 2}}

	item, _ := stack.Pop()

	assert.Equal(t, 2, item)
	assert.Equal(t, []int{1}, stack.items)

	item, _ = stack.Pop()

	assert.Equal(t, 1, item)
	assert.Equal(t, []int{}, stack.items)
}

func TestStackPopErrorsWhenStackEmpty(t *testing.T) {
	stack := Stack{}

	item, err := stack.Pop()

	assert.Equal(t, -1, item)
	assert.Equal(t, errors.New("stack is empty"), err)
}
