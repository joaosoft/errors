package main

import (
	"fmt"

	errors ".."
)

var (
	ErrorOne   = errors.New(errors.LevelError, 1, "error one")
	ErrorTwo   = errors.New(errors.LevelError, 2, "error two")
	ErrorThree = errors.New(errors.LevelError, 3, "error three")
)

func main() {
	exampleSimple()
	exampleList()
	exampleError()
}

func exampleSimple() {

	err := errors.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	fmt.Printf("\nerror: %s, Cause: %s", err.String(), err.Cause())
}

func exampleList() {

	var errs errors.ErrorList
	errs.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	fmt.Printf("\n\nErrors: %s", errs.String())
	fmt.Printf("\n\nErrors: %s", ErrorOne.Error())
}

func exampleError() {

	var err error
	err = errors.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	fmt.Printf("\nerr: %s, Cause: %s", err, err)
}
