package myslice

import "fmt"

// Reverses the slice iteratively by swapping left-side element with right-side element
// Any changes to the slice inside this function will be reflected in the caller's slice.
// Slices in Go are reference types. Passing a slice to a function does not copy the underlying array.
func reverseSlice(dataSlice []int) {
	length := len(dataSlice)
	fmt.Printf("length: %v\n", length)
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		// Swapping array elements dataSlice[i], dataSlice[j]
		dataSlice[i], dataSlice[j] = dataSlice[j], dataSlice[i]
	}
}

func ReverseSliceDemo1() {
	dataSlice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("dataSlice1 before: %v\n", dataSlice1)
	reverseSlice(dataSlice1)
	fmt.Printf("dataSlice1 after: %v\n", dataSlice1)
}

func ReverseSliceDemo2() {
	dataSlice2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Printf("dataSlice2 before: %v\n", dataSlice2)
	reverseSlice(dataSlice2)
	fmt.Printf("dataSlice2 after: %v\n", dataSlice2)
}

func ReverseSliceDemo3() {
	dataSlice3 := []int{11, 91, 42, 56, 76, 84, 11, 23, 31, 9}
	fmt.Printf("dataSlice3 before: %v\n", dataSlice3)
	reverseSlice(dataSlice3)
	fmt.Printf("dataSlice3 after: %v\n", dataSlice3)
}
