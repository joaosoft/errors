package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ErrorOne   = New(ErrorLevel, 1, "Error one")
	ErrorTwo   = New(ErrorLevel, 2, "Error two")
	ErrorThree = New(ErrorLevel, 3, "Error three")
)

func TestExampleSimple(t *testing.T) {

	err := Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	fmt.Printf("Error: %s, Cause: %s", err.String(), err.Cause())

	assert.Equal(t, err.String(), `Error: {"previous":{"previous":{"level":2,"code":1,"message":"Error one","stack":"main.exampleSimple()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:22 +0x3f\nmain.main()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:16 +0x20\n"},"level":2,"code":2,"message":"Error two","stack":"main.exampleSimple()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:22 +0x59\nmain.main()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:16 +0x20\n"},"level":2,"code":3,"message":"Error three","stack":"main.exampleSimple()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:22 +0x73\nmain.main()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:16 +0x20\n"}`)
	assert.Equal(t, err.Cause(), `'Error three', caused by 'Error two', caused by 'Error one''`)
}

func TestExampleList(t *testing.T) {

	var errs ListErr
	errs.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	fmt.Printf("Errors: %s", errs.String())

	assert.Equal(t, errs.String(), `[{"level":2,"code":1,"message":"Error one","stack":"main.init()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:10 +0x99\n"},{"level":2,"code":2,"message":"Error two","stack":"main.init()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:11 +0xfc\n"},{"level":2,"code":3,"message":"Error three","stack":"main.init()\n\t/Users/joaoribeiro/workspace/go/personal/src/errors/examples/main.go:12 +0x15f\n"}]`)
}
