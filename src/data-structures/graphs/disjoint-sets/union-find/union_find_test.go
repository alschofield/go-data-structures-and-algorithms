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
