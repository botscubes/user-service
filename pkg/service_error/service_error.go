package service_error

// New error for service.
func New(code int, message string) *ServiceError {
	return &ServiceError{code, message}
}

type ServiceError struct {
	code    int
	message string
}

func (e *ServiceError) Error() string {
	return e.message
}
