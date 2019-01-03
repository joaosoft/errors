package errors

import (
	"encoding/json"
	"fmt"
)

func (e *Err) Add(newErr *Err) {
	prevErr := &Err{
		Previous: e.Previous,
		Code:     e.Code,
		Message:  e.Message,
	}

	e.Previous = prevErr
	e.Code = newErr.Code
	e.Message = newErr.Message
}

func (e *Err) Error() string {
	return e.Message
}

func (e *Err) Cause() string {
	str := fmt.Sprintf("'%s'", e.Message)

	prevErr := e.Previous
	for prevErr != nil {
		str += fmt.Sprintf(", caused by '%s'", prevErr.Message)
		prevErr = prevErr.Previous
	}
	return str
}

func (e *Err) SetError(newErr *Err) {
	*e = *newErr
}

func (e *Err) GetError() *Err {
	return e
}

func (e *Err) GetPrevious() *Err {
	return e.Previous
}

func (e *Err) GetErrors() []*Err {
	errors := make([]*Err, 0)
	errors = append(errors, e)

	nextErr := e.Previous
	for nextErr != nil {
		errors = append(errors, e.Previous)
		nextErr = nextErr.Previous
	}

	return errors
}

func (e *Err) SetCode(code string) {
	e.Code = code

}
func (e *Err) GetCode() string {
	return e.Code
}

func (e *Err) Format(values ...interface{}) *Err {
	e.SetError(New(e.Code, fmt.Sprintf(e.Error(), values)))
	return e
}

func (e *Err) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}
