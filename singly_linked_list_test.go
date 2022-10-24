package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListPrepend(t *testing.T) {
	ll := LinkedList{}
	assert.Equal(t, "", ll.String())

	ll.Prepend(&node{data: 1})

	assert.Equal(t, uint(1), ll.length)
	assert.Equal(t, 1, ll.head.data)
	assert.Nil(t, ll.head.nextNode)
	assert.Equal(t, "1", ll.String())

	ll.Prepend(&node{data: 2})

	assert.Equal(t, uint(2), ll.length)
	assert.Equal(t, 2, ll.head.data)
	assert.Equal(t, &node{data: 1}, ll.head.nextNode)
	assert.Equal(t, "2, 1", ll.String())
}

func TestLinkedListDeleteValue(t *testing.T) {
	ll := LinkedList{}
	ll.Prepend(&node{data: 3})
	ll.Prepend(&node{data: 2})
	ll.Prepend(&node{data: 1})
	assert.Equal(t, "1, 2, 3", ll.String())
	assert.Equal(t, uint(3), ll.length)

	ll.DeleteValue(2)
	assert.Equal(t, "1, 3", ll.String())
	assert.Equal(t, uint(2), ll.length)

	ll.DeleteValue(1)
	assert.Equal(t, "3", ll.String())
	assert.Equal(t, uint(1), ll.length)

	ll.DeleteValue(3)
	assert.Equal(t, "", ll.String())
	assert.Equal(t, uint(0), ll.length)
}

func TestLinkedListDeleteMultipleOfValue(t *testing.T) {
	ll := LinkedList{}
	ll.Prepend(&node{data: 2})
	ll.Prepend(&node{data: 3})
	ll.Prepend(&node{data: 2})
	ll.Prepend(&node{data: 1})
	ll.Prepend(&node{data: 2})
	assert.Equal(t, "2, 1, 2, 3, 2", ll.String())

	ll.DeleteValue(2)

	assert.Equal(t, "1, 3", ll.String())
	assert.Equal(t, uint(2), ll.length)
}

func TestLinkedListDeleteValueDoesNotErrorWhenValueNotFound(t *testing.T) {
	ll := LinkedList{}
	ll.Prepend(&node{data: 1})

	ll.DeleteValue(99)
	assert.Equal(t, "1", ll.String())
	assert.Equal(t, uint(1), ll.length)

	ll.DeleteValue(1)
	assert.Equal(t, uint(0), ll.length)

	ll.DeleteValue(99)
	assert.Equal(t, "", ll.String())
	assert.Equal(t, uint(0), ll.length)
}
