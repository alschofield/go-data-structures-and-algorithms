package doubly_linked_list

import (
	"errors"

	graph "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

// ErrInvalidIndex reports an index outside the range accepted by Get, Insert, or Remove.
var ErrInvalidIndex = errors.New("function requires a valid index.")

// DoublyLinkedList tracks both ends so operations at either end avoid traversal.
type DoublyLinkedList[T any] struct {
	size int
	head *graph.Node[T]
	tail *graph.Node[T]
}

// NewDoublyLinkedList returns an initialized empty list.
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// PushFront links value before the current head.
func (ll *DoublyLinkedList[T]) PushFront(value T) bool {
	// The former head, if present, follows the new node.
	var node *graph.Node[T] = &graph.Node[T]{Value: &value, Prev: nil, Next: ll.head}
	if ll.size == 0 {
		// The first node is both ends of the list.
		ll.tail = node
	} else {
		// The old head must point back to its new predecessor.
		ll.head.Prev = node
	}

	// The new node now begins the forward chain.
	ll.head = node
	ll.size++
	return true
}

// PushBack links value after the current tail.
func (ll *DoublyLinkedList[T]) PushBack(value T) bool {
	// The former tail, if present, precedes the new node.
	var node *graph.Node[T] = &graph.Node[T]{Value: &value, Prev: ll.tail, Next: nil}
	if ll.size == 0 {
		// The first node is both ends of the list.
		ll.head = node
	} else {
		// The old tail must point forward to its new successor.
		ll.tail.Next = node
	}

	// The new node now ends the forward chain.
	ll.tail = node
	ll.size++
	return true
}

// PopFront removes and returns the head value when one exists.
func (ll *DoublyLinkedList[T]) PopFront() (T, bool) {
	// An empty list has no head to remove.
	if ll.size == 0 {
		var zero T
		return zero, false
	}

	// Save the old head, then advance head to its successor.
	var return_node *graph.Node[T] = ll.head
	ll.head = return_node.Next

	if ll.head != nil {
		// The new head has no predecessor.
		ll.head.Prev = nil
	}

	if ll.size == 1 {
		// Removing the only node clears the opposite end too.
		ll.tail = nil
	}

	ll.size--

	// Return the removed value after the list state is consistent.
	return *return_node.Value, true
}

// PopBack removes and returns the tail value when one exists.
func (ll *DoublyLinkedList[T]) PopBack() (T, bool) {
	// An empty list has no tail to remove.
	if ll.size == 0 {
		var zero T
		return zero, false
	}

	// Save the old tail, then retreat tail to its predecessor.
	var return_node *graph.Node[T] = ll.tail
	ll.tail = return_node.Prev

	if ll.tail != nil {
		// The new tail has no successor.
		ll.tail.Next = nil
	}

	if ll.size == 1 {
		// Removing the only node clears the opposite end too.
		ll.head = nil
	}

	ll.size--

	// Return the removed value after the list state is consistent.
	return *return_node.Value, true
}

// Get returns the value at index without changing the list.
func (ll *DoublyLinkedList[T]) Get(index int) (T, bool, error) {
	// Readable indexes must identify an existing node.
	if index < 0 || index >= ll.size {
		var zero T
		return zero, false, ErrInvalidIndex
	}

	var temp *graph.Node[T]
	// Starting from the nearer end limits traversal to roughly half the list.
	if index < (ll.size / 2) {
		temp = ll.head
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
	} else if index >= ll.size/2 {
		temp = ll.tail
		for i := ll.size - 1; i > index; i-- {
			temp = temp.Prev
		}
	}

	// The traversal found the requested existing node.
	return *temp.Value, true, nil
}

// Remove unlinks and returns the value at index.
func (ll *DoublyLinkedList[T]) Remove(index int) (T, bool, error) {
	// Removable indexes must identify an existing node.
	if index < 0 || index >= ll.size {
		var zero T
		return zero, false, ErrInvalidIndex
	}

	var return_node *graph.Node[T]
	// Reuse the end operations because they also maintain head and tail.
	if index == 0 {
		return_node, status := ll.PopFront()
		return return_node, status, nil
	}

	if index == ll.size-1 {
		return_node, status := ll.PopBack()
		return return_node, status, nil
	}

	var temp *graph.Node[T]
	// Traverse from the closer end to the interior node.
	if index < (ll.size / 2) {
		temp = ll.head
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
	} else if index >= ll.size/2 {
		temp = ll.tail
		for i := ll.size - 1; i > index; i-- {
			temp = temp.Prev
		}
	}

	// Link the removed node's neighbors directly to each other.
	temp.Next.Prev = temp.Prev
	temp.Prev.Next = temp.Next
	return_node = temp
	ll.size--

	// The node's value remains available after it is unlinked.
	return *return_node.Value, true, nil
}

// Insert links value at index, allowing Len() to append.
func (ll *DoublyLinkedList[T]) Insert(index int, value T) (bool, error) {
	// Insertion may target every existing position and the position after tail.
	if index < 0 || index > ll.size {
		return false, ErrInvalidIndex
	}

	// The end cases already maintain the corresponding boundary pointers.
	if index == 0 {
		return ll.PushFront(value), nil
	}

	if index == ll.size {
		return ll.PushBack(value), nil
	}

	var temp *graph.Node[T]
	// Find the current node that will follow the inserted node.
	if index < (ll.size / 2) {
		temp = ll.head
		for i := 0; i < index; i++ {
			temp = temp.Next
		}
	} else if index >= ll.size/2 {
		temp = ll.tail
		for i := ll.size - 1; i > index; i-- {
			temp = temp.Prev
		}
	}

	// Stitch the new node between its predecessor and successor.
	var new_node *graph.Node[T] = &graph.Node[T]{Value: &value}
	temp.Prev.Next = new_node
	new_node.Prev = temp.Prev
	new_node.Next = temp
	temp.Prev = new_node
	ll.size++

	return true, nil
}

// Len returns the number of nodes recorded in the list.
func (ll *DoublyLinkedList[T]) Len() int {
	return ll.size
}

// IsEmpty reports whether the list has no nodes.
func (ll *DoublyLinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
