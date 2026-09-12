package prefix_trie

import (
	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

type PrefixTrie struct {
	char_count int
	word_count int
	next_key   int
	root       *graph.Node[string]
}

func NewPrefixTrie() *PrefixTrie {
	var root_char string = ""
	root := &graph.Node[string]{
		Key:   0,
		Value: &root_char,
	}

	return &PrefixTrie{
		char_count: 0,
		word_count: 0,
		next_key:   1,
		root:       root,
	}
}

func (pt *PrefixTrie) Insert(str string) bool {
	return false
}

func (pt *PrefixTrie) Contains(str string) bool {
	return false
}

func (pt *PrefixTrie) StartsWith(str string) bool {
	return false
}

func (pt *PrefixTrie) Remove(str string) bool {
	return false
}

func (pt *PrefixTrie) Len() int {
	return pt.char_count
}
