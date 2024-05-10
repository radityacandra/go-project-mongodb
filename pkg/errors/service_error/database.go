package service_error

type DatabaseError struct {
	message string
}

func NewDatabaseError(message string) *DatabaseError {
	return &DatabaseError{
		message: message,
	}
}

func (e *DatabaseError) Error() string {
	return e.message
}
