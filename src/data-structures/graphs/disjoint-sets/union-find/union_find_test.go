//go:build contract

package union_find

import "testing"

func TestUnionFind(t *testing.T) {
	unionFind := NewUnionFind(4)
	for _, pair := range [][2]int{{0, 1}, {2, 3}, {1, 2}} {
		if merged, ok := unionFind.Union(pair[0], pair[1]); !ok || !merged {
			t.Fatalf("Union(%v) must merge distinct sets", pair)
		}
	}
	if merged, ok := unionFind.Union(0, 3); !ok || merged {
		t.Fatal("existing connection must not change set count")
	}
	if connected, ok := unionFind.Connected(0, 3); !ok || !connected || unionFind.SetCount() != 1 {
		t.Fatal("all elements must be connected in one set")
	}
}
