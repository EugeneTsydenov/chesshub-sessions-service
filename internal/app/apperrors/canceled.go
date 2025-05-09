package apperrors

func NewCanceledError(msg string) *AppError {
	return &AppError{
		Code:    Canceled,
		Message: msg,
	}
}
