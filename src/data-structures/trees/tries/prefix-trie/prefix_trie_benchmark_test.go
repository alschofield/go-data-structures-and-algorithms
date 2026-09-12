//go:build contract

package prefix_trie

import (
	"fmt"
	"testing"
)

func BenchmarkPrefixTrie(b *testing.B) {
	for _, size := range []int{256, 1_024} {
		keys := trieKeys(size)

		b.Run(fmt.Sprintf("insert/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				trie := NewPrefixTrie()
				for _, key := range keys {
					if !trie.Insert(key) {
						b.Fatal("Insert() rejected a unique key")
					}
				}
			}
		})

		b.Run(fmt.Sprintf("contains/Size=%d", size), func(b *testing.B) {
			trie := trieFromKeys(b, keys)
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if !trie.Contains(keys[iteration%size]) {
					b.Fatal("Contains() missed an inserted key")
				}
			}
		})

		b.Run(fmt.Sprintf("shared-prefix/Size=%d", size), func(b *testing.B) {
			trie := trieFromKeys(b, keys)
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				if !trie.StartsWith("shared-") {
					b.Fatal("StartsWith() missed the shared prefix")
				}
			}
		})

		b.Run(fmt.Sprintf("remove/Size=%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for iteration := 0; iteration < b.N; iteration++ {
				trie := trieFromKeys(b, keys)
				if !trie.Remove(keys[iteration%size]) {
					b.Fatal("Remove() missed an inserted key")
				}
			}
		})
	}
}

func trieFromKeys(b *testing.B, keys []string) *PrefixTrie {
	trie := NewPrefixTrie()
	for _, key := range keys {
		if !trie.Insert(key) {
			b.Fatal("trie setup rejected a unique key")
		}
	}
	return trie
}

func trieKeys(size int) []string {
	keys := make([]string, size)
	for index := range keys {
		keys[index] = fmt.Sprintf("shared-%04d", index)
	}
	return keys
}
