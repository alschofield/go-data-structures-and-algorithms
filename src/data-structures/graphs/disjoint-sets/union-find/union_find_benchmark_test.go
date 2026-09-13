//go:build contract

package union_find

import (
	"fmt"
	"testing"
)

func BenchmarkUnionFind(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		b.Run(fmt.Sprintf("find-compressed/Size=%d", size), func(b *testing.B) {
			union_find := unionFindChain(b, size)
			if _, ok, err := union_find.Find(size - 1); err != nil || !ok {
				b.Fatal("setup Find() failed")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if _, ok, err := union_find.Find(size - 1); err != nil || !ok {
					b.Fatal("Find() failed")
				}
			}
		})

		b.Run(fmt.Sprintf("connected/Size=%d", size), func(b *testing.B) {
			union_find := unionFindChain(b, size)
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if connected, err := union_find.Connected(0, size-1); err != nil || !connected {
					b.Fatal("Connected() failed")
				}
			}
		})

		b.Run(fmt.Sprintf("union/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				union_find, err := NewUnionFind(size)
				if err != nil {
					b.Fatal(err)
				}
				if merged, err := union_find.Union(size-2, size-1); err != nil || !merged {
					b.Fatal("Union() failed")
				}
			}
		})
	}
}

func unionFindChain(b *testing.B, size int) *UnionFind {
	union_find, err := NewUnionFind(size)
	if err != nil {
		b.Fatal(err)
	}
	for index := 1; index < size; index++ {
		if merged, err := union_find.Union(index-1, index); err != nil || !merged {
			b.Fatal("benchmark setup Union() failed")
		}
	}
	return union_find
}
