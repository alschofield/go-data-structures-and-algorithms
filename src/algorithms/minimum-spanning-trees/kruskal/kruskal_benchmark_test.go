//go:build contract

package kruskal

import (
	"fmt"
	"testing"
)

func BenchmarkKruskal(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		connected := kruskalCycleGraph(size)
		b.Run(fmt.Sprintf("connected-cycle/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				forest, _, err := Kruskal(connected)
				if err != nil || len(forest) != size-1 {
					b.Fatal("Kruskal spanning tree failed")
				}
			}
		})

		forest := kruskalForestGraph(size)
		b.Run(fmt.Sprintf("disconnected-forest/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				selected, _, err := Kruskal(forest)
				if err != nil || len(selected) != size/2 {
					b.Fatal("Kruskal spanning forest failed")
				}
			}
		})
	}
}

// kruskalCycleGraph builds one connected cycle with varied weights so the sort
// order differs from the discovery order and exactly one edge must be skipped.
func kruskalCycleGraph(size int) graphFixture {
	values := make([]string, size)
	for index := range values {
		values[index] = fmt.Sprintf("node-%d", index)
	}
	edges := make([]edge, 0, size)
	for index := 0; index < size-1; index++ {
		edges = append(edges, edge{from: index, to: index + 1, weight: int64((index*7)%13 + 1)})
	}
	edges = append(edges, edge{from: size - 1, to: 0, weight: 100})
	return undirectedFixture(values, edges)
}

// kruskalForestGraph builds size/2 disjoint pairs so the result is a forest.
func kruskalForestGraph(size int) graphFixture {
	values := make([]string, size)
	for index := range values {
		values[index] = fmt.Sprintf("node-%d", index)
	}
	edges := make([]edge, 0, size/2)
	for index := 0; index+1 < size; index += 2 {
		edges = append(edges, edge{from: index, to: index + 1, weight: int64(index%9 + 1)})
	}
	return undirectedFixture(values, edges)
}
