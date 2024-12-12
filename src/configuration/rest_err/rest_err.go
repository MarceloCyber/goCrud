package rest_err

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Err   string `json:"error"`
	Code int64 `json:"code"`
	Causes []Causes `json:"causes"`
}

type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}
func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
	}
}

func NewValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "validation_error",
		Code: 422, // Unprocessable Entity
		Causes: causes,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not_found",
		Code: http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "internal_server_error",
		Code: http.StatusInternalServerError,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "unauthorized",
		Code: http.StatusUnauthorized,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "forbidden",
		Code: http.StatusForbidden,
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "conflict",
		Code: http.StatusConflict,
	}
}
