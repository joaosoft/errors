package errors

import (
	"errors"
	"fmt"
)

type ErrorData struct {
	previous *ErrorData
	current  error
}

func New(err interface{}) *ErrorData {

	switch v := err.(type) {
	case error:
		return &ErrorData{current: v}

	case string:
		return &ErrorData{current: errors.New(v)}

	default:
		return &ErrorData{current: errors.New(fmt.Sprint(v))}

	}
}

func (e *ErrorData) Add(newErr interface{}) {
	prevErr := &ErrorData{
		previous: e.previous,
		current:  e.current,
	}
	e.previous = prevErr

	switch v := newErr.(type) {
	case error:
		e.current = v

	case string:
		e.current = errors.New(v)

	default:
		e.current = errors.New(fmt.Sprint(v))

	}
}

func (e *ErrorData) Error() string {
	return e.current.Error()
}

func (e *ErrorData) Cause() string {
	str := fmt.Sprintf("'%s'", e.current.Error())
	nextErr := e.previous

	for nextErr != nil {
		str += fmt.Sprintf(", caused by '%s'", e.previous.current.Error())
		nextErr = nextErr.previous
	}
	return str
}
