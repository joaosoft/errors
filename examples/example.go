package examples

import (
	"fmt"

	"../../errors"
)

func main() {

	err := errors.New(fmt.Errorf("erro 1"))
	err.Add(fmt.Errorf("erro 2"))
	err.Add(fmt.Errorf("erro 3"))

	fmt.Printf("Error: %s, Cause: %s", err.Error(), err.Cause())
}
