package main

type node[T comparable] struct {
	value    T
	nextNode *node[T]
}

// A SinglyLinkedList is a structure of nodes linked using pointers.
type SinglyLinkedList[T comparable] struct {
	head   *node[T]
	length uint
}

// Prepend sets a new node at the head of the LinkedList.
func (ll *SinglyLinkedList[T]) Prepend(value T) {
	node := &node[T]{value: value}

	if ll.head != nil {
		node.nextNode = ll.head
	}

	ll.head = node
	ll.length++
}

// DeleteValue will remove any nodes with the given value from the LinkedList.
func (ll *SinglyLinkedList[T]) DeleteValue(value T) {
	if ll.length == 0 {
		return
	}

	node := ll.head

	if node.value == value {
		// remove head node
		ll.length--
		ll.head = node.nextNode
	}

	for node.nextNode != nil {
		if node.nextNode.value == value {
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
