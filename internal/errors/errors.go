package errors

// New error for service.
func New(code int, message string) error {
	return &ServiceError{code, message}
}

type ServiceError struct {
	code    int
	message string
}

func (e *ServiceError) Error() string {
	return e.message
}
