package prefix_trie

import (
	"unicode/utf8"

	"github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

type PrefixTrie struct {
	char_count int
	word_count int
	next_key   int
	root       *graph.Node[string]
}

func NewPrefixTrie() *PrefixTrie {
	children := make(map[rune]*graph.Node[string])
	var root_char string = ""
	root := &graph.Node[string]{
		Value:    &root_char,
		Children: children,
	}

	return &PrefixTrie{
		char_count: 0,
		word_count: 0,
		root:       root,
	}
}

func (pt *PrefixTrie) Insert(str string) bool {
	candidate := pt.root
	for _, r := range str {
		if candidate.Children[r] == nil {
			char := string(r)
			children := make(map[rune]*graph.Node[string])
			candidate.Children[r] = &graph.Node[string]{
				Value:    &char,
				Children: children,
			}

			pt.char_count++
		}

		candidate = candidate.Children[r]
	}

	candidate.IsEndOfWord = true
	pt.word_count++

	return true
}

func (pt *PrefixTrie) Contains(str string) bool {
	candidate := pt.root
	for _, r := range str {
		if candidate.Children[r] != nil {
			candidate = candidate.Children[r]
		} else {
			return false
		}
	}

	return candidate.IsEndOfWord
}

func (pt *PrefixTrie) StartsWith(str string) bool {
	candidate := pt.root
	for _, r := range str {
		if candidate.Children[r] != nil {
			candidate = candidate.Children[r]
		} else {
			return false
		}
	}

	return true
}

func recurse(pt *PrefixTrie, node *graph.Node[string], substring string) bool {
	if substring == "" {
		return true
	}

	first_rune, _ := utf8.DecodeRuneInString(substring)

	if node.Children[first_rune] != nil {
		if recurse(pt, node.Children[first_rune], substring[1:]) {
			if len(node.Children[first_rune].Children) == 0 {
				node.Children[first_rune] = nil
				pt.char_count--
			}

			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (pt *PrefixTrie) Remove(str string) bool {
	if recurse(pt, pt.root, str) {
		pt.word_count--
		return true
	} else {
		return false
	}
}

func (pt *PrefixTrie) Len() int {
	return pt.char_count
}
