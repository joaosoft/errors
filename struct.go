package errors

type ListErr []*Err

type Err struct {
	previous *Err   `json:"previous"`
	error    error  `json:"error"`
	code     string `json:"code"`
}
