package errors

type InvalidFieldError struct {
	message string
}

func NewInvalidFieldError(message string) *InvalidFieldError {
	return &InvalidFieldError{
		message: message,
	}
}

func (e InvalidFieldError) Error() string {
	return e.message
}
