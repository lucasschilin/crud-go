package rest_err

import (
	"net/http"
)

type RestErr struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, nil)
}

func NewBadRequestValidationError(message string, causes []Cause) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, causes)
}

func NewInternalServerError(message string) *RestErr {
	return NewRestErr(
		message, "internal_server_error", http.StatusInternalServerError, nil,
	)
}

func NewNotFoundError(message string) *RestErr {
	return NewRestErr(message, "not_found", http.StatusNotFound, nil)
}

func NewForbiddenError(message string) *RestErr {
	return NewRestErr(message, "forbidden", http.StatusForbidden, nil)
}
