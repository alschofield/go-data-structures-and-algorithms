//go:build contract

package adjacency_matrix

import (
	"fmt"
	"testing"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func BenchmarkAdjacencyMatrix(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		b.Run(fmt.Sprintf("add-edge/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				matrix, nodes := adjacencyMatrixNodes(b, size)
				if added, err := matrix.AddEdge(nodes[size-2].Key, nodes[size-1].Key, 1); err != nil || !added {
					b.Fatal("AddEdge() failed")
				}
			}
		})

		b.Run(fmt.Sprintf("has-edge/Size=%d", size), func(b *testing.B) {
			matrix, nodes := adjacencyMatrixNodes(b, size)
			if added, err := matrix.AddEdge(nodes[0].Key, nodes[size-1].Key, 1); err != nil || !added {
				b.Fatal("benchmark setup AddEdge() failed")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if found, err := matrix.HasEdge(nodes[0].Key, nodes[size-1].Key); err != nil || !found {
					b.Fatal("HasEdge() missed an inserted edge")
				}
			}
		})

		b.Run(fmt.Sprintf("neighbors/Size=%d", size), func(b *testing.B) {
			matrix, nodes := adjacencyMatrixNodes(b, size)
			for index := 1; index < size; index++ {
				if added, err := matrix.AddEdge(nodes[0].Key, nodes[index].Key, 1); err != nil || !added {
					b.Fatal("benchmark setup AddEdge() failed")
				}
			}
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				visits := 0
				if complete, err := matrix.Neighbors(nodes[0].Key, func(*graph.Node[int], int64) bool { visits++; return true }); err != nil || !complete || visits != size-1 {
					b.Fatal("Neighbors() did not visit the complete matrix row")
				}
			}
		})
	}
}

func adjacencyMatrixNodes(b *testing.B, size int) (*AdjacencyMatrix[int], []*graph.Node[int]) {
	matrix := NewAdjacencyMatrix[int](true)
	nodes := make([]*graph.Node[int], size)
	for index := range nodes {
		value := index
		nodes[index] = matrix.AddVertex(&value)
	}
	return matrix, nodes
}
