package main

import "errors"

type Stack struct {
	items []int
}

func (stack *Stack) Push(item int) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() (int, error) {
	if len(stack.items) == 0 {
		return -1, errors.New("stack is empty")
	}
	lastIdx := len(stack.items) - 1
	poppedItem := stack.items[lastIdx]
	stack.items = stack.items[:lastIdx]
	return poppedItem, nil
}
