package adjacency_matrix

import "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"

// AdjacencyMatrix stores edge records in an N by N row matrix.
type AdjacencyMatrix[T any] struct {
	directed   bool
	node_count int
	edge_count int
	next_key   int
	nodes      []*graph.Node[T]
	edges      [][]*graph.Edge[T]
}

// NewAdjacencyMatrix creates an empty directed or undirected square matrix.
func NewAdjacencyMatrix[T any](directed bool) *AdjacencyMatrix[T] {
	return &AdjacencyMatrix[T]{
		directed: directed,
		next_key: 0,
		nodes:    []*graph.Node[T]{},
		edges:    [][]*graph.Edge[T]{},
	}
}

// AddVertex grows every matrix row and appends one new empty row.
func (am *AdjacencyMatrix[T]) AddVertex(value *T) *graph.Node[T] {
	node := &graph.Node[T]{
		Key:   am.next_key,
		Value: value,
	}

	am.nodes = append(am.nodes, node)

	// Preserve the square invariant by adding one empty column to every old row.
	for i := 0; i < len(am.edges); i++ {
		am.edges[i] = append(am.edges[i], nil)
	}

	am.edges = append(am.edges, make([]*graph.Edge[T], len(am.nodes)))

	am.node_count++
	am.next_key++

	return node
}

// NodeByKey resolves the direct matrix row and column key.
func (am *AdjacencyMatrix[T]) NodeByKey(key int) (*graph.Node[T], bool) {
	if key < 0 || key >= len(am.nodes) {
		return nil, false
	}

	return am.nodes[key], am.nodes[key] != nil
}

// AddEdge writes one matrix cell and its reciprocal cell when undirected.
func (am *AdjacencyMatrix[T]) AddEdge(from_key int, to_key int, weight int64) (bool, error) {
	from_node, status := am.NodeByKey(from_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	has_edge, err := am.HasEdge(from_key, to_key)

	if err != nil {
		return false, err
	}

	if has_edge {
		return false, nil
	}

	to_node, status := am.NodeByKey(to_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	am.edges[from_key][to_key] = &graph.Edge[T]{
		From:   from_node,
		To:     to_node,
		Weight: weight,
	}

	if !am.Directed() {
		// The reciprocal record lets each row expose its own outgoing neighbor.
		am.edges[to_key][from_key] = &graph.Edge[T]{
			From:   to_node,
			To:     from_node,
			Weight: weight,
		}
	}

	am.edge_count++

	return true, nil
}

// RemoveEdge clears one logical edge and its reciprocal cell when needed.
func (am *AdjacencyMatrix[T]) RemoveEdge(from_key int, to_key int) (bool, error) {
	_, status := am.NodeByKey(from_key)
	if !status {
		return false, graph.ErrInvalidKey
	}

	_, status = am.NodeByKey(to_key)
	if !status {
		return false, graph.ErrInvalidKey
	}

	has_edge, err := am.HasEdge(from_key, to_key)

	if err != nil {
		return false, err
	}

	if !has_edge {
		return false, nil
	}

	am.edges[from_key][to_key] = nil

	if !am.Directed() {
		am.edges[to_key][from_key] = nil
	}

	am.edge_count--

	return true, nil
}

// HasEdge checks one matrix cell after validating both node keys.
func (am *AdjacencyMatrix[T]) HasEdge(from_key int, to_key int) (bool, error) {
	_, status := am.NodeByKey(from_key)
	if !status {
		return false, graph.ErrInvalidKey
	}

	_, status = am.NodeByKey(to_key)
	if !status {
		return false, graph.ErrInvalidKey
	}

	return am.edges[from_key][to_key] != nil, nil
}

// Neighbors scans one full matrix row in ascending key order.
func (am *AdjacencyMatrix[T]) Neighbors(key int, visit func(*graph.Node[T], int64) bool) (bool, error) {
	_, status := am.NodeByKey(key)
	if !status {
		return false, graph.ErrInvalidKey
	}

	for i := 0; i < len(am.edges[key]); i++ {
		if am.edges[key][i] != nil {
			if !visit(am.edges[key][i].To, am.edges[key][i].Weight) {
				return false, nil
			}
		}
	}

	return true, nil
}

// NodeCount returns the current matrix width.
func (am *AdjacencyMatrix[T]) NodeCount() int {
	return am.node_count
}

// EdgeCount returns logical edges rather than reciprocal stored records.
func (am *AdjacencyMatrix[T]) EdgeCount() int {
	return am.edge_count
}

// Directed reports whether edges have orientation.
func (am *AdjacencyMatrix[T]) Directed() bool {
	return am.directed
}

// Edges exposes each logical undirected edge once in row-major key order.
func (am *AdjacencyMatrix[T]) Edges(visit func(graph.Edge[T]) bool) (bool, error) {
	if am.Directed() {
		return false, graph.ErrDirectedGraph
	}

	for i := 0; i < am.node_count; i++ {
		for n := 0; n < am.node_count; n++ {
			// The lower-key direction canonically represents mirrored edge storage.
			if am.edges[i][n] != nil {
				if am.edges[i][n].From.Key <= am.edges[i][n].To.Key {
					if !visit(*am.edges[i][n]) {
						return false, nil
					}
				}
			}
		}
	}

	return true, nil
}
