package errors

import "errors"

func NewInternalError(msg string, cause error) *AppError {
	return NewAppError(Internal, msg, cause, nil)
}

func IsInternalError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == Internal
}
