package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedListPrepend(t *testing.T) {
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

func TestSinglyLinkedListDelete(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(3)
	ll.Prepend(2)
	ll.Prepend(1)
	assert.Equal(t, uint(3), ll.length)

	ll.Delete(2)
	assert.Equal(t, uint(2), ll.length)
	assert.Equal(t, 1, ll.head.value)
	assert.Equal(t, 3, ll.head.nextNode.value)

	ll.Delete(1)
	assert.Equal(t, uint(1), ll.length)
	assert.Equal(t, 3, ll.head.value)

	ll.Delete(3)
	assert.Equal(t, uint(0), ll.length)
}

func TestSinglyLinkedListDeleteMultipleOfValue(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(2)
	ll.Prepend(3)
	ll.Prepend(2)
	ll.Prepend(1)
	ll.Prepend(2)

	ll.Delete(2)

	assert.Equal(t, uint(2), ll.length)
	assert.Equal(t, 1, ll.head.value)
	assert.Equal(t, 3, ll.head.nextNode.value)
}

func TestSinglyLinkedListDeleteDoesNotErrorWhenValueNotFound(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(1)

	ll.Delete(99)
	assert.Equal(t, uint(1), ll.length)

	ll.Delete(1)
	assert.Equal(t, uint(0), ll.length)

	ll.Delete(99)
	assert.Equal(t, uint(0), ll.length)
}

func TestSinglyLinkedListSearch(t *testing.T) {
	ll := SinglyLinkedList[int]{}
	ll.Prepend(1)
	ll.Prepend(2)

	assert.True(t, ll.Search(1))
	assert.True(t, ll.Search(1))
	assert.False(t, ll.Search(99))
}
