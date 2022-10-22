package main

import "errors"

type Queue struct {
	items []int
}

func (queue *Queue) Enqueue(item int) {
	queue.items = append(queue.items, item)
}

func (queue *Queue) Dequeue() (int, error) {
	if len(queue.items) == 0 {
		return -1, errors.New("queue is empty")
	}
	poppedItem := queue.items[0]
	queue.items = queue.items[1:]
	return poppedItem, nil
}
