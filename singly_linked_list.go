package main

import (
	"fmt"
)

type node struct {
	data     int
	nextNode *node
}

type linkedList struct {
	head   *node
	length uint
}

func (ll *linkedList) prepend(node *node) {
	if ll.head != nil {
		node.nextNode = ll.head
	}
	ll.head = node
	ll.length++
}

func (ll *linkedList) deleteValue(value int) {
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

func (ll linkedList) displayValues() string {
	values := ""

	if ll.head == nil {
		return values
	}

	node := ll.head
	for node != nil {
		values += fmt.Sprintf("%d", node.data)
		node = node.nextNode
		if node != nil {
			values += fmt.Sprint(", ")
		}
	}

	return values
}
