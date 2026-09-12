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

// NewPrefixTrie creates an empty root with no character edges.
func NewPrefixTrie() *PrefixTrie {
	children := make(map[rune]*graph.Node[string])
	var root_char string = ""
	root := &graph.Node[string]{
		Key:      0,
		Value:    &root_char,
		Children: children,
	}

	return &PrefixTrie{
		char_count: 0,
		word_count: 0,
		next_key:   1,
		root:       root,
	}
}

// Insert adds a key's missing rune path and marks its terminal node as a word.
func (pt *PrefixTrie) Insert(str string) bool {
	candidate := pt.root
	for _, r := range str {
		if candidate.Children[r] == nil {
			// Each new rune gets a stable shared-node key and its own child map.
			char := string(r)
			children := make(map[rune]*graph.Node[string])
			candidate.Children[r] = &graph.Node[string]{
				Key:      pt.next_key,
				Value:    &char,
				Children: children,
			}

			pt.next_key++
			pt.char_count++
		}

		candidate = candidate.Children[r]
	}

	if !candidate.IsEndOfWord {
		candidate.IsEndOfWord = true
		pt.word_count++
		return true
	} else {
		return false
	}
}

// Contains reports whether str reaches a terminal word node.
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

// StartsWith reports whether str reaches any existing trie node.
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

// recurse clears one terminal marker and prunes only unreachable suffix nodes.
func recurse(pt *PrefixTrie, node *graph.Node[string], substring string) bool {
	if substring == "" {
		if !node.IsEndOfWord {
			return false
		}

		node.IsEndOfWord = false
		return true
	}

	first_rune, size := utf8.DecodeRuneInString(substring)

	if node.Children[first_rune] != nil {
		if recurse(pt, node.Children[first_rune], substring[size:]) {
			// A child remains necessary when it starts another word or has descendants.
			if len(node.Children[first_rune].Children) == 0 && !node.Children[first_rune].IsEndOfWord {
				delete(node.Children, first_rune)
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

// Remove clears a stored nonempty key and prunes any now-unused suffix nodes.
func (pt *PrefixTrie) Remove(str string) bool {
	if str == "" {
		return false
	}

	if recurse(pt, pt.root, str) {
		pt.word_count--
		return true
	} else {
		return false
	}
}

// Len returns the number of complete stored words rather than structural runes.
func (pt *PrefixTrie) Len() int {
	return pt.word_count
}
