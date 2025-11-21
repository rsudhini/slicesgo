package myslice

import "fmt"

// For small slices (<1024 cap) → capacity generally doubles
// For large slices (≥1024 cap) → capacity grows by about +25%
// Go may grow differently depending on: element size, append pattern, and Go version
func SliceCapacityDemo() {
	// Empty slice created with zero length and zero capacity
	numSlice := []int{}
	for i := 0; i < 20; i++ {
		// Capacity increases automatically when appending beyond current capacity.
		numSlice = append(numSlice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(numSlice), cap(numSlice))
	}
}
