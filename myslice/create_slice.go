package myslice

import "fmt"

// CreateEmptySliceDemo creates empty slice using string literal with zero length and default capacity
func CreateEmptySliceDemo() {
	slice := []string{}
	fmt.Println("CreateEmptySliceDemo - slice: ", slice)
	fmt.Println("CreateEmptySliceDemo - length of slice: ", len(slice))
	fmt.Println("CreateEmptySliceDemo - capacity of slice: ", cap(slice))
}

// CreateEmptySliceUsingMakeDemo creates an empty slice using the make function with zero length.
func CreateEmptySliceUsingMakeDemo() {
	slice := make([]string, 0)
	fmt.Println("CreateEmptySliceUsingMakeDemo - slice: ", slice)
	fmt.Println("CreateEmptySliceUsingMakeDemo - length of slice: ", len(slice))
	fmt.Println("CreateEmptySliceUsingMakeDemo - capacity of slice: ", cap(slice))
}

// CreateSliceWithSameLenAndCapDemo creates slice using make function, slice length and capacity are same
func CreateSliceWithSameLenCapDemo() {
	fmt.Println("Creating slice demo4")
	slice := make([]string, 3)
	slice[0] = "Hello"
	fmt.Println("CreateSliceWithSameLenAndCapDemo - slice: ", slice)
	fmt.Println("CreateSliceWithSameLenAndCapDemo - length of slice: ", len(slice))
	fmt.Println("CreateSliceWithSameLenAndCapDemo - capacity of slice4: ", cap(slice))
}

// CreateSliceWithDiffLenCapDemo creates a slice with different length and capacity
func CreateSliceWithDiffLenCapDemo() {
	slice := make([]int, 3, 6)
	fmt.Printf("CreateSliceWithDiffLenCapDemo - slice: %v\n", slice)
	fmt.Printf("CreateSliceWithDiffLenCapDemo - slice length: %d\n", len(slice))
	fmt.Printf("CreateSliceWithDiffLenCapDemo - slice capacity: %d\n", cap(slice))
}

// CreateSliceExplDeclDemo demonstrates explicit slice declaration using both
// the left-hand side (var keyword) and the right-hand side (slice literal).
func CreateSliceExplDeclDemo() {
	var slice []int = []int{} // explicit declaration, no shorthand
	fmt.Printf("CreateSliceExplDeclDemo - slice: %v\n", slice)
	fmt.Printf("CreateSliceExplDeclDemo - slice length: %v\n", len(slice))
	fmt.Printf("CreateSliceExplDeclDemo - slice capacity: %v\n", cap(slice))
}

// CreateSliceWithInitValuesDemo creates a slice using a slice literal with initialized values.
func CreateSliceWithInitValuesDemo() {
	slice := []int{1, 2, 3}
	fmt.Printf("CreateSliceWithInitValuesDemo - slice: %v\n", slice)
	fmt.Printf("CreateSliceWithInitValuesDemo - slice length: %v\n", len(slice))
	fmt.Printf("CreateSliceWithInitValuesDemo - slice capacity: %v\n", cap(slice))
}

// CreateNilSliceDemo creates a nil slice.
func CreateNilSliceDemo() {
	var slice []int = CreateNilSlice()
	fmt.Printf("CreateNilSliceDemo - slice: %v\n", slice)
	fmt.Printf("CreateNilSliceDemo - slice length: %d\n", len(slice))
	fmt.Printf("CreateNilSliceDemo - slice capacity: %d\n", cap(slice))

	if slice == nil {
		fmt.Println("CreateNilSliceDemo - slice is nil")
	}
}

// CreateNilSlice is a helper func for CreateNilSliceDemo
func CreateNilSlice() []int {
	var slice []int
	return slice
}
