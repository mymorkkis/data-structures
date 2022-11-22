package main

import "errors"

// A Queue implements first in first out storing of its items.
type Queue[T comparable] struct {
	Length int
	head   *node[T]
	tail   *node[T]
}

// Enqueue adds the item to the Queue.
func (q *Queue[T]) Enqueue(item T) {
	n := node[T]{value: item}

	if q.Length == 0 {
		q.head = &n
	} else {
		q.tail.nextNode = &n
	}

	q.tail = &n
	q.Length++
}

// Dequeue removes the first item added to the Queue and returns an error when the Queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.Length == 0 {
		var emptyValue T
		return emptyValue, errors.New("queue is empty")
	}

	head := q.head
	q.head = q.head.nextNode

	q.Length--
	if q.Length == 0 {
		q.tail = nil
	}

	return head.value, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.Length == 0 {
		var emptyValue T
		return emptyValue, errors.New("queue is empty")
	}
	return q.head.value, nil
}
