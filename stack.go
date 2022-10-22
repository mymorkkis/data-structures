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
	poppedItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return poppedItem, nil
}
