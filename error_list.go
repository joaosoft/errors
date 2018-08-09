package errors

func (e *ListErr) Len() int {
	return len(*e)
}

func (e *ListErr) IsEmpty() bool {
	return len(*e) == 0
}