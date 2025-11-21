package myslice

import "fmt"

func CopySliceDemo() {
	slice := []int{1, 5, 7}
	fmt.Printf("slice: %v\n", slice)
	newSlice := make([]int, 10, 20)
	fmt.Printf("new slice before: %v\n", newSlice)
	// Values from slice are copied into newSlice, overwriting newSlice's elements up to len(slice) elements.
	copy(newSlice, slice)
	fmt.Printf("new slice after: %v\n", newSlice)
}
