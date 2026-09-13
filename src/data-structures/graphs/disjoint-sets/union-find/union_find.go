package union_find

import (
	"errors"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

// ErrInvalidCapacity reports a negative element count at construction.
var ErrInvalidCapacity error = errors.New("function requires a valid capacity.")

// ErrInvalidIndex reports an element outside the Union-Find node slice.
var ErrInvalidIndex error = errors.New("function requires a valid index.")

// UnionFind groups shared nodes into disjoint parent-pointer trees.
type UnionFind struct {
	size  int
	sets  int
	nodes []*graph.Node[struct{}]
}

// NewUnionFind initializes count singleton sets with self-parented nodes.
func NewUnionFind(count int) (*UnionFind, error) {
	if count < 0 {
		return nil, ErrInvalidCapacity
	}

	nodes := make([]*graph.Node[struct{}], count)

	for i := 0; i < len(nodes); i++ {
		// Each root starts at rank zero and points to itself.
		nodes[i] = &graph.Node[struct{}]{
			Key:   i,
			Value: &struct{}{},
			Rank:  0,
		}

		nodes[i].Parent = nodes[i]
	}

	union := &UnionFind{
		size:  count,
		sets:  count,
		nodes: nodes,
	}

	return union, nil
}

// Find returns a set representative and applies path halving on the way up.
func (uf *UnionFind) Find(index int) (int, bool, error) {
	if index < 0 || index >= uf.size {
		return 0, false, ErrInvalidIndex
	}

	candidate := uf.nodes[index]
	for candidate.Parent != candidate {
		// Point the current node at its grandparent before advancing upward.
		prev := candidate
		candidate = candidate.Parent
		prev.Parent = candidate.Parent
	}

	return candidate.Key, true, nil
}

// Union joins two distinct roots by rank and preserves shallow trees.
func (uf *UnionFind) Union(thing_one int, thing_two int) (bool, error) {
	thing_one_root, status, err := uf.Find(thing_one)
	if err != nil {
		return false, err
	}

	if !status {
		return false, ErrInvalidIndex
	}

	thing_two_root, status, err := uf.Find(thing_two)
	if err != nil {
		return false, err
	}

	if !status {
		return false, ErrInvalidIndex
	}

	if uf.nodes[thing_one_root] == uf.nodes[thing_two_root] {
		return false, nil
	}

	if uf.nodes[thing_one_root].Rank >= uf.nodes[thing_two_root].Rank {
		uf.nodes[thing_two_root].Parent = uf.nodes[thing_one_root]

		// Equal-rank trees gain one possible level after either root wins.
		if uf.nodes[thing_one_root].Rank == uf.nodes[thing_two_root].Rank {
			uf.nodes[thing_one_root].Rank++
		}
	} else {
		uf.nodes[thing_one_root].Parent = uf.nodes[thing_two_root]
	}

	uf.sets--
	return true, nil
}

// Connected compares the representatives returned by Find.
func (uf *UnionFind) Connected(thing_one int, thing_two int) (bool, error) {
	thing_one_root, status, err := uf.Find(thing_one)
	if err != nil {
		return false, err
	}

	if !status {
		return false, ErrInvalidIndex
	}

	thing_two_root, status, err := uf.Find(thing_two)
	if err != nil {
		return false, err
	}

	if !status {
		return false, ErrInvalidIndex
	}

	return uf.nodes[thing_one_root].Key == uf.nodes[thing_two_root].Key, nil
}

// SetCount returns the number of remaining disjoint sets.
func (uf *UnionFind) SetCount() int {
	return uf.sets
}
