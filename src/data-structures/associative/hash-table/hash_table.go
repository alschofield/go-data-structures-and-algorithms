// Package hash_table implements a generic separate-chaining hash table.
package hash_table

// errors supplies sentinel errors for invalid constructor arguments.
import (
	"errors"

	graph "github.com/alschofield/go-data-structures-and-algorithms/src/data-structures/graphs/graph"
)

// Constructor errors describe invalid configuration before a table exists.
var (
	// ErrInvalidCapacity reports a non-positive bucket count.
	ErrInvalidCapacity = errors.New("function requires a valid capacity.")
	// ErrNilHash reports a missing function that assigns keys to buckets.
	ErrNilHash = errors.New("function needs a valid hash function.")
	// ErrNilEqual reports a missing function that compares keys in a chain.
	ErrNilEqual = errors.New("function needs a valid equal function.")
)

// entry stores one hash key/value pair inside a shared chain node.
type entry[K any, V any] struct {
	// key stays unchanged when an equal replacement updates value.
	key K
	// value is returned by Get and replaced by Set for an equal key.
	value V
}

// HashTable owns bucket storage plus the functions that interpret its keys.
type HashTable[K any, V any] struct {
	// size is the number of stored keys across all buckets.
	size uint
	// capacity is the number of bucket heads in items.
	capacity uint
	// hash maps a key to an unsigned value before bucket reduction.
	hash func(K) uint
	// equal decides whether a searched key matches a stored key.
	equal func(K, K) bool
	// items holds one linked-list head per bucket.
	items []*graph.Node[entry[K, V]]
}

// NewHashTable validates its configuration and creates an empty bucket array.
func NewHashTable[K any, V any](initial_capacity int, hash func(K) uint, equal func(K, K) bool) (*HashTable[K, V], error) {
	// A zero-length bucket array would make modulo reduction invalid.
	if initial_capacity <= 0 {
		return nil, ErrInvalidCapacity
		// A hash function is required to choose a bucket for every operation.
	} else if hash == nil {
		return nil, ErrNilHash
		// Equality is required to distinguish keys that share a bucket.
	} else if equal == nil {
		return nil, ErrNilEqual
	}

	// Allocate exactly one empty chain head for each requested bucket.
	return &HashTable[K, V]{
		size:     0,
		capacity: uint(initial_capacity),
		hash:     hash,
		equal:    equal,
		items:    make([]*graph.Node[entry[K, V]], initial_capacity),
	}, nil
}

// Set inserts a key or replaces its value without changing table capacity.
func (ht *HashTable[K, V]) Set(key K, value V) (V, bool) {
	// Reduce the hash to a valid bucket index.
	var bucket uint = ht.hash(key) % ht.capacity

	// Scan the chain so equal keys replace rather than duplicate an entry.
	for list := ht.items[bucket]; list != nil; list = list.Next {
		if ht.equal(key, list.Value.key) {
			// Preserve the original stored key and return its previous value.
			old_value := list.Value.value
			list.Value.value = value
			return old_value, true
		}
	}

	// Prepend a new collision node so insertion does not traverse the chain again.
	new_entry := entry[K, V]{key: key, value: value}
	ht.items[bucket] = &graph.Node[entry[K, V]]{Value: &new_entry, Next: ht.items[bucket]}

	// Only a new key increases the number of entries.
	ht.size++

	// New insertions have no prior value of V.
	var zero V
	return zero, false
}

// SetResizing grows and rehashes before a new insertion would exceed 75% load.
func (ht *HashTable[K, V]) SetResizing(key K, value V) (V, bool) {
	// Replacements do not increase load, so only a new threshold-crossing key grows.
	if ((ht.size+1)*4 > ht.capacity*3) && !ht.Contains(key) {
		// Doubling keeps growth geometric and preserves a positive capacity.
		var new_capacity uint = ht.capacity * 2
		// New bucket heads receive every existing node under the new modulus.
		var items []*graph.Node[entry[K, V]] = make([]*graph.Node[entry[K, V]], new_capacity)
		var candidate *graph.Node[entry[K, V]]
		for i := 0; i < int(ht.capacity); i++ {
			candidate = ht.items[i]

			// Empty buckets have no chain to move.
			if candidate == nil {
				continue
			}

			// Relink each old chain node into its new bucket.
			for candidate != nil {
				// Save the old successor before overwriting candidate.next.
				next := candidate.Next
				// The new capacity can change this node's bucket.
				new_bucket := ht.hash(candidate.Value.key) % uint(new_capacity)
				// Prepend the node to its new collision chain.
				candidate.Next = items[new_bucket]
				items[new_bucket] = candidate
				// Continue through the old chain using the saved link.
				candidate = next
			}
		}

		// Publish the fully built bucket array only after rehashing completes.
		ht.items = items
		ht.capacity = new_capacity
	}

	// Set performs the actual insert or replacement after any required growth.
	return ht.Set(key, value)
}

// Get returns the value for key, or the zero value and false when it is absent.
func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	// Search only the chain selected by this key's hash.
	var bucket uint = ht.hash(key) % ht.capacity
	var node *graph.Node[entry[K, V]] = ht.items[bucket]
	for node != nil {
		if ht.equal(key, node.Value.key) {
			return node.Value.value, true
		}

		// Advance through collisions until a match or chain end.
		node = node.Next
	}

	// Absence is a normal lookup result, not an error.
	var zero V
	return zero, false
}

// Remove unlinks and returns key's entry, or reports false without mutation.
func (ht *HashTable[K, V]) Remove(key K) (V, bool) {
	// Track both the current node and its predecessor in the target chain.
	var bucket uint = ht.hash(key) % ht.capacity
	var node *graph.Node[entry[K, V]] = ht.items[bucket]
	var prev *graph.Node[entry[K, V]]
	for node != nil {
		if ht.equal(key, node.Value.key) {
			if prev == nil {
				// Removing the head changes the bucket's chain head.
				ht.items[bucket] = node.Next
			} else {
				// Removing elsewhere skips the node from its predecessor.
				prev.Next = node.Next
			}

			// A successful removal decreases entry count and detaches the node.
			ht.size--
			node.Next = nil

			return node.Value.value, true
		}

		// Advance both pointers while preserving their predecessor relationship.
		prev = node
		node = node.Next
	}

	// An absent key leaves every bucket and counter unchanged.
	var zero V
	return zero, false
}

// Contains reports whether Get finds a value for key.
func (ht *HashTable[K, V]) Contains(key K) bool {
	_, status := ht.Get(key)
	return status
}

// Cap returns the current number of buckets.
func (ht *HashTable[K, V]) Cap() int {
	return int(ht.capacity)
}

// Len returns the current number of entries.
func (ht *HashTable[K, V]) Len() int {
	return int(ht.size)
}

// IsEmpty reports whether no entries are stored.
func (ht *HashTable[K, V]) IsEmpty() bool {
	return int(ht.size) == 0
}
