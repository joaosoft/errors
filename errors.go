package errors

import (
	"fmt"
)

type IErr interface {
	Add(newErr *Err)
	Error() string
	Cause() string

	SetError(newE *Err)
	GetError() *Err
	GetPrevious() *Err
	GetErrors() []*Err

	SetCode(code string)
	GetCode() string
}

func (e *Err) Add(newErr *Err) {
	prevErr := &Err{
		previous: e.previous,
		error:    e.error,
	}

	e.previous = prevErr
	e.error = newErr
}

func (e *Err) Error() string {
	return e.error.Error()
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

func (e *Err) SetError(newE *Err) {
	*e = *newE
}

func (e *Err) GetError() *Err {
	return e
}

func (e *Err) GetPrevious() *Err {
	return e.previous
}

func (e *Err) GetErrors() []*Err {
	errors := make([]*Err, 0)
	errors = append(errors, e)

	nextErr := e.previous
	for nextErr != nil {
		errors = append(errors, e.previous)
		nextErr = nextErr.previous
	}

	return errors
}

func (e *Err) SetCode(code string) {
	e.code = code

}
func (e *Err) GetCode() string {
	return e.code
}

func (e *Err) Format(values ...interface{}) *Err {
	e.SetError(New(e.code, fmt.Sprintf(e.Error(), values)))
	return e
}
