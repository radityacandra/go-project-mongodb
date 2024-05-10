package service_error

type ConflictError struct {
	message string
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		message: message,
	}
}

func (e *ConflictError) Error() string {
	return e.message
}
