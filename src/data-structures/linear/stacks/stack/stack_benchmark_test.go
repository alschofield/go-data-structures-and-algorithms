//go:build contract

package stack

import (
	"strconv"
	"testing"
)

var stackBenchmarkBool bool
var stackBenchmarkValue int

func BenchmarkStackPush(b *testing.B) {
	for _, size := range []int{0, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			stack := NewStack[int]()
			for value := 0; value < size; value++ {
				stack.Push(value)
			}

			b.ReportAllocs()
			b.ResetTimer()
			for value := 0; value < b.N; value++ {
				stackBenchmarkBool = stack.Push(value)
			}
			b.StopTimer()

			if !stackBenchmarkBool || stack.Len() != size+b.N {
				b.Fatalf("Push() produced Len() = %d, want %d", stack.Len(), size+b.N)
			}
		})
	}
}

func BenchmarkStackPop(b *testing.B) {
	for _, size := range []int{0, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			stack := NewStack[int]()
			for value := 0; value < size+b.N; value++ {
				stack.Push(value)
			}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				value, err := stack.Pop()
				if err != nil {
					b.Fatal(err)
				}
				stackBenchmarkValue = value
			}
			b.StopTimer()

			if stack.Len() != size || stackBenchmarkValue != size {
				b.Fatalf("Pop() left Len() = %d and returned %d last, want %d and %d", stack.Len(), stackBenchmarkValue, size, size)
			}
		})
	}
}

func BenchmarkStackPeek(b *testing.B) {
	for _, size := range []int{1, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			stack := NewStack[int]()
			for value := 0; value < size; value++ {
				stack.Push(value)
			}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				value, err := stack.Peek()
				if err != nil {
					b.Fatal(err)
				}
				stackBenchmarkValue = value
			}
			b.StopTimer()

			if stack.Len() != size || stackBenchmarkValue != size-1 {
				b.Fatalf("Peek() left Len() = %d and returned %d, want %d and %d", stack.Len(), stackBenchmarkValue, size, size-1)
			}
		})
	}
}
