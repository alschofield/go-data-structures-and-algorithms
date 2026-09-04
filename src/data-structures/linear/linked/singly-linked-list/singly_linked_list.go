package singly_linked_list

import (
	"errors"
)

// ErrInvalidIndex reports an index outside the range accepted by Get, Insert, or Remove.
var ErrInvalidIndex = errors.New("function requires a correct index.")

// Node holds one value and a link to the following node in the list.
type Node[T any] struct {
	value T
	next  *Node[T]
}

// SinglyLinkedList tracks the first node and the number of values it contains.
// It deliberately has no tail pointer, so operations at the back traverse from head.
type SinglyLinkedList[T any] struct {
	size int
	head *Node[T]
}

// NewSinglyLinkedList returns an initialized empty list.
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// PushFront links a new node before the current head.
func (ll *SinglyLinkedList[T]) PushFront(value T) bool {
	var node *Node[T] = &Node[T]{value: value, next: ll.head}
	ll.head = node
	ll.size++
	return true
}

// PushBack walks to the final node, then links a new node after it.
func (ll *SinglyLinkedList[T]) PushBack(value T) bool {
	// An empty list has no final node, so adding at the back is the same as adding at the front.
	if ll.head == nil {
		return ll.PushFront(value)
	}

	// Start at head and follow each next link until reaching the tail.
	var node *Node[T] = ll.head
	for node.next != nil {
		node = node.next
	}

	// A nil next pointer marks the new node as the tail.
	var empty *Node[T]
	var new *Node[T] = &Node[T]{value, empty}
	node.next = new
	ll.size++
	return true
}

// PopFront removes and returns the head value when one exists.
func (ll *SinglyLinkedList[T]) PopFront() (T, bool) {
	// Return T's zero value for an empty list without changing its state.
	if ll.size == 0 {
		var zero T
		return zero, false
	}

	// Advance head past the removed node before reducing the recorded size.
	var node *Node[T] = ll.head
	ll.head = node.next
	ll.size--
	return node.value, true
}

// PopBack removes and returns the tail value when one exists.
func (ll *SinglyLinkedList[T]) PopBack() (T, bool) {
	// Return T's zero value for an empty list without changing its state.
	if ll.size == 0 {
		var zero T
		return zero, false
	}

	// Keep a cursor so it can stop at the node immediately before the tail.
	var node *Node[T] = ll.head
	// A one-node list becomes empty after its head is removed.
	if ll.size == 1 {
		ll.head = nil
		ll.size--
		return node.value, true
	}

	// Stop at the penultimate node, whose next node is the value to remove.
	for node.next.next != nil {
		node = node.next
	}

	// Disconnect the tail so the penultimate node becomes the new tail.
	var temp = node.next
	node.next = nil
	ll.size--
	return temp.value, true
}

// Get returns the value at index without changing the list.
func (ll *SinglyLinkedList[T]) Get(index int) (T, bool, error) {
	// Every readable index must identify an existing node.
	if index < 0 || index >= ll.size {
		var zeroed T
		return zeroed, false, ErrInvalidIndex
	}

	// Walk forward exactly index links from head.
	var node *Node[T] = ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node.value, true, nil
}

// Insert links value at index, allowing index Len() to append.
func (ll *SinglyLinkedList[T]) Insert(index int, value T) (bool, error) {
	// Insertion may target every existing position plus the position after the tail.
	if index < 0 || index > ll.size {
		return false, ErrInvalidIndex
	}

	// The head case needs no traversal.
	if index == 0 {
		return ll.PushFront(value), nil
	}

	// Stop at the node that will precede the inserted node.
	var node *Node[T] = ll.head
	for i := 0; i < (index - 1); i++ {
		node = node.next
	}

	// Preserve the successor before inserting the new link between the two nodes.
	var new *Node[T] = &Node[T]{value: value, next: node.next}
	node.next = new
	ll.size++
	return true, nil
}

// Remove unlinks and returns the value at index.
func (ll *SinglyLinkedList[T]) Remove(index int) (T, bool, error) {
	// Removal requires index to identify an existing node.
	if index < 0 || index >= ll.size {
		var zeroed T
		return zeroed, false, ErrInvalidIndex
	}

	// Start at head so index zero can update it directly.
	var node *Node[T] = ll.head
	// Removing head advances head to the second node, if any.
	if index == 0 {
		ll.head = node.next
		ll.size--
		return node.value, true, nil
	}

	// Stop at the node immediately before the one to remove.
	for i := 0; i < (index - 1); i++ {
		node = node.next
	}

	// Bypass the removed node by linking its predecessor to its successor.
	var temp *Node[T] = node.next
	node.next = temp.next
	ll.size--
	return temp.value, true, nil
}

// Len returns the number of nodes recorded in the list.
func (ll *SinglyLinkedList[T]) Len() int {
	return ll.size
}

// IsEmpty reports whether the list has no nodes.
func (ll *SinglyLinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
