package main

import (
	"fmt"
	"go-error/app"
)

func main() {

	err := goerror.NewError(fmt.Errorf("erro 1"))
	err.AddError(fmt.Errorf("erro 2"))
	err.AddError(fmt.Errorf("erro 3"))

	fmt.Printf("Error: %s, Cause: %s", err.Error(), err.Cause())
}
