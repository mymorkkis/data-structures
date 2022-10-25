package main

import "errors"

// A Queue implements first in first out storing of its items.
type Queue[T any] struct {
	items []T
}

// Enqueue adds the item to the Queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes the first item added to the Queue and returns an error when the Queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		var emptyValue T
		return emptyValue, errors.New("queue is empty")
	}
	poppedItem := q.items[0]
	q.items = q.items[1:]
	return poppedItem, nil
}
