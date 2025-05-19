package errors

func NewNotFoundError(msg string, cause error) *AppError {
	return NewAppError(NotFound, msg, cause, nil)
}
