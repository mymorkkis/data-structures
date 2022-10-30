package main

import (
	"errors"
	"strings"
)

type trieNode struct {
	children  [26]*trieNode
	isWordEnd bool
}

// An ASCIITrie is a Trie structure that stores words compiled of the lower case ASCII alphabet.
// It enables quick word look up at the expense of storage space.
type ASCIITrie struct {
	root *trieNode
}

// Insert inserts the given word in the ASCIITrie or returns an error if non ASCII alphabet character found.
func (t *ASCIITrie) Insert(word string) error {
	for _, r := range word {
		if !(r >= 'a' && r <= 'z') && !(r >= 'A' && r <= 'Z') {
			return errors.New("only ASCII alphabet characters accepted")
		}
	}

	if t.root == nil {
		t.root = &trieNode{}
	}

	node := t.root
	for _, r := range strings.ToLower(word) {
		if node.children[r-97] == nil {
			node.children[r-97] = &trieNode{}
		}
		node = node.children[r-97]
	}
	node.isWordEnd = true

	return nil
}

// Search returns whether the word was found in ehe ASCIITrie or not.
func (t *ASCIITrie) Search(word string) bool {
	if t.root == nil {
		return false
	}

	node := t.root
	for _, r := range strings.ToLower(word) {
		if node.children[r-97] == nil {
			return false
		}
		node = node.children[r-97]
	}

	return node.isWordEnd
}
