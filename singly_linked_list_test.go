package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListPrepend(t *testing.T) {
	ll := linkedList{}
	assert.Equal(t, "", ll.displayValues())

	ll.prepend(&node{data: 1})

	assert.Equal(t, uint(1), ll.length)
	assert.Equal(t, 1, ll.head.data)
	assert.Nil(t, ll.head.nextNode)
	assert.Equal(t, "1", ll.displayValues())

	ll.prepend(&node{data: 2})

	assert.Equal(t, uint(2), ll.length)
	assert.Equal(t, 2, ll.head.data)
	assert.Equal(t, &node{data: 1}, ll.head.nextNode)
	assert.Equal(t, "2, 1", ll.displayValues())
}

func TestLinkedListDeleteValue(t *testing.T) {
	ll := linkedList{}
	ll.prepend(&node{data: 3})
	ll.prepend(&node{data: 2})
	ll.prepend(&node{data: 1})
	assert.Equal(t, "1, 2, 3", ll.displayValues())
	assert.Equal(t, uint(3), ll.length)

	ll.deleteValue(2)
	assert.Equal(t, "1, 3", ll.displayValues())
	assert.Equal(t, uint(2), ll.length)

	ll.deleteValue(1)
	assert.Equal(t, "3", ll.displayValues())
	assert.Equal(t, uint(1), ll.length)

	ll.deleteValue(3)
	assert.Equal(t, "", ll.displayValues())
	assert.Equal(t, uint(0), ll.length)
}

func TestLinkedListDeleteMultipleOfValue(t *testing.T) {
	ll := linkedList{}
	ll.prepend(&node{data: 2})
	ll.prepend(&node{data: 3})
	ll.prepend(&node{data: 2})
	ll.prepend(&node{data: 1})
	ll.prepend(&node{data: 2})
	assert.Equal(t, "2, 1, 2, 3, 2", ll.displayValues())

	ll.deleteValue(2)

	assert.Equal(t, "1, 3", ll.displayValues())
	assert.Equal(t, uint(2), ll.length)
}

func TestLinkedListDeleteNonExistentValueDoesNotError(t *testing.T) {
	ll := linkedList{}
	ll.prepend(&node{data: 1})

	ll.deleteValue(99)
	assert.Equal(t, "1", ll.displayValues())
	assert.Equal(t, uint(1), ll.length)

	ll.deleteValue(1)
	assert.Equal(t, uint(0), ll.length)

	ll.deleteValue(99)
	assert.Equal(t, "", ll.displayValues())
	assert.Equal(t, uint(0), ll.length)
}
