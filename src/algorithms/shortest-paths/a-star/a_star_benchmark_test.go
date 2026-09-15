//go:build contract

package a_star

import (
	"fmt"
	"testing"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

func BenchmarkAStar(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		input, sourceKey, goalKey := aStarChainGraph(size)
		for _, workload := range []struct {
			name      string
			heuristic Heuristic
		}{
			{"zero-heuristic", func(int) int64 { return 0 }},
			{"exact-chain-heuristic", func(key int) int64 { return int64(goalKey - key) }},
		} {
			b.Run(fmt.Sprintf("%s/Size=%d", workload.name, size), func(b *testing.B) {
				b.ReportAllocs()
				for iteration := 0; iteration < b.N; iteration++ {
					path, err := AStar(input, sourceKey, goalKey, workload.heuristic)
					if err != nil || len(path) != size {
						b.Fatal("A* pathfinding failed")
					}
				}
			})
		}
	}
}

func aStarChainGraph(size int) (fixture, int, int) {
	nodes := make([]*graph.Node[string], size)
	edges := make(map[int][]edge, size)
	for index := range nodes {
		value := fmt.Sprintf("node-%d", index)
		nodes[index] = &graph.Node[string]{Key: index, Value: &value}
		if index > 0 {
			edges[nodes[index-1].Key] = append(edges[nodes[index-1].Key], edge{to: nodes[index], weight: 1})
		}
	}
	return fixture{nodes: nodes, edges: edges}, nodes[0].Key, nodes[len(nodes)-1].Key
}
