package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListPrepend(t *testing.T) {
	ll := SinglyLinkedList[int]{}

	ll.Prepend(1)

	assert.Equal(t, uint(1), ll.length)
	assert.Equal(t, 1, ll.head.value)
	assert.Nil(t, ll.head.nextNode)

	ll.Prepend(2)

	assert.Equal(t, uint(2), ll.length)
	assert.Equal(t, 2, ll.head.value)
	assert.Equal(t, 1, ll.head.nextNode.value)
}

func TestLinkedListDeleteValue(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(3)
	ll.Prepend(2)
	ll.Prepend(1)
	assert.Equal(t, uint(3), ll.length)

	ll.DeleteValue(2)
	assert.Equal(t, uint(2), ll.length)
	assert.Equal(t, 1, ll.head.value)
	assert.Equal(t, 3, ll.head.nextNode.value)

	ll.DeleteValue(1)
	assert.Equal(t, uint(1), ll.length)
	assert.Equal(t, 3, ll.head.value)

	ll.DeleteValue(3)
	assert.Equal(t, uint(0), ll.length)
}

func TestLinkedListDeleteMultipleOfValue(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(2)
	ll.Prepend(3)
	ll.Prepend(2)
	ll.Prepend(1)
	ll.Prepend(2)

	ll.DeleteValue(2)

	assert.Equal(t, uint(2), ll.length)
	assert.Equal(t, 1, ll.head.value)
	assert.Equal(t, 3, ll.head.nextNode.value)
}

func TestLinkedListDeleteValueDoesNotErrorWhenValueNotFound(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(1)

	ll.DeleteValue(99)
	assert.Equal(t, uint(1), ll.length)

	ll.DeleteValue(1)
	assert.Equal(t, uint(0), ll.length)

	ll.DeleteValue(99)
	assert.Equal(t, uint(0), ll.length)
}
