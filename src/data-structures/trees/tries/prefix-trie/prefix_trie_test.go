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

func TestPrefixTrieGuardsMissingRemovalAndPreservesUnicodePrefixes(t *testing.T) {
	trie := NewPrefixTrie()
	for _, key := range []string{"café", "caféteria"} {
		if !trie.Insert(key) {
			t.Fatalf("Insert(%q) = false, want true", key)
		}
	}
	if trie.Remove("caf") || trie.Len() != 2 {
		t.Fatal("removing a prefix that is not a word must preserve trie state")
	}
	if !trie.Remove("café") || trie.Contains("café") || !trie.Contains("caféteria") || !trie.StartsWith("café") || trie.Len() != 1 {
		t.Fatal("Unicode removal must preserve longer shared-prefix words")
	}
}
