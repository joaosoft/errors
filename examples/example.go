package examples

import (
	"fmt"
	"../../errors"
)

func main() {

	err := errors.NewError(fmt.Errorf("erro 1"))
	err.AddError(fmt.Errorf("erro 2"))
	err.AddError(fmt.Errorf("erro 3"))

	fmt.Printf("Error: %s, Cause: %s", err.Error(), err.Cause())
}
