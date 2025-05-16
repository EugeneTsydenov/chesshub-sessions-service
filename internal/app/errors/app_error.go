package errors

import (
	"errors"
	"fmt"
)

type ErrorCode int

const (
	InvalidArgument ErrorCode = iota
	NotFound
	Conflict
	Internal
	Unauthenticated
	Forbidden
	Canceled
	DeadlineExceeded
)

func (c ErrorCode) String() string {
	switch c {
	case InvalidArgument:
		return "INVALID_ARGUMENT_ERROR"
	case NotFound:
		return "NOT_FOUND_ERROR"
	case Conflict:
		return "CONFLICT_ERROR"
	case Internal:
		return "INTERNAL_ERROR"
	case Unauthenticated:
		return "UNAUTHENTICATED_ERROR"
	case Forbidden:
		return "FORBIDDEN_ERROR"
	case DeadlineExceeded:
		return "DEADLINE_EXCEEDED_ERROR"
	case Canceled:
		return "CANCELED_ERROR"
	default:
		return "UNKNOWN_ERROR"
	}
}

type ErrorDetails map[string]string

type AppError struct {
	Code    ErrorCode
	Message string
	Cause   error
	Details ErrorDetails
}

func NewAppError(code ErrorCode, msg string, cause error, details ErrorDetails) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
		Cause:   cause,
		Details: details,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%v: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func (e *AppError) Join() error {
	return errors.Join(e, e.Cause)
}
