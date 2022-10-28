package main

type TreeNode struct {
	key        int
	leftChild  *TreeNode
	rightChild *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func (bt *BinarySearchTree) Insert(node TreeNode) {
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

func (bt *BinarySearchTree) Search(key int) bool {
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
