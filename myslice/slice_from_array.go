package myslice

import "fmt"

func CreateSliceFromArray() {
	// Creating an array
	arr := [7]string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}
	fmt.Printf("Array: %v\n", arr)

	// Creating a slice from an array
	length := len(arr)
	slice := arr[0:length]

	// Display slice, length of slice and capacity of slice
	fmt.Printf("Slice: %v\n", slice)

}
