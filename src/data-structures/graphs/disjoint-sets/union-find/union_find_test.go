//go:build contract

package union_find

import (
	"errors"
	"testing"
)

func TestUnionFind(t *testing.T) {
	unionFind, err := NewUnionFind(4)
	if err != nil {
		t.Fatalf("constructor error = %v", err)
	}
	for _, pair := range [][2]int{{0, 1}, {2, 3}, {1, 2}} {
		if merged, err := unionFind.Union(pair[0], pair[1]); err != nil || !merged {
			t.Fatalf("Union(%v) must merge distinct sets", pair)
		}
	}
	if merged, err := unionFind.Union(0, 3); err != nil || merged {
		t.Fatal("existing connection must not change set count")
	}
	if connected, err := unionFind.Connected(0, 3); err != nil || !connected || unionFind.SetCount() != 1 {
		t.Fatal("all elements must be connected in one set")
	}
	if _, err := NewUnionFind(-1); !errors.Is(err, ErrInvalidCapacity) {
		t.Fatalf("negative count error = %v, want ErrInvalidCapacity", err)
	}
}

func TestUnionFindCompressionAndInvalidIndexes(t *testing.T) {
	union_find, err := NewUnionFind(4)
	if err != nil {
		t.Fatal(err)
	}
	for _, pair := range [][2]int{{0, 1}, {2, 3}, {0, 2}} {
		if merged, err := union_find.Union(pair[0], pair[1]); err != nil || !merged {
			t.Fatalf("Union(%v) = (%t, %v), want (true, nil)", pair, merged, err)
		}
	}
	root, ok, err := union_find.Find(3)
	if err != nil || !ok || union_find.nodes[3].Parent != union_find.nodes[root] {
		t.Fatal("Find must compress a traversed path toward its root")
	}
	if _, _, err := union_find.Find(-1); !errors.Is(err, ErrInvalidIndex) {
		t.Fatalf("Find(-1) error = %v, want ErrInvalidIndex", err)
	}
	if _, err := union_find.Connected(0, 4); !errors.Is(err, ErrInvalidIndex) {
		t.Fatalf("Connected(0, 4) error = %v, want ErrInvalidIndex", err)
	}
}
