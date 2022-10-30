package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieInsertAndSearch(t *testing.T) {
	trie := ASCIITrie{}
	assert.False(t, trie.Search("any"))

	err := trie.Insert("oreo")
	assert.Nil(t, err)
	err = trie.Insert("OrEGano")
	assert.Nil(t, err)

	assert.True(t, trie.Search("OReo"))
	assert.True(t, trie.Search("oregano"))
	assert.False(t, trie.Search("or"))

	err = trie.Insert("or")
	assert.Nil(t, err)

	assert.True(t, trie.Search("or"))
}

func TestTrieInsertErrorsWhenNonASCIIAlphabetUsed(t *testing.T) {
	trie := ASCIITrie{}

	err := trie.Insert("héllò")
	assert.Equal(t, errors.New("only ASCII alphabet characters accepted"), err)
}
