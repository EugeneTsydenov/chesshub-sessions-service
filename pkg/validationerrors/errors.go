package validationerrors

type ErrorDetails map[string]string

type Error struct {
	Details ErrorDetails
	Msg     string
}

func New(msg string, details ErrorDetails) *Error {
	return &Error{
		Details: details,
		Msg:     msg,
	}
}

func (e *Error) Error() string {
	return e.Msg
}
