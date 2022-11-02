package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTableInsertAndSearchAndDelete(t *testing.T) {
	ht := HashTable{}
	assert.False(t, ht.Search("any"))

	ht.Insert("ray")
	assert.True(t, ht.Search("ray"))

	ht.Insert("yar") // same hash as "ray"
	assert.True(t, ht.Search("yar"))

	ht.Delete("ray")
	assert.False(t, ht.Search("ray"))

	ht.Delete("yar")
	assert.False(t, ht.Search("yar"))
}
