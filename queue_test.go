package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueImplementsFirstInFirstOut(t *testing.T) {
	queue := Queue{}

	queue.Enqueue(1)

	assert.Equal(t, []int{1}, queue.items)

	queue.Enqueue(2)

	assert.Equal(t, []int{1, 2}, queue.items)

	item, _ := queue.Dequeue()

	assert.Equal(t, 1, item)
	assert.Equal(t, []int{2}, queue.items)

	item, _ = queue.Dequeue()

	assert.Equal(t, 2, item)
	assert.Equal(t, []int{}, queue.items)
}

func TestQueueDequeueErrorsWhenQueueEmpty(t *testing.T) {
	queue := Queue{}

	item, err := queue.Dequeue()

	assert.Equal(t, -1, item)
	assert.Equal(t, errors.New("queue is empty"), err)
}
