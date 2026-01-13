package myslice

import "fmt"

// SliceCapacityDemo1 creates slice with same length and capacity
func SliceCapacityDemo1() {
	var slicelmn = make([]int, 3)
	fmt.Printf("slice3: %v\n", slicelmn)
	fmt.Printf("slice3 length: %v\n", len(slicelmn))
	fmt.Printf("slice3 capacity: %v\n", cap(slicelmn))
}

// SliceCapacityDemo2 creates a slice with different length and capacity
func SliceCapacityDemo2() {
	slice := make([]int, 3, 6)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("slice length: %d\n", len(slice))
	fmt.Printf("slice capacity: %d\n", cap(slice))
}

// SliceCapacityDemo2 appends value to slice with no capacity remaining, this cause Go to
// create a new slice with capacity increase by double and length increase by one
func SliceCapacityDemo3() {
	var slice = make([]int, 3)
	slice2 := append(slice, 4)
	fmt.Printf("SliceCapacityDemo1: slice2: %v\n", slice2)
	fmt.Printf("SliceCapacityDemo1: slice2 length: %v\n", len(slice2))
	fmt.Printf("SliceCapacityDemo1: slice2 capacity: %v\n", cap(slice2))
}

// For small slices (<1024 cap) → capacity generally doubles
// For large slices (≥1024 cap) → capacity grows by about +25%
// Go may grow differently depending on: element size, append pattern, and Go version
func SliceCapacityDemo4() {
	// Empty slice created with zero length and zero capacity
	numSlice := []int{}
	for i := 0; i < 20; i++ {
		// Capacity increases automatically when appending beyond current capacity.
		numSlice = append(numSlice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(numSlice), cap(numSlice))
	}
}
