package graph

import "errors"

// ErrInvalidKey reports a key that does not identify a node in the graph.
var ErrInvalidKey error = errors.New("function requires a valid node key.")

// Node is the shared storage record for node-backed data structures.
// Structures use only the links they need and leave the remaining fields nil.
type Node[T any] struct {
	Key   int
	Value *T

	Occurences int

	Next   *Node[T]
	Prev   *Node[T]
	Left   *Node[T]
	Right  *Node[T]
	Parent *Node[T]

	Children map[rune]*Node[T]
	Edges    []Edge[T]
}

// Edge describes one weighted directed connection between stable node keys.
type Edge[T any] struct {
	From   *Node[T]
	To     *Node[T]
	Weight int64
}

// Graph is the read-only weighted-graph boundary shared by graph algorithms.
// Concrete adjacency-list and adjacency-matrix structs implement this interface
// directly while retaining ownership of their storage and mutation methods.
type Graph[T any] interface {
	// Directed reports whether edges have an orientation.
	Directed() bool
	// NodeCount returns the number of nodes currently in the graph.
	NodeCount() int
	// NodeByKey returns the node with key or ErrInvalidKey when it is absent.
	NodeByKey(key int) (*Node[T], bool, error)
	// Neighbors visits outgoing edges in deterministic representation order.
	// It returns false when visit requests an early stop.
	Neighbors(key int, visit func(neighbor *Node[T], weight int64) bool) (bool, error)
}

// UndirectedEdgeGraph extends Graph with each logical undirected edge once.
// Kruskal requires this view so it does not need to reconstruct edges from
// mirrored representation storage.
type UndirectedEdgeGraph[T any] interface {
	Graph[T]
	Edges(visit func(Edge[T]) bool) bool
}
