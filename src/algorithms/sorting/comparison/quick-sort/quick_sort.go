package quick_sort

import (
	"errors"
)

// ErrNilComparator reports that QuickSort cannot order items without a comparator.
var ErrNilComparator = errors.New("function needs non nil compare function.")

// recursion sorts the inclusive range from left through right in place.
func recursion[T any](items []T, left int, right int, compare func(T, T) int) bool {
	// A range with zero or one item is already sorted.
	if left >= right {
		return true
	}

	// Consider the first, middle, and last items when choosing a pivot.
	var middle int = ((right - left) / 2) + left
	var pivot_index int = middle
	// Select the median of the three candidates according to the comparator.
	if compare(items[left], items[middle]) < 0 {
		if compare(items[middle], items[right]) < 0 {
			pivot_index = middle
		} else if compare(items[left], items[right]) > 0 {
			pivot_index = right
		} else {
			pivot_index = left
		}
	} else if compare(items[left], items[right]) < 0 {
		pivot_index = left
	} else if compare(items[middle], items[right]) < 0 {
		pivot_index = right
	}

	// Keep the pivot value even when swaps move its original item.
	var pivot T = items[pivot_index]
	// The two bounds grow the less-than and greater-than regions.
	var less_than_index int = left
	var greater_than_index int = right

	// Partition into less-than, equal-to, and greater-than pivot regions.
	for i := left; i <= greater_than_index; {
		comparison := compare(pivot, items[i])
		if comparison > 0 {
			// Move an item smaller than the pivot into the left region.
			items[less_than_index], items[i] = items[i], items[less_than_index]
			less_than_index++
			i++
		} else if comparison < 0 {
			// Move an item larger than the pivot into the right region for later inspection.
			items[greater_than_index], items[i] = items[i], items[greater_than_index]
			greater_than_index--
		} else {
			// Equal items already belong in the middle region.
			i++
		}
	}

	// Sort the lower region before the greater region; equal items need no recursion.
	if !recursion(items, left, less_than_index-1, compare) {
		return false
	}

	if !recursion(items, greater_than_index+1, right, compare) {
		return false
	}

	return true
}

// QuickSort orders items in place according to compare.
func QuickSort[T any](items []T, compare func(T, T) int) (bool, error) {
	// Reject a missing comparator before inspecting or changing the input.
	if compare == nil {
		return false, ErrNilComparator
	}

	// Sort the complete slice, including the empty-slice range of 0 through -1.
	return recursion(items, 0, len(items)-1, compare), nil
}
