package service_error

type GeneralError struct {
	message string
}

func NewGeneralError(message string) *GeneralError {
	return &GeneralError{
		message: message,
	}
}

func (e *GeneralError) Error() string {
	return e.message
}
