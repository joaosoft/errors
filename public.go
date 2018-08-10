package errors

import (
	"errors"
	"fmt"
)

func New(code string, err interface{}, params ...interface{}) *Err {

	switch v := err.(type) {
	case error:
		return &Err{code: code, error: v}

	case string:
		return &Err{code: code, error: errors.New(fmt.Sprintf(v, params))}

	default:
		return &Err{code: code, error: errors.New(fmt.Sprint(v))}

	}
}
