//go:build contract

package dijkstra

import (
	"fmt"
	"testing"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func BenchmarkDijkstra(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		input, sourceKey := dijkstraChainGraph(size)
		b.Run(fmt.Sprintf("weighted-chain/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				result, err := Dijkstra(input, sourceKey)
				if err != nil {
					b.Fatal(err)
				}
				if distance, present := result.Distance(size - 1); !present || distance != int64(size-1) {
					b.Fatal("shortest-path result failed")
				}
			}
		})
	}
}

func dijkstraChainGraph(size int) (fixture, int) {
	nodes := make([]*graph.Node[string], size)
	edges := make(map[int][]edge, size)
	for index := range nodes {
		value := fmt.Sprintf("node-%d", index)
		nodes[index] = &graph.Node[string]{Key: index, Value: &value}
		if index > 0 {
			edges[nodes[index-1].Key] = append(edges[nodes[index-1].Key], edge{to: nodes[index], weight: 1})
		}
	}
	return fixture{nodes: nodes, edges: edges}, nodes[0].Key
}
