//go:build contract

package queue

import (
	"strconv"
	"testing"
)

var queueBenchmarkBool bool
var queueBenchmarkValue int

func BenchmarkQueueEnqueue(b *testing.B) {
	for _, size := range []int{0, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			queue := NewQueue[int]()
			for value := 0; value < size; value++ {
				queue.Enqueue(value)
			}

			b.ReportAllocs()
			b.ResetTimer()
			for value := 0; value < b.N; value++ {
				queueBenchmarkBool = queue.Enqueue(value)
			}
			b.StopTimer()

			if !queueBenchmarkBool || queue.Len() != size+b.N {
				b.Fatalf("Enqueue() produced Len() = %d, want %d", queue.Len(), size+b.N)
			}
		})
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	for _, size := range []int{0, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			queue := NewQueue[int]()
			for value := 0; value < size+b.N; value++ {
				queue.Enqueue(value)
			}

			b.ReportAllocs()
			b.ResetTimer()
			for range b.N {
				value, err := queue.Dequeue()
				if err != nil {
					b.Fatal(err)
				}
				queueBenchmarkValue = value
			}
			b.StopTimer()

			if queue.Len() != size || queueBenchmarkValue != b.N-1 {
				b.Fatalf("Dequeue() left Len() = %d and returned %d last, want %d and %d", queue.Len(), queueBenchmarkValue, size, b.N-1)
			}
		})
	}
}

func BenchmarkQueuePeek(b *testing.B) {
	for _, size := range []int{1, 1_024, 65_536} {
		b.Run("Size="+strconv.Itoa(size), func(b *testing.B) {
			queue := NewQueue[int]()
			for value := 0; value < size; value++ {
				queue.Enqueue(value)
			}

			b.ReportAllocs()
			b.ResetTimer()
			for range b.N {
				value, err := queue.Peek()
				if err != nil {
					b.Fatal(err)
				}
				queueBenchmarkValue = value
			}
			b.StopTimer()

			if queue.Len() != size || queueBenchmarkValue != 0 {
				b.Fatalf("Peek() left Len() = %d and returned %d, want %d and 0", queue.Len(), queueBenchmarkValue, size)
			}
		})
	}
}

func BenchmarkQueueWraparound(b *testing.B) {
	queue := NewQueue[int]()
	queue.Enqueue(0)
	queue.Enqueue(1)

	b.ReportAllocs()
	b.ResetTimer()
	for value := 0; value < b.N; value++ {
		dequeued, err := queue.Dequeue()
		if err != nil {
			b.Fatal(err)
		}
		queueBenchmarkValue = dequeued
		queueBenchmarkBool = queue.Enqueue(value + 2)
	}
	b.StopTimer()

	if !queueBenchmarkBool || queue.Len() != 2 {
		b.Fatalf("wraparound left Len() = %d, want 2", queue.Len())
	}
}

func BenchmarkQueueGrowWrapped(b *testing.B) {
	const size = 1_024

	b.ReportAllocs()
	b.ResetTimer()
	for value := 0; value < b.N; value++ {
		b.StopTimer()
		queue := fullWrappedQueue(size)
		b.StartTimer()
		queueBenchmarkBool = queue.Enqueue(value)
		b.StopTimer()

		if !queueBenchmarkBool || queue.Len() != size+1 || queue.items[0] != size/2 {
			b.Fatalf("wrapped growth produced invalid queue state")
		}
	}
}

func fullWrappedQueue(size int) *Queue[int] {
	queue := NewQueue[int]()
	for value := 0; value < size; value++ {
		queue.Enqueue(value)
	}
	for range size / 2 {
		queue.Dequeue()
	}
	for value := size; value < size+size/2; value++ {
		queue.Enqueue(value)
	}
	return queue
}
