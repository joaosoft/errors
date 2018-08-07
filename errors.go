package errors

import (
	"errors"
	"fmt"
)

type ListErr []*Err

type Err struct {
	previous *Err
	error    error
	code     string
}

func New(code string, err interface{}) *Err {

	switch v := err.(type) {
	case error:
		return &Err{code: code, error: v}

	case string:
		return &Err{code: code, error: errors.New(v)}

	default:
		return &Err{code: code, error: errors.New(fmt.Sprint(v))}

	}
}

func (e *Err) Add(newErr *Err) {
	prevErr := &Err{
		previous: e.previous,
		error:    e.error,
	}

	e.previous = prevErr
	e.error = newErr
}

func (e *Err) Err() error {
	return e.error
}

func (e *Err) Code() string {
	return e.code
}

func (e *Err) Error() string {
	return e.error.Error()
}

func (e *Err) Errors() []error {
	errors := make([]error, 0)
	errors = append(errors, e.error)

	nextErr := e.previous
	for nextErr != nil {
		errors = append(errors, e.previous.error)
		nextErr = nextErr.previous
	}

	return errors
}

func (e *Err) Cause() string {
	str := fmt.Sprintf("'%s'", e.error.Error())

	nextErr := e.previous
	for nextErr != nil {
		str += fmt.Sprintf(", caused by '%s'", e.previous.error.Error())
		nextErr = nextErr.previous
	}
	return str
}
