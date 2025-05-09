package apperrors

import (
	"fmt"

	"github.com/EugeneTsydenov/chesshub-sessions-service/pkg/validationerrors"
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

type AppError struct {
	Code    ErrorCode
	Message string
	Details validationerrors.ErrorDetails
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%v: %s", e.Code, e.Message)
}
