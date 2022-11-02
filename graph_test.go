package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanAddNodeToGraph(t *testing.T) {
	g := Graph[int]{}
	assert.Nil(t, g.GetNode(1))

	g.AddNode(1)
	assert.NotNil(t, g.GetNode(1))
}

func TestAddNodeErrorsIfTheNodeAlreadyInTheGraph(t *testing.T) {
	g := Graph[int]{}
	g.AddNode(1)

	err := g.AddNode(1)
	assert.Equal(t, errors.New("duplicate node insertion attempted"), err)
}

func TestCanAddEdgeToGraphq(t *testing.T) {
	g := Graph[int]{}
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(1, 2)

	node := g.GetNode(1)
	assert.Equal(t, 1, len(node.adjacentNodes))
	assert.Equal(t, 2, node.adjacentNodes[0].key)
}

func TestAddEdgeErrorsIfEitherKeyProvidedIsNotInTheGraph(t *testing.T) {
	g := Graph[int]{}

	err := g.AddEdge(1, 2)
	assert.Equal(t, errors.New("key not found in the Graph"), err)

	g.AddNode(1)

	err = g.AddEdge(1, 2)
	assert.Equal(t, errors.New("adjacentKey not found in the Graph"), err)
}

func TestAddEdgeErrorsIfNodeAlreadyHasAdjacentNode(t *testing.T) {
	g := Graph[int]{}
	g.AddNode(1)
	g.AddNode(2)
	g.AddEdge(1, 2)

	err := g.AddEdge(1, 2)
	assert.Equal(t, errors.New("provided adjacentNode already linked to node"), err)
}
