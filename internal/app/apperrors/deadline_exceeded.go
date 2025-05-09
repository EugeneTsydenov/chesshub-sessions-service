package apperrors

func NewDeadlineExceededError(msg string) *AppError {
	return &AppError{
		Code:    DeadlineExceeded,
		Message: msg,
	}
}
