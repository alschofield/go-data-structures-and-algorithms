package stack

import "errors"

// ErrEmptyStack reports an attempt to read from a stack with no items.
var ErrEmptyStack = errors.New("stack is empty.")

// Stack stores values in insertion order, with the last slice element as the top.
type Stack[T any] struct {
	// size records the number of values currently available to callers.
	size int
	// items is the backing slice that holds the stack values from bottom to top.
	items []T
}

// NewStack creates an empty stack with no allocated backing storage.
func NewStack[T any]() *Stack[T] {
	// Return a zero-valued stack so allocation happens only on the first Push.
	return &Stack[T]{}
}

// Push places value at the top of the stack.
func (s *Stack[T]) Push(value T) bool {
	// Append makes value the last element, which is this implementation's top.
	s.items = append(s.items, value)
	// Keep the explicit size synchronized with the accessible slice length.
	s.size++
	// Push cannot fail for a valid receiver, so report success.
	return true
}

// Pop removes and returns the value at the top of the stack.
func (s *Stack[T]) Pop() (T, error) {
	// An empty stack has no top value to remove.
	if s.size == 0 {
		// Declare T's zero value for the required empty-result return.
		var empty T
		// Return the stable sentinel without changing stack state.
		return empty, ErrEmptyStack
	}

	// Create T's zero value before clearing the removed backing-slice slot.
	var zeroed T
	// Read the last element because it is the current stack top.
	item := s.items[s.size-1]
	// Drop references held by the removed slot so its referent can be collected.
	s.items[s.size-1] = zeroed
	// Record that the top value is no longer part of the stack.
	s.size--
	// Shorten the accessible slice while retaining its capacity for future Push calls.
	s.items = s.items[:s.size]
	// Return the removed value and a successful nil error.
	return item, nil
}

// Peek returns the value at the top of the stack without removing it.
func (s *Stack[T]) Peek() (T, error) {
	// An empty stack has no top value to inspect.
	if s.size == 0 {
		// Declare T's zero value for the required empty-result return.
		var empty T
		// Return the stable sentinel without changing stack state.
		return empty, ErrEmptyStack
	}

	// Read the last element without changing either size or the backing slice.
	return s.items[s.size-1], nil
}

// Len returns the number of values currently in the stack.
func (s *Stack[T]) Len() int {
	// size is maintained by Push and Pop, so no slice traversal is necessary.
	return s.size
}

// IsEmpty reports whether the stack contains no values.
func (s *Stack[T]) IsEmpty() bool {
	// A stack is empty exactly when its tracked size is zero.
	return s.size == 0
}
