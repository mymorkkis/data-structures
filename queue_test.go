package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueImplementsFirstInFirstOut(t *testing.T) {
	q := Queue{}

	q.Enqueue(1)

	assert.Equal(t, []int{1}, q.items)

	q.Enqueue(2)

	assert.Equal(t, []int{1, 2}, q.items)

	item, _ := q.Dequeue()

	assert.Equal(t, 1, item)
	assert.Equal(t, []int{2}, q.items)

	item, _ = q.Dequeue()

	assert.Equal(t, 2, item)
	assert.Equal(t, []int{}, q.items)
}

func TestQueueDequeueErrorsWhenQueueEmpty(t *testing.T) {
	q := Queue{}

	item, err := q.Dequeue()

	assert.Equal(t, -1, item)
	assert.Equal(t, errors.New("queue is empty"), err)
}
