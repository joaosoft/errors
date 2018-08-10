package errors

func (e *ListErr) Len() int {
	return len(*e)
}

func (e *ListErr) IsEmpty() bool {
	return len(*e) == 0
}

func (e *ListErr) Add(err IErr) *ListErr {
	*e = append(*e, err)
	return e
}
