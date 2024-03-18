package ignoreerror

type IgnorableError struct {
	Msg string
}

// Error implements the error interface.
func (e *IgnorableError) Error() string {
	return e.Msg
}
