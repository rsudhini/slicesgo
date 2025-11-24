package myslice

import "fmt"

// ReverseSliceWithAdditionalSliceDemo reverses a slice by creating a new slice and
// then copying elements from the original slice into the new slice from right to left.
func ReverseSliceWithAdditionalSliceDemo() {

	slice := []int{1, 2, 5, 7}
	fmt.Printf("slice: %v\n", slice)
	// Create a new slice with the same length as the original
	reversedSlice := make([]int, len(slice))

	length := len(slice)
	// Reverse the original slice by copying elements into the new slice from right to left
	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversedSlice[j] = slice[i]
	}

	fmt.Printf("reversedSlice: %v\n", reversedSlice)
}
