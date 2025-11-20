package myslice

import "fmt"

func CreateSliceDemo1() {
	fmt.Println("Creating slice...")
	slice1 := []string{}
	// here appending a string not bytes
	slice1 = append(slice1, "Hello")
	fmt.Println("length of slice: ", len(slice1))
	fmt.Println("capacity of slice: ", cap(slice1))

}

func CreateSliceDemo2() {
	fmt.Println("Creating slice...")
	slice2 := []string{}
	fmt.Println("length of slice: ", len(slice2))
	fmt.Println("capacity of slice: ", cap(slice2))
}

func CreateSliceDemo3() {
	// Case 3 using make function
	// slice3 := make([]string, 3, 3)
	slice3 := make([]string, 3) // succinct version, length and capacity are same
	slice3[0] = "Hello"
	// slice3[4] = "Hi" // This will give runtime error
	fmt.Println("length of slice3: ", len(slice3))
	fmt.Println("capacity of slice3: ", cap(slice3))
}
