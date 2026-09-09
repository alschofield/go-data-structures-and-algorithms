package radix_sort

func RadixCount(items []uint32, shift uint) ([]uint32, bool) {
	// One output slot per value keeps this byte pass stable.
	sorted_items := make([]uint32, len(items))
	// A byte has 256 possible values, independent of the full uint32 range.
	keys := make([]uint, 256)
	for i := 0; i < len(items); i++ {
		// Count the byte selected by this pass.
		keys[(items[i]>>shift)&0xff]++
	}

	// Convert each frequency into that bucket's next output position.
	var total uint = 0
	for i := 0; i < len(keys); i++ {
		count := keys[i]
		keys[i] = total
		total += count
	}

	// Move complete values left to right so equal byte keys keep their order.
	for i := 0; i < len(items); i++ {
		key := (items[i] >> shift) & 0xff
		sorted_items[keys[key]] = items[i]
		keys[key]++
	}

	return sorted_items, true
}

func RadixSort(items []uint32) bool {
	// A nil or short slice is already sorted.
	if items == nil {
		return true
	}

	if len(items) == 0 || len(items) == 1 {
		return true
	}

	// Each pass reads one byte and writes a stable ordering for the next pass.
	sorted_items := items
	var status bool = false
	for shift := 0; shift < 32; shift += 8 {
		sorted_items, status = RadixCount(sorted_items, uint(shift))
		if !status {
			return false
		}
	}

	// The last pass owns a new buffer; copy it into the caller's slice.
	copy(items, sorted_items)

	return true
}
