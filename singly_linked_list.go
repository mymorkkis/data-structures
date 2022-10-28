package main

import (
	"strconv"
	"strings"
)

type node struct {
	data     int
	nextNode *node
}

// A LinkedList is a structure of nodes linked using pointers.
type LinkedList struct {
	head   *node
	length uint
}

// Prepend sets a new node at the head of the LinkedList.
func (ll *LinkedList) Prepend(node *node) {
	if ll.head != nil {
		node.nextNode = ll.head
	}
	ll.head = node
	ll.length++
}

// DeleteValue will remove any nodes with the given value from the LinkedList.
func (ll *LinkedList) DeleteValue(value int) {
	if ll.length == 0 {
		return
	}

	node := ll.head

	if node.data == value {
		// remove head node
		ll.length--
		ll.head = node.nextNode
	}

	for node.nextNode != nil {
		if node.nextNode.data == value {
			ll.length--
			if node.nextNode.nextNode == nil {
				// remove tail node
				node.nextNode = nil
				break
			}
			// remove a middle node
			node.nextNode = node.nextNode.nextNode
		}
		node = node.nextNode
	}
}

// String will return any values in the LinkedList in a comma seperated string.
func (ll *LinkedList) String() string {
	if ll.head == nil {
		return ""
	}

	var values []string

	node := ll.head
	for node != nil {
		values = append(values, strconv.Itoa(node.data))
		node = node.nextNode
	}

	return strings.Join(values, ", ")
}
