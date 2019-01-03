package errors

type ListErr []*Err

type Err struct {
	Previous *Err   `json:"previous,omitempty"`
	Level    Level  `json:"code"`
	Code     int    `json:"code"`
	Message  string `json:"error"`
	Stack    string `json:"stack"`
}
