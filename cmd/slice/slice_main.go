package main

// import "thirdpartlibs"

// import "usedeflibs"

import (
	"fmt"
	"os"

	"github.com/rsudhini/slicesgo/myslice"
)

func main() {

	var v int
	fmt.Printf("%d\n", v)

	var caseVar string = "case1"
	if len(os.Args) < 2 {
		fmt.Print("usage: go run cmd/app/sliceversusarray. case1")
		caseVar = "case1"
	} else {
		caseVar = os.Args[1]
	}
	fmt.Printf("%s\n", caseVar)

	switch caseVar {

	// Case 1
	case "case1":
		myslice.CreateSliceDemo1() // 1.1
		myslice.CreateSliceDemo2() // 1.2
		myslice.CreateSliceDemo3() // 1.3

		// // 	Case 2
		// case "case2":
		// 	myslice.CreateSliceFromArray()

		// // 	Case 3
		// case "case3":
		// 	myslice.SliceCapacityDemo()

		// // Case 4
		// case "case4":
		// 	numSlice := []int{1, 5, 7}
		// 	var newSlice = make([]int, 10, 20)
		// 	fmt.Printf("new slice before: %v\n", newSlice)
		// 	myslice.CopySlice(newSlice, numSlice)
		// 	fmt.Printf("new slice after: %v\n", newSlice)

		// // 	Case 5
		// case "case5":
		// 	dataSlice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		// 	// dataSlice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		// 	// dataSlice3 := []int{11, 91, 42, 56, 76, 84, 11, 23, 31, 9}
		// 	myslice.ReverseSliceDemo(dataSlice1)

		// // 	Case 6
		// case "case6":
		// 	myslice.ReverseSliceWithAdditionalSlice()
	}
}
