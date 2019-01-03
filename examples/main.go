package main

import (
	"fmt"

	".."
)

var (
	ErrorOne = errors.New(errors.ErrorLevel, 1, "Error one")
	ErrorTwo = errors.New(errors.ErrorLevel, 2, "Error two")
)

func main() {
	fmt.Println("\nADDING ERRORS!\n")

	errs := errors.Add(ErrorOne).
		Add(ErrorTwo)

	fmt.Println(errs.Cause())

	fmt.Println(errs.Stack)
	fmt.Println(errs.Previous.Stack)

	fmt.Println("\nDONE!")
}
