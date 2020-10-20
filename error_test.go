package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ErrorOne   = New(LevelError, 1, "error one")
	ErrorTwo   = New(LevelError, 2, "error two")
	ErrorThree = New(LevelError, 3, "error three")
)

func TestExampleSimple(t *testing.T) {

	err := Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	assert.Equal(t, err.Message, `error three`)
	assert.Equal(t, err.Cause(), `'error three', caused by 'error two', caused by 'error one'`)
}

func TestExampleList(t *testing.T) {

	var errs ErrorList
	errs.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	assert.Equal(t, errs.Len(), 3)
}
