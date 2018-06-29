package errors

import (
	"fmt"
)

func main() {

	err := NewError(fmt.Errorf("erro 1"))
	err.AddError(fmt.Errorf("erro 2"))
	err.AddError(fmt.Errorf("erro 3"))

	fmt.Printf("Error: %s, Cause: %s", err.Error(), err.Cause())
}
