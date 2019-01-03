package errors

type ListErr []*Err

type Err struct {
	Previous *Err   `json:"previous,omitempty"`
	Code     string `json:"code"`
	Message  string `json:"error"`
}
