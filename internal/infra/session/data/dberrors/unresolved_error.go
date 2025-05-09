package dberrors

import "fmt"

type UnresolvedError struct {
	message string
}

func NewUnresolvedError(message string) *UnresolvedError {
	return &UnresolvedError{message: message}
}

func (e *UnresolvedError) Error() string {
	return fmt.Sprintf("unresolved database error: %s", e.message)
}
