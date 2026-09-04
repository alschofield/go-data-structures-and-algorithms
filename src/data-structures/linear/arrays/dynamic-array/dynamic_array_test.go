//go:build contract

package dynamic_array

import "testing"

func TestDynamicArray(t *testing.T) {
	a := NewDynamicArray[int]()
	for _, value := range []int{1, 3} {
		if !a.Insert(a.Len(), value) {
			t.Fatal("append insert failed")
		}
	}
	if !a.Insert(1, 2) || a.Len() != 3 {
		t.Fatal("insert must shift and grow length")
	}
	for index, want := range []int{1, 2, 3} {
		if got, ok := a.Get(index); !ok || got != want {
			t.Fatalf("Get(%d) = (%d, %t), want (%d, true)", index, got, ok, want)
		}
	}
	if got, ok := a.Remove(1); !ok || got != 2 {
		t.Fatal("Remove must return and remove indexed value")
	}
	if a.Insert(-1, 0) || a.Len() != 2 {
		t.Fatal("invalid insert must preserve state")
	}
}
