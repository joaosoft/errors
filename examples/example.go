package examples

import (
	"fmt"

	"../../errors"
)

func main() {

	err := errors.New("0", "erro 1")
	err.Add(errors.New("0", "erro 2"))
	err.Add(errors.New("0", "erro 3"))

	fmt.Printf("Error: %s, Cause: %s", err.Error(), err.Cause())
}
