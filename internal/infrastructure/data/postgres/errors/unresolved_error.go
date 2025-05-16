package errors

import "fmt"

type UnresolvedError struct {
	message string
	cause   error
}

func NewUnresolvedError(message string, cause error) *UnresolvedError {
	return &UnresolvedError{
		message: message,
		cause:   cause,
	}
}

func (e *UnresolvedError) Error() string {
	return fmt.Sprintf("unresolved postgres error: %s: %v", e.message, e.cause)
}

func (e *UnresolvedError) Unwrap() error {
	return e.cause
}
