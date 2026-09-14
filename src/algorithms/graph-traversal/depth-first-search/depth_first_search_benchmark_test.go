//go:build contract

package depth_first_search

import (
	"fmt"
	"testing"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/representations/adjacency-list"
)

func BenchmarkDepthFirstSearch(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		list, source_key := dfsStarGraph(b, size)

		b.Run(fmt.Sprintf("full-star/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				path, complete, err := DepthFirstSearch(list, source_key, func(*graph.Node[int]) bool { return true })
				if err != nil || !complete || len(path) != size {
					b.Fatal("full DFS traversal failed")
				}
			}
		})

		b.Run(fmt.Sprintf("early-stop/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				path, complete, err := DepthFirstSearch(list, source_key, func(node *graph.Node[int]) bool { return node.Key != size-1 })
				if err != nil || !complete || len(path) != 2 {
					b.Fatal("early-stop DFS traversal failed")
				}
			}
		})
	}
}

func dfsStarGraph(b *testing.B, size int) (*adjacency_list.AdjacencyList[int], int) {
	list := adjacency_list.NewAdjacencyList[int](true)
	nodes := make([]*graph.Node[int], size)
	for index := range nodes {
		value := index
		nodes[index] = list.AddVertex(&value)
	}
	for index := 1; index < size; index++ {
		if added, err := list.AddEdge(nodes[0].Key, nodes[index].Key, 1); err != nil || !added {
			b.Fatal("benchmark graph setup failed")
		}
	}
	return list, nodes[0].Key
}
