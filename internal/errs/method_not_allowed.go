package errs

import "net/http"

type MethodNotAllowedError struct {
	message string
}

func (e MethodNotAllowedError) Error() string {
	return e.message
}

func (e MethodNotAllowedError) Code() int {
	return http.StatusMethodNotAllowed
}

func MethodNotAllowed(message string) MethodNotAllowedError {
	return MethodNotAllowedError{
		message: message,
	}
}
