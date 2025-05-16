package errors

import "errors"

func NewCanceledError(msg string, cause error) *AppError {
	return NewAppError(Canceled, msg, cause, nil)
}

func IsCanceledError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == Canceled
}
