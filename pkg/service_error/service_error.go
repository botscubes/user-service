package service_error

// New error for service.
func New(code int, message string) *ServiceError {
	return &ServiceError{code, message}
}

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ServiceError) Error() string {
	return e.Message
}
