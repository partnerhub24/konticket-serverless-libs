package errs

import "net/http"

type AppErr struct {
	Message string
	Code    int
}

func (e AppErr) Error() string {
	return e.Message
}

func NewAppErr(message string, code int) AppErr {
	return AppErr{
		Message: message,
		Code:    code,
	}
}

func NewUnexpectedError() error {
	return AppErr{
		Code:    http.StatusInternalServerError,
		Message: "Unexpected error",
	}
}

func NewBadRequestError(message string) error {
	return AppErr{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewNotFoundError(message string) error {
	return AppErr{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
