package main

import "errors"

type Node[T comparable] struct {
	key           T
	adjacentNodes []*Node[T]
}

func (n Node[T]) GetAdjacentNode(key T) *Node[T] {
	for _, an := range n.adjacentNodes {
		if an.key == key {
			return an
		}
	}
	return nil
}

// A Graph is a data structure using adjecency lists to link Nodes
type Graph[T comparable] struct {
	nodes []*Node[T]
}

// AddNode adds a Node to the Graph or errors if it is a duplicate
func (g *Graph[T]) AddNode(key T) error {
	if g.GetNode(key) != nil {
		return errors.New("duplicate node insertion attempted")
	}

	g.nodes = append(g.nodes, &Node[T]{key: key})

	return nil
}

// AddEdge sets a Node as an adjacent Node on another Node
func (g *Graph[T]) AddEdge(key, adjacentKey T) error {
	node := g.GetNode(key)
	if node == nil {
		return errors.New("key not found in the Graph")
	}
	if node.GetAdjacentNode(adjacentKey) != nil {
		return errors.New("provided adjacentNode already linked to node")
	}

	adjacentNode := g.GetNode(adjacentKey)
	if adjacentNode == nil {
		return errors.New("adjacentKey not found in the Graph")
	}

	node.adjacentNodes = append(node.adjacentNodes, adjacentNode)

	return nil
}

// GetNode retrieves a Node from the Graph
func (g *Graph[T]) GetNode(key T) *Node[T] {
	for _, n := range g.nodes {
		if n.key == key {
			return n
		}
	}
	return nil
}
