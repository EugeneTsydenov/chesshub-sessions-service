package apperrors

func NewInternalError(msg string) *AppError {
	return &AppError{
		Code:    Internal,
		Message: msg,
	}
}
