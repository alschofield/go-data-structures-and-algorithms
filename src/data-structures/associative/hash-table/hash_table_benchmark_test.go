//go:build contract

package hash_table

import (
	"strconv"
	"testing"
)

var hashTableBenchmarkBool bool
var hashTableBenchmarkValue int

func BenchmarkHashTableSet(b *testing.B) {
	for _, test := range []struct {
		name string
		hash func(int) uint
	}{
		{name: "uniform", hash: intHash},
		{name: "colliding", hash: func(int) uint { return 0 }},
	} {
		b.Run(test.name, func(b *testing.B) {
			table := benchmarkTable(b, b.N+1, test.hash)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, hashTableBenchmarkBool = table.Set(i, i)
			}
			b.StopTimer()
			if table.Len() != b.N || hashTableBenchmarkBool {
				b.Fatalf("Set() left Len() = %d and replaced = %t", table.Len(), hashTableBenchmarkBool)
			}
		})
	}
}

func BenchmarkHashTableGet(b *testing.B) {
	for _, test := range []struct {
		name string
		hash func(int) uint
	}{
		{name: "uniform", hash: intHash},
		{name: "colliding", hash: func(int) uint { return 0 }},
	} {
		for _, size := range []int{1_024, 65_536} {
			b.Run(test.name+"/Size="+strconv.Itoa(size), func(b *testing.B) {
				table := benchmarkTableOfSize(b, size, test.hash)
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					value, ok := table.Get(i % size)
					if !ok {
						b.Fatal("Get() = absent, want present key")
					}
					hashTableBenchmarkValue = value
				}
				b.StopTimer()
				if table.Len() != size || hashTableBenchmarkValue < 0 {
					b.Fatalf("Get() left Len() = %d and value = %d", table.Len(), hashTableBenchmarkValue)
				}
			})
		}
	}
}

func BenchmarkHashTableRemove(b *testing.B) {
	for _, test := range []struct {
		name string
		hash func(int) uint
	}{
		{name: "uniform", hash: intHash},
		{name: "colliding", hash: func(int) uint { return 0 }},
	} {
		b.Run(test.name, func(b *testing.B) {
			table := benchmarkTableOfSize(b, b.N, test.hash)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				value, ok := table.Remove(i)
				if !ok {
					b.Fatal("Remove() = absent, want present key")
				}
				hashTableBenchmarkValue = value
			}
			b.StopTimer()
			if table.Len() != 0 || hashTableBenchmarkValue != b.N-1 {
				b.Fatalf("Remove() left Len() = %d and value = %d", table.Len(), hashTableBenchmarkValue)
			}
		})
	}
}

func BenchmarkHashTableInsertionModes(b *testing.B) {
	for _, test := range []struct {
		name string
		set  func(*HashTable[int, int], int, int) (int, bool)
	}{
		{name: "fixed", set: (*HashTable[int, int]).Set},
		{name: "resizing", set: (*HashTable[int, int]).SetResizing},
	} {
		b.Run(test.name, func(b *testing.B) {
			table := benchmarkTable(b, 1, intHash)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, hashTableBenchmarkBool = test.set(table, i, i)
			}
			b.StopTimer()
			if table.Len() != b.N || hashTableBenchmarkBool {
				b.Fatalf("%s insert left Len() = %d and replaced = %t", test.name, table.Len(), hashTableBenchmarkBool)
			}
		})
	}
}

func BenchmarkHashTableRehash(b *testing.B) {
	for _, size := range []int{1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				table := benchmarkTable(b, size, intHash)
				for key := 0; key < size*3/4; key++ {
					table.Set(key, key)
				}
				b.StartTimer()
				_, hashTableBenchmarkBool = table.SetResizing(size*2+i, i)
				b.StopTimer()
				if table.Cap() != size*2 {
					b.Fatalf("SetResizing() Cap() = %d, want %d", table.Cap(), size*2)
				}
			}
			if hashTableBenchmarkBool {
				b.Fatal("rehash insert replaced an existing value")
			}
		})
	}
}

func benchmarkTable(b *testing.B, capacity int, hash func(int) uint) *HashTable[int, int] {
	b.Helper()
	table, err := NewHashTable[int, int](capacity, hash, func(a, b int) bool { return a == b })
	if err != nil {
		b.Fatalf("NewHashTable() error = %v", err)
	}
	return table
}

func benchmarkTableOfSize(b *testing.B, size int, hash func(int) uint) *HashTable[int, int] {
	table := benchmarkTable(b, size+1, hash)
	for i := 0; i < size; i++ {
		table.Set(i, i)
	}
	return table
}

func intHash(value int) uint {
	return uint(value)
}
