//go:build contract

package hash_table

import (
	"errors"
	"testing"
)

func TestNewHashTableRejectsInvalidArguments(t *testing.T) {
	hash := func(string) uint { return 0 }
	equal := func(a, b string) bool { return a == b }

	for _, test := range []struct {
		name string
		call func() (*HashTable[string, int], error)
		want error
	}{
		{name: "zero capacity", call: func() (*HashTable[string, int], error) { return NewHashTable[string, int](0, hash, equal) }, want: ErrInvalidCapacity},
		{name: "negative capacity", call: func() (*HashTable[string, int], error) { return NewHashTable[string, int](-1, hash, equal) }, want: ErrInvalidCapacity},
		{name: "nil hash", call: func() (*HashTable[string, int], error) { return NewHashTable[string, int](1, nil, equal) }, want: ErrNilHash},
		{name: "nil equal", call: func() (*HashTable[string, int], error) { return NewHashTable[string, int](1, hash, nil) }, want: ErrNilEqual},
	} {
		t.Run(test.name, func(t *testing.T) {
			table, err := test.call()
			if table != nil || !errors.Is(err, test.want) {
				t.Fatalf("NewHashTable() = (%v, %v), want (nil, %v)", table, err, test.want)
			}
		})
	}
}

func TestHashTableSetCollisionChainsAndReplacement(t *testing.T) {
	table := newCollidingTable(t, 2)
	for _, entry := range []struct {
		key   string
		value int
	}{{"a", 1}, {"b", 2}, {"c", 3}} {
		if old, replaced := table.Set(entry.key, entry.value); replaced || old != 0 {
			t.Fatalf("Set(%q, %d) = (%d, %t), want (0, false)", entry.key, entry.value, old, replaced)
		}
	}
	if table.Cap() != 2 {
		t.Fatalf("Set() changed Cap() = %d, want 2", table.Cap())
	}
	assertTableContents(t, table, 2, map[string]int{"a": 1, "b": 2, "c": 3})

	if old, replaced := table.Set("b", 20); !replaced || old != 2 {
		t.Fatalf("Set replacement = (%d, %t), want (2, true)", old, replaced)
	}
	assertTableContents(t, table, 2, map[string]int{"a": 1, "b": 20, "c": 3})
}

func TestHashTableReplacementRetainsOriginalKey(t *testing.T) {
	type key struct {
		id   int
		name string
	}
	original := key{id: 7, name: "original"}
	replacement := key{id: 7, name: "replacement"}
	table, err := NewHashTable[key, int](2, func(key) uint { return 0 }, func(a, b key) bool { return a.id == b.id })
	if err != nil {
		t.Fatalf("NewHashTable() error = %v", err)
	}
	table.Set(original, 1)
	if old, replaced := table.Set(replacement, 2); !replaced || old != 1 {
		t.Fatalf("Set replacement = (%d, %t), want (1, true)", old, replaced)
	}
	if got := table.items[0].key; got != original {
		t.Fatalf("stored key = %#v, want original %#v", got, original)
	}
	if got, ok := table.Get(replacement); !ok || got != 2 {
		t.Fatalf("Get(replacement) = (%d, %t), want (2, true)", got, ok)
	}
}

func TestHashTableSetResizingThresholdReplacementAndRehash(t *testing.T) {
	table, err := NewHashTable[string, int](4, stringHash, stringEqual)
	if err != nil {
		t.Fatalf("NewHashTable() error = %v", err)
	}
	for _, entry := range []struct {
		key   string
		value int
	}{{"a", 1}, {"b", 2}, {"c", 3}} {
		table.SetResizing(entry.key, entry.value)
	}
	assertTableContents(t, table, 4, map[string]int{"a": 1, "b": 2, "c": 3})

	if old, replaced := table.SetResizing("b", 20); !replaced || old != 2 {
		t.Fatalf("SetResizing replacement = (%d, %t), want (2, true)", old, replaced)
	}
	assertTableContents(t, table, 4, map[string]int{"a": 1, "b": 20, "c": 3})

	if old, replaced := table.SetResizing("d", 4); replaced || old != 0 {
		t.Fatalf("SetResizing new value = (%d, %t), want (0, false)", old, replaced)
	}
	assertTableContents(t, table, 8, map[string]int{"a": 1, "b": 20, "c": 3, "d": 4})
}

func TestHashTableGetContainsAndRemove(t *testing.T) {
	table := newCollidingTable(t, 4)
	table.Set("a", 1)
	table.Set("b", 2)
	table.Set("c", 3)

	if got, ok := table.Get("b"); !ok || got != 2 {
		t.Fatalf("Get(b) = (%d, %t), want (2, true)", got, ok)
	}
	if !table.Contains("b") || table.Contains("missing") {
		t.Fatal("Contains() did not distinguish present and absent keys")
	}
	assertMissingDoesNotMutate(t, table, "missing")

	for _, test := range []struct {
		name  string
		key   string
		want  int
		after map[string]int
	}{
		{name: "head", key: "c", want: 3, after: map[string]int{"a": 1, "b": 2}},
		{name: "middle", key: "b", want: 2, after: map[string]int{"a": 1}},
		{name: "tail", key: "a", want: 1, after: map[string]int{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got, ok := table.Remove(test.key); !ok || got != test.want {
				t.Fatalf("Remove(%q) = (%d, %t), want (%d, true)", test.key, got, ok, test.want)
			}
			assertTableContents(t, table, 4, test.after)
		})
	}
	assertMissingDoesNotMutate(t, table, "missing")
}

func TestHashTableLenCapAndIsEmpty(t *testing.T) {
	table := newCollidingTable(t, 3)
	assertTableContents(t, table, 3, map[string]int{})
	table.Set("a", 1)
	assertTableContents(t, table, 3, map[string]int{"a": 1})
	table.Remove("a")
	assertTableContents(t, table, 3, map[string]int{})
}

func newCollidingTable(t *testing.T, capacity int) *HashTable[string, int] {
	t.Helper()
	table, err := NewHashTable[string, int](capacity, func(string) uint { return 0 }, stringEqual)
	if err != nil {
		t.Fatalf("NewHashTable() error = %v", err)
	}
	return table
}

func assertMissingDoesNotMutate(t *testing.T, table *HashTable[string, int], key string) {
	t.Helper()
	beforeLen, beforeCap := table.Len(), table.Cap()
	if got, ok := table.Get(key); ok || got != 0 {
		t.Fatalf("Get(%q) = (%d, %t), want (0, false)", key, got, ok)
	}
	if got, ok := table.Remove(key); ok || got != 0 {
		t.Fatalf("Remove(%q) = (%d, %t), want (0, false)", key, got, ok)
	}
	if table.Len() != beforeLen || table.Cap() != beforeCap {
		t.Fatalf("missing operations changed Len/Cap from (%d, %d) to (%d, %d)", beforeLen, beforeCap, table.Len(), table.Cap())
	}
}

func assertTableContents(t *testing.T, table *HashTable[string, int], wantCap int, want map[string]int) {
	t.Helper()
	if table.Len() != len(want) || table.Cap() != wantCap || table.IsEmpty() != (len(want) == 0) {
		t.Fatalf("state = Len:%d Cap:%d IsEmpty:%t, want Len:%d Cap:%d IsEmpty:%t", table.Len(), table.Cap(), table.IsEmpty(), len(want), wantCap, len(want) == 0)
	}
	for key, value := range want {
		if got, ok := table.Get(key); !ok || got != value {
			t.Fatalf("Get(%q) = (%d, %t), want (%d, true)", key, got, ok, value)
		}
	}
}

func stringHash(value string) uint {
	var hash uint
	for i := range len(value) {
		hash = hash*31 + uint(value[i])
	}
	return hash
}

func stringEqual(a, b string) bool {
	return a == b
}
