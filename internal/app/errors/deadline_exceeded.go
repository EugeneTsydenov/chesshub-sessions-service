package errors

import "errors"

func NewDeadlineExceededError(msg string, cause error) *AppError {
	return NewAppError(DeadlineExceeded, msg, cause, nil)
}

func IsDeadlineExceededError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == DeadlineExceeded
}
