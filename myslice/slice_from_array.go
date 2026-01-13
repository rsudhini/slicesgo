package myslice

import "fmt"

func CreateSliceFromArrayDemo() {
	// Creating an array
	arr := [7]string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}
	fmt.Printf("Array: %v\n", arr)

	// Creating a slice from an array
	length := len(arr)
	slice := arr[0:length]

	// Display slice, length of slice and capacity of slice
	fmt.Printf(" Demo1 Slice: %v\n", slice)

}

func CreateSliceFromArrayFullDemo() {
	// Creating an array
	arr := [7]string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}
	fmt.Printf("Array: %v\n", arr)

	// Creating a slice from an array, idiomatic way
	slice := arr[:]

	// Display slice, length of slice and capacity of slice
	fmt.Printf("Demo2 Slice: %v\n", slice)

}
