package graph

import "errors"

// ErrInvalidVertex reports an index outside the graph's dense vertex range.
var ErrInvalidVertex error = errors.New("function requires a valid vertex.")

// Graph is the read-only weighted-graph boundary shared by graph algorithms.
// Concrete adjacency-list and adjacency-matrix structs implement this interface
// directly while retaining ownership of their storage and mutation methods.
type Graph interface {
	// Directed reports whether edges have an orientation.
	Directed() bool
	// VertexCount returns the number of dense vertices indexed from zero.
	VertexCount() int
	// Neighbors visits outgoing edges in deterministic representation order.
	// It returns false when visit requests an early stop.
	Neighbors(vertex int, visit func(neighbor int, weight int64) bool) (bool, error)
}
