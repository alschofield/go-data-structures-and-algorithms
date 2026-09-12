//go:build contract

package adjacency_list

import (
	"fmt"
	"testing"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func BenchmarkAdjacencyList(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		b.Run(fmt.Sprintf("add-edge/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				list, nodes := adjacencyListChain(b, size)
				if added, err := list.AddEdge(nodes[size-2].Key, nodes[size-1].Key, 1); err != nil || !added {
					b.Fatal("AddEdge() failed")
				}
			}
		})

		b.Run(fmt.Sprintf("has-edge/Size=%d", size), func(b *testing.B) {
			list, nodes := adjacencyListChain(b, size)
			if added, err := list.AddEdge(nodes[0].Key, nodes[size-1].Key, 1); err != nil || !added {
				b.Fatal("benchmark setup AddEdge() failed")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if found, err := list.HasEdge(nodes[0].Key, nodes[size-1].Key); err != nil || !found {
					b.Fatal("HasEdge() missed an inserted edge")
				}
			}
		})

		b.Run(fmt.Sprintf("neighbors/Size=%d", size), func(b *testing.B) {
			list, nodes := adjacencyListChain(b, size)
			for index := 1; index < size; index++ {
				if added, err := list.AddEdge(nodes[0].Key, nodes[index].Key, 1); err != nil || !added {
					b.Fatal("benchmark setup AddEdge() failed")
				}
			}
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				visits := 0
				if complete, err := list.Neighbors(nodes[0].Key, func(*graph.Node[int], int64) bool { visits++; return true }); err != nil || !complete || visits != size-1 {
					b.Fatal("Neighbors() did not visit the complete adjacency list")
				}
			}
		})
	}
}

func adjacencyListChain(b *testing.B, size int) (*AdjacencyList[int], []*graph.Node[int]) {
	list := NewAdjacencyList[int](true)
	nodes := make([]*graph.Node[int], size)
	for index := range nodes {
		value := index
		nodes[index] = list.AddVertex(&value)
	}
	return list, nodes
}
