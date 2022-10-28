package main

import "golang.org/x/exp/constraints"

// A TreeNode is a node in a BinarySearchTree
type TreeNode[T constraints.Ordered] struct {
	key        T
	leftChild  *TreeNode[T]
	rightChild *TreeNode[T]
}

// A BinarySearchTree is an ordered structure of pointer linked nodes
// Nodes with higher value keys are placed to the right of the parent
// Nodes with lower value keys are placed to the left of the parent
type BinarySearchTree[T constraints.Ordered] struct {
	root *TreeNode[T]
}

// Insert inserts the given node as a leaf node
func (bt *BinarySearchTree[T]) Insert(node TreeNode[T]) {
	if bt.root == nil {
		bt.root = &node
		return
	}

	for parentNode := bt.root; ; {
		if node.key < parentNode.key {
			if parentNode.leftChild == nil {
				parentNode.leftChild = &node
				break
			}
			parentNode = parentNode.leftChild
		} else if node.key > parentNode.key {
			if parentNode.rightChild == nil {
				parentNode.rightChild = &node
				break
			}
			parentNode = parentNode.rightChild
		} else {
			// Not allowing duplicate entries
			// Can implement `count` functionality on the node later if needed
			break
		}
	}
}

// Search returns whether the given key is found in the BinarySearchTree
func (bt *BinarySearchTree[T]) Search(key T) bool {
	if bt.root == nil {
		return false
	}

	for node := bt.root; ; {
		if key == node.key {
			return true
		} else if key < node.key {
			if node.leftChild == nil {
				return false
			}
			node = node.leftChild
		} else {
			if node.rightChild == nil {
				return false
			}
			node = node.rightChild
		}
	}
}
