package myslice

import "fmt"

func InsertAtDemo1() {
	slice := []int{1, 2, 4, 5}
	index := 2
	value := 3
	fmt.Printf("InsertAtDemo1: before insert: %v\n", slice)
	slice = InsertAt(slice, index, value)
	fmt.Printf("InsertAtDemo1: after insert: %v\n", slice)
}

func InsertAtDemo2() {
	slice := []int{1, 2, 4, 5}
	index := 4
	value := 9
	fmt.Printf("InsertAtDemo2: before insert: %v\n", slice)
	slice = InsertAt(slice, index, value)
	fmt.Printf("InsertAtDemo2: after insert: %v\n", slice)
}

// InsertAt inserts a value into the slice at the given index and returns the updated slice.
// Example:
//     slice := []int{1, 2, 4, 5}
//     slice = InsertAt(slice, 2, 3)
//     Result: []int{1, 2, 3, 4, 5}
func InsertAt(slice []int, index int, value int) []int {
	fmt.Printf("InsertAt: before insert: %v\n", slice)
	if index < 0 || index > len(slice) {
		panic("InsertAt: index out of range")
	}
	slice = append(slice, 0) // increase the slice length by 1
	fmt.Printf("InsertAt: after extend: %v\n", slice)
	copy(slice[index+1:], slice[index:len(slice)-1]) // shift right
	fmt.Printf("InsertAt: after shift: %v\n", slice)
	slice[index] = value
	fmt.Printf("InsertAt: after insert: %v\n", slice)
	return slice
}
