package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleSimple(t *testing.T) {

	err := New("1", "erro 1")
	err.Add(New("2", "erro 2"))
	err.Add(New("3", "erro 3"))

	fmt.Printf("Error: %s, Cause: %s", err.String(), err.Cause())

	assert.Equal(t, err.String(), `{"previous":{"previous":{"code":"1","error":"erro 1"},"code":"2","error":"erro 2"},"code":"3","error":"erro 3"}`)
	assert.Equal(t, err.Cause(), `'erro 3', caused by 'erro 2', caused by 'erro 1'`)
}

func TestExampleList(t *testing.T) {

	var errs ListErr
	errs.Add(New("1", "erro 1"))
	errs.Add(New("2", "erro 2"))
	errs.Add(New("3", "erro 3"))

	fmt.Printf("Errors: %s", errs.String())

	assert.Equal(t, errs.String(), `[{"code":"1","error":"erro 1"},{"code":"2","error":"erro 2"},{"code":"3","error":"erro 3"}]`)
}
