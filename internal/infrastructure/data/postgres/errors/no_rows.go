package errors

type NoRowsError struct {
	message string
}

func NewNoRowsError(msg string) *NoRowsError {
	return &NoRowsError{
		message: msg,
	}
}

func (e *NoRowsError) Error() string {
	return e.message
}
