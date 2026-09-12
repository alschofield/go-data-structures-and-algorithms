//go:build contract

package binary_search_tree

import (
	"fmt"
	"testing"
)

func BenchmarkBinarySearchTree(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		b.Run(fmt.Sprintf("insert-random/Size=%d", size), func(b *testing.B) {
			values := bstRandomValues(size)
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				tree, err := NewBinarySearchTree(compareInts)
				if err != nil {
					b.Fatal(err)
				}
				for _, value := range values {
					if _, added := tree.Insert(value); !added {
						b.Fatal("random insert unexpectedly found a duplicate")
					}
				}
			}
		})

		b.Run(fmt.Sprintf("insert-sorted/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				tree, err := NewBinarySearchTree(compareInts)
				if err != nil {
					b.Fatal(err)
				}
				for value := 0; value < size; value++ {
					if _, added := tree.Insert(value); !added {
						b.Fatal("sorted insert unexpectedly found a duplicate")
					}
				}
			}
		})

		b.Run(fmt.Sprintf("find/Size=%d", size), func(b *testing.B) {
			tree := bstTree(b, bstRandomValues(size))
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if tree.Find(iteration%size) == nil {
					b.Fatal("Find() missed an inserted value")
				}
			}
		})

		b.Run(fmt.Sprintf("remove/Size=%d", size), func(b *testing.B) {
			values := bstRandomValues(size)
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				tree := bstTree(b, values)
				if _, removed := tree.Remove(values[iteration%size]); !removed {
					b.Fatal("Remove() missed an inserted value")
				}
			}
		})
	}
}

func compareInts(left, right int) int { return left - right }

func bstRandomValues(size int) []int {
	values := make([]int, size)
	for index := range values {
		values[index] = (index*7919 + 104729) % size
	}
	return values
}

func bstTree(b *testing.B, values []int) *BinarySearchTree[int] {
	tree, err := NewBinarySearchTree(compareInts)
	if err != nil {
		b.Fatal(err)
	}
	for _, value := range values {
		if _, added := tree.Insert(value); !added {
			b.Fatal("tree setup unexpectedly found a duplicate")
		}
	}
	return tree
}
