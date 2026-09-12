package adjacency_list

import "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"

// AdjacencyList stores each node's outgoing weighted edges on that node.
type AdjacencyList[T any] struct {
	node_count int
	edge_count int
	next_key   int
	directed   bool
	nodes      []*graph.Node[T]
}

// NewAdjacencyList creates an empty directed or undirected graph.
func NewAdjacencyList[T any](directed bool) *AdjacencyList[T] {
	return &AdjacencyList[T]{
		node_count: 0,
		edge_count: 0,
		next_key:   0,
		directed:   directed,
		nodes:      []*graph.Node[T]{},
	}
}

// AddVertex appends one node and assigns its never-reused graph key.
func (al *AdjacencyList[T]) AddVertex(value *T) *graph.Node[T] {
	node := &graph.Node[T]{
		Key:   al.next_key,
		Value: value,
		Edges: []graph.Edge[T]{},
	}

	al.nodes = append(al.nodes, node)
	al.node_count++
	al.next_key++

	return node
}

// NodeByKey uses the preserved node-key slot for constant-time lookup.
func (al *AdjacencyList[T]) NodeByKey(key int) (*graph.Node[T], bool) {
	if key < 0 || key >= al.node_count {
		return nil, false
	}

	return al.nodes[key], al.nodes[key] != nil
}

// AddEdge adds one logical edge and mirrors it when the graph is undirected.
func (al *AdjacencyList[T]) AddEdge(from_key int, to_key int, weight int64) (bool, error) {
	from_node, status := al.NodeByKey(from_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	// Scan the source adjacency list so duplicate logical edges remain absent.
	has_edge, err := al.HasEdge(from_key, to_key)
	if err != nil {
		return false, err
	}

	if has_edge {
		return false, nil
	}

	to_node, status := al.NodeByKey(to_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	edge := &graph.Edge[T]{
		From:   from_node,
		To:     to_node,
		Weight: weight,
	}

	from_node.Edges = append(from_node.Edges, *edge)

	if !al.directed && from_key != to_key {
		// Undirected storage has a reciprocal arc but still one logical edge count.
		edge = &graph.Edge[T]{
			From:   to_node,
			To:     from_node,
			Weight: weight,
		}

		to_node.Edges = append(to_node.Edges, *edge)
	}

	al.edge_count++

	return true, nil
}

// RemoveEdge removes one logical edge and its reciprocal arc when needed.
func (al *AdjacencyList[T]) RemoveEdge(from_key int, to_key int) (bool, error) {
	from_node, status := al.NodeByKey(from_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	to_node, status := al.NodeByKey(to_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	removed := false
	for i := 0; i < len(from_node.Edges); i++ {
		if from_node.Edges[i].To == to_node {
			// Preserve the remaining edge order while removing the matching arc.
			from_node.Edges = append(from_node.Edges[:i], from_node.Edges[i+1:]...)
			removed = true
			break
		}
	}

	if !al.directed && from_key != to_key {
		for i := 0; i < len(to_node.Edges); i++ {
			if to_node.Edges[i].To == from_node {
				to_node.Edges = append(to_node.Edges[:i], to_node.Edges[i+1:]...)
				removed = true
				break
			}
		}
	}

	if removed {
		al.edge_count--
		return true, nil
	} else {
		return false, nil
	}
}

// HasEdge scans only the source node's outgoing adjacency list.
func (al *AdjacencyList[T]) HasEdge(from_key int, to_key int) (bool, error) {
	from_node, status := al.NodeByKey(from_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	to_node, status := al.NodeByKey(to_key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	for i := 0; i < len(from_node.Edges); i++ {
		if from_node.Edges[i].To == to_node {
			return true, nil
		}
	}

	return false, nil
}

// Neighbors visits outgoing arcs until all are visited or visit requests a stop.
func (al *AdjacencyList[T]) Neighbors(key int, visit func(*graph.Node[T], int64) bool) (bool, error) {
	node, status := al.NodeByKey(key)

	if !status {
		return false, graph.ErrInvalidKey
	}

	for i := 0; i < len(node.Edges); i++ {
		if !visit(node.Edges[i].To, node.Edges[i].Weight) {
			return false, nil
		}
	}

	return true, nil
}

// NodeCount returns the number of stored nodes.
func (al *AdjacencyList[T]) NodeCount() int {
	return al.node_count
}

// EdgeCount returns the number of logical edges, not stored reciprocal arcs.
func (al *AdjacencyList[T]) EdgeCount() int {
	return al.edge_count
}

// Directed reports whether AddEdge stores only its requested orientation.
func (al *AdjacencyList[T]) Directed() bool {
	return al.directed
}

// Edges exposes each logical undirected edge once for algorithms such as Kruskal.
func (al *AdjacencyList[T]) Edges(visit func(graph.Edge[T]) bool) (bool, error) {
	if al.Directed() {
		return false, graph.ErrDirectedGraph
	}

	for i := 0; i < al.node_count; i++ {
		for n := 0; n < len(al.nodes[i].Edges); n++ {
			// The lower-key orientation is the canonical representative of a pair.
			if al.nodes[i].Edges[n].From.Key <= al.nodes[i].Edges[n].To.Key {
				if !visit(al.nodes[i].Edges[n]) {
					return false, nil
				}
			}
		}
	}

	return true, nil
}
