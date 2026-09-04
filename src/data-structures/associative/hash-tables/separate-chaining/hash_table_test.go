//go:build contract

package separate_chaining

import "testing"

func TestHashTable(t *testing.T) {
	equal := func(a, b string) bool { return a == b }
	table := NewHashTable[string, int](2, func(string) uint { return 0 }, equal)
	if table == nil {
		t.Fatal("nonzero capacity must create table")
	}
	table.Set("a", 1)
	table.Set("b", 2)
	if table.Cap() != 2 {
		t.Fatal("Set must not resize")
	}
	if prior, replaced := table.Set("a", 3); !replaced || prior != 1 {
		t.Fatal("Set must return replaced value")
	}
	table.SetResizing("c", 3)
	if table.Cap() != 4 || table.Len() != 3 {
		t.Fatal("SetResizing must rehash before exceeding 0.75 load")
	}
	if got, ok := table.Get("b"); !ok || got != 2 {
		t.Fatal("rehash must retain collision-chain entries")
	}
}
