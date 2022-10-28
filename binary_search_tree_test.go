package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTreeInsert(t *testing.T) {
	bst := BinarySearchTree[int]{}

	bst.Insert(TreeNode[int]{key: 100})
	assert.Equal(t, 100, bst.root.key)
	assert.Nil(t, bst.root.leftChild)
	assert.Nil(t, bst.root.rightChild)

	bst.Insert(TreeNode[int]{key: 50})
	assert.Equal(t, 100, bst.root.key)
	assert.Equal(t, 50, bst.root.leftChild.key)
	assert.Nil(t, bst.root.rightChild)

	bst.Insert(TreeNode[int]{key: 150})
	assert.Equal(t, 100, bst.root.key)
	assert.Equal(t, 150, bst.root.rightChild.key)
	assert.Equal(t, 50, bst.root.leftChild.key)

	bst.Insert(TreeNode[int]{key: 75})
	rlc := bst.root.leftChild
	assert.Equal(t, 75, rlc.rightChild.key)
	assert.Nil(t, rlc.leftChild)

	bst.Insert(TreeNode[int]{key: 125})
	rrc := bst.root.rightChild
	assert.Equal(t, 125, rrc.leftChild.key)
	assert.Nil(t, rrc.rightChild)
}

func TestBinarySearchTreeSearch(t *testing.T) {
	bst := BinarySearchTree[int]{}
	assert.False(t, bst.Search(999))

	bst.Insert(TreeNode[int]{key: 100})
	bst.Insert(TreeNode[int]{key: 50})
	bst.Insert(TreeNode[int]{key: 150})

	assert.True(t, bst.Search(100))
	assert.True(t, bst.Search(50))
	assert.True(t, bst.Search(150))
	assert.False(t, bst.Search(999))
}
