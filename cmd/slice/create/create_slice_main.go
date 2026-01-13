package main

// import "thirdpartlibs"

// import "usedeflibs"

import (
	"fmt"
	"os"

	"github.com/rsudhini/slicesgo/myslice"
)

const (
	cmdEmptySlice          = "empty-slice"
	cmdEmptySliceUsingMake = "empty-slice-using-make"
	cmdSliceWithSameLenCap = "slice-with-same-len-cap"
	cmdSliceWithDiffLenCap = "slice-with-diff-len-cap"
	cmdSliceExplDecl       = "slice-with-expl-decl"
	cmdSliceWithInitValues = "slice-with-init-values"
	cmdNilSlice            = "nil-slice"
	cmdSliceFromArray      = "slice-from-array"
	cmdSliceFromArrayFull  = "slice-from-array-full"
	cmdHelp                = "help"
)

func main() {

	fmt.Print("Inside slice main\n")

	var caseVar string

	if len(os.Args) < 2 {
		caseVar = "empty-slice"
	} else {
		caseVar = os.Args[1]
	}
	fmt.Printf("%s\n", caseVar)

	switch caseVar {

	case cmdEmptySlice:
		myslice.CreateEmptySliceDemo()
	case cmdEmptySliceUsingMake:
		myslice.CreateEmptySliceUsingMakeDemo()
	case cmdSliceWithSameLenCap:
		myslice.CreateSliceWithSameLenCapDemo()
	case cmdSliceWithDiffLenCap:
		myslice.CreateSliceWithDiffLenCapDemo()
	case cmdSliceExplDecl:
		myslice.CreateSliceExplDeclDemo()
	case cmdSliceWithInitValues:
		myslice.CreateSliceWithInitValuesDemo()
	case cmdNilSlice:
		myslice.CreateNilSliceDemo()
	case cmdHelp:
		printHelp()
	case cmdSliceFromArray:
		myslice.CreateSliceFromArrayDemo()
	case cmdSliceFromArrayFull:
		myslice.CreateSliceFromArrayFullDemo()
	default:
		fmt.Printf("Unknown command: %s\n", caseVar)
		printHelp()

	// Case 3
	case "case3":
		myslice.SliceCapacityDemo1()

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

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run ./cmd/slice/create <command>")
	fmt.Println()

	fmt.Println("Commands:")
	fmt.Printf("  %-25s %s\n", cmdEmptySlice, "create an empty slice using slice literal")
	fmt.Printf("  %-25s %s\n", cmdEmptySliceUsingMake, "create an empty slice using make")
	fmt.Printf("  %-25s %s\n", cmdNilSlice, "create a nil slice")
	fmt.Printf("  %-25s %s\n", cmdSliceWithSameLenCap, "create slice when len and cap are same")
	fmt.Printf("  %-25s %s\n", cmdSliceWithDiffLenCap, "create slice when len and cap are different")
	fmt.Printf("  %-25s %s\n", cmdSliceExplDecl, "create slice using explicit declaration")
	fmt.Printf("  %-25s %s\n", cmdSliceWithInitValues, "create slice with initialized values")
	fmt.Printf("  %-25s %s\n", cmdSliceFromArray, "create a slice from an array using range")
	fmt.Printf("  %-25s %s\n", cmdSliceFromArrayFull, "create a full slice from array")
	fmt.Printf("  %-25s %s\n", cmdHelp, "print this help message")
	fmt.Println()

	fmt.Println("Examples:")
	fmt.Printf("  go run ./cmd/slice/create %s\n", cmdEmptySlice)
	fmt.Printf("  go run ./cmd/slice/create %s\n", cmdSliceFromArrayFull)
	fmt.Printf("  go run ./cmd/slice/create %s\n", cmdHelp)
}
