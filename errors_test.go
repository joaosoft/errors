package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleSimple(t *testing.T) {

	err := New(InfoLevel, 1, "error 1")
	err.Add(New(ErrorLevel, 2, "error 2"))
	err.Add(New(PanicLevel, 3, "error 3"))

	fmt.Printf("Error: %s, Cause: %s", err.String(), err.Cause())

	assert.Equal(t, err.String(), `{"previous":{"previous":{"code":1,"error":"error 1"},"code":2,"error":"error 2"},"code":3,"error":"error 3"}`)
	assert.Equal(t, err.Cause(), `'error 3', caused by 'error 2', caused by 'error 1'`)
}

func TestExampleList(t *testing.T) {

	var errs ListErr
	errs.Add(New(InfoLevel, 1, "error 1"))
	errs.Add(New(ErrorLevel, 2, "error 2"))
	errs.Add(New(PanicLevel, 3, "error 3"))

	fmt.Printf("Errors: %s", errs.String())

	assert.Equal(t, errs.String(), `[{"code":1,"error":"error 1"},{"code":2,"error":"error 2"},{"code":3,"error":"error 3"}]`)
}
