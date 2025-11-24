package main

// import "thirdpartlibs"

// import "usedeflibs"

import (
	"fmt"
	"os"

	"github.com/rsudhini/slicesgo/myslice"
)

func main() {

	fmt.Print("Inside slice main\n")

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
		myslice.CreateSliceDemo4() // 1.4

	// 	Case 2
	case "case2":
		myslice.CreateSliceFromArrayDemo1() // 2.1
		myslice.CreateSliceFromArrayDemo2() // 2.2

	// Case 3
	case "case3":
		myslice.SliceCapacityDemo()

	// Case 4 Copying slice values from one slice to another slice
	case "case4":
		myslice.CopySliceDemo()

	// Case 5
	case "case5":
		myslice.ReverseSliceDemo1() // 5.1
		myslice.ReverseSliceDemo2() // 5.2
		myslice.ReverseSliceDemo3() // 5.3

		//  Case 6
	case "case6":
		myslice.ReverseSliceWithAdditionalSliceDemo()
	}
}
