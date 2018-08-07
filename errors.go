package errors

import (
	"errors"
	"fmt"
)

type ListErr []*Err

type Err struct {
	Previous *Err   `json:"previous"`
	Err      error  `json:"error"`
	Code     string `json:"code"`
}

func New(code string, err interface{}) *Err {

	switch v := err.(type) {
	case error:
		return &Err{Code: code, Err: v}

	case string:
		return &Err{Code: code, Err: errors.New(v)}

	default:
		return &Err{Code: code, Err: errors.New(fmt.Sprint(v))}

	}
}

func (e *Err) Add(newErr *Err) {
	prevErr := &Err{
		Previous: e.Previous,
		Err:      e.Err,
	}

	e.Previous = prevErr
	e.Err = newErr
}

func (e *Err) Error() string {
	return e.Err.Error()
}

func (e *Err) Errors() []*Err {
	errors := make([]*Err, 0)
	errors = append(errors, e)

	nextErr := e.Previous
	for nextErr != nil {
		errors = append(errors, e.Previous)
		nextErr = nextErr.Previous
	}

	return errors
}

func (e *Err) Cause() string {
	str := fmt.Sprintf("'%s'", e.Err.Error())

	nextErr := e.Previous
	for nextErr != nil {
		str += fmt.Sprintf(", caused by '%s'", e.Previous.Err.Error())
		nextErr = nextErr.Previous
	}
	return str
}
