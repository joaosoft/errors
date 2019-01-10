package main

import (
	"fmt"

	".."
)

var (
	ErrorOne   = errors.New(errors.ErrorLevel, 1, "Error one")
	ErrorTwo   = errors.New(errors.ErrorLevel, 2, "Error two")
	ErrorThree = errors.New(errors.ErrorLevel, 3, "Error three")
)

func main() {
	exampleSimple()
	exampleList()
}

func exampleSimple() {

	err := errors.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	fmt.Printf("\nError: %s, Cause: %s", err.String(), err.Cause())
}

func exampleList() {

	var errs errors.ListErr
	errs.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	fmt.Printf("\n\nErrors: %s", errs.String())
}
