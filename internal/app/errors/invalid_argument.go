package errors

func NewInvalidArgument(msg string, cause error, details ErrorDetails) *AppError {
	return NewAppError(InvalidArgument, msg, cause, details)
}
