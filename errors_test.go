package errors

import (
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

	assert.Equal(t, err.Message, `Error three`)
	assert.Equal(t, err.Cause(), `'Error three', caused by 'Error two', caused by 'Error one'`)
}

func TestExampleList(t *testing.T) {

	var errs ListErr
	errs.Add(ErrorOne).Add(ErrorTwo).Add(ErrorThree)

	assert.Equal(t, errs.Len(), 3)
}
