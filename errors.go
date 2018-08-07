package errors

import (
	"errors"
	"fmt"
)

type ErrorData struct {
	previous *ErrorData
	error    error
	code     string
}

func New(err interface{}) *ErrorData {

	switch v := err.(type) {
	case error:
		return &ErrorData{error: v}

	case string:
		return &ErrorData{error: errors.New(v)}

	default:
		return &ErrorData{error: errors.New(fmt.Sprint(v))}

	}
}

func (e *ErrorData) Add(newErr interface{}) {
	prevErr := &ErrorData{
		previous: e.previous,
		error:    e.error,
	}
	e.previous = prevErr

	switch v := newErr.(type) {
	case error:
		e.error = v

	case string:
		e.error = errors.New(v)

	default:
		e.error = errors.New(fmt.Sprint(v))

	}
}

func (e *ErrorData) Err() error {
	return e.error
}

func (e *ErrorData) Code() string {
	return e.code
}

func (e *ErrorData) Error() string {
	return e.error.Error()
}

func (e *ErrorData) Errors() []error {
	errors := make([]error, 0)
	errors = append(errors, e.error)

	nextErr := e.previous
	for nextErr != nil {
		errors = append(errors, e.previous.error)
		nextErr = nextErr.previous
	}

	return errors
}

func (e *ErrorData) Cause() string {
	str := fmt.Sprintf("'%s'", e.error.Error())

	nextErr := e.previous
	for nextErr != nil {
		str += fmt.Sprintf(", caused by '%s'", e.previous.error.Error())
		nextErr = nextErr.previous
	}
	return str
}
