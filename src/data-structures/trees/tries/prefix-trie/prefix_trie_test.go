//go:build contract

package prefix_trie

import "testing"

func TestPrefixTrie(t *testing.T) {
	trie := NewPrefixTrie()
	for _, key := range []string{"car", "cart", "cat"} {
		if !trie.Insert(key) {
			t.Fatal("new insert failed")
		}
	}
	if trie.Insert("car") {
		t.Fatal("duplicate insert must be idempotent")
	}
	if !trie.StartsWith("ca") || !trie.StartsWith("") || !trie.Remove("cart") || !trie.Contains("car") || trie.Contains("cart") {
		t.Fatal("remove must prune only unneeded nodes")
	}
}
