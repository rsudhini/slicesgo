package myslice

import "fmt"

func CreateSliceDemo1() {
	fmt.Println("Creating slice demo1")
	// Empty slice is created with zero length and zero capacity
	slice1 := []string{}
	fmt.Println("slice1: ", slice1)
	fmt.Println("length of slice: ", len(slice1))
	fmt.Println("capacity of slice: ", cap(slice1))
}

func CreateSliceDemo2() {
	fmt.Println("Creating slice demo2")
	slice2 := []string{}
	// Appending a string so length and capacity increased by 1
	slice2 = append(slice2, "Hello")
	fmt.Println("slice2: ", slice2)
	fmt.Println("length of slice: ", len(slice2))
	fmt.Println("capacity of slice: ", cap(slice2))

	// Appending a string so length and capacity increased by 1
	slice2 = append(slice2, "Hi")
	fmt.Println("slice2: ", slice2)
	fmt.Println("length of slice: ", len(slice2))
	fmt.Println("capacity of slice: ", cap(slice2))

	// Appending a string so length increased by 1 and capacity increased by 2 because doubles
	// once slice capacity is reached it normally doubles but this behavior can change slightly
	// for higher capacity
	slice2 = append(slice2, "Howdy")
	fmt.Println("slice2: ", slice2)
	fmt.Println("length of slice: ", len(slice2))
	fmt.Println("capacity of slice: ", cap(slice2))
}

func CreateSliceDemo3() {
	fmt.Println("Creating slice demo3")
	// slice3 := make([]string, 3, 3)
	slice3 := make([]string, 3) // succinct version, length and capacity are same
	slice3[0] = "Hello"
	// slice3[4] = "Hi" // This will give runtime error
	fmt.Println("slice3: ", slice3)
	fmt.Println("length of slice3: ", len(slice3))
	fmt.Println("capacity of slice3: ", cap(slice3))
}

func CreateSliceDemo4() {
	fmt.Println("Creating slice demo4")
	// succinct version, length and capacity are same
	slice4 := make([]string, 3)
	slice4[0] = "Hello"
	fmt.Println("slice4: ", slice4)
	fmt.Println("length of slice4: ", len(slice4))
	fmt.Println("capacity of slice4: ", cap(slice4))
}

//
