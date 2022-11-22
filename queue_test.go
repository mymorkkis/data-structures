package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueImplementsFirstInFirstOut(t *testing.T) {
	q := Queue[int]{}
	assert.Equal(t, 0, q.Length)

	q.Enqueue(1)
	assert.Equal(t, 1, q.Length)

	head, _ := q.Peek()
	assert.Equal(t, 1, head)

	q.Enqueue(2)
	assert.Equal(t, 2, q.Length)

	head, _ = q.Peek()
	assert.Equal(t, 1, head)

	item, _ := q.Dequeue()
	assert.Equal(t, 1, item)
	assert.Equal(t, 1, q.Length)

	head, _ = q.Peek()
	assert.Equal(t, 2, head)

	item, _ = q.Dequeue()
	assert.Equal(t, 2, item)
	assert.Equal(t, 0, q.Length)
}

func TestQueueDequeueErrorsWhenQueueEmpty(t *testing.T) {
	q := Queue[int]{}

	item, err := q.Dequeue()

	assert.Empty(t, item)
	assert.Equal(t, errors.New("queue is empty"), err)
}

func TestQueuePeekErrorsWhenQueueEmpty(t *testing.T) {
	q := Queue[int]{}

	item, err := q.Peek()

	assert.Empty(t, item)
	assert.Equal(t, errors.New("queue is empty"), err)
}
