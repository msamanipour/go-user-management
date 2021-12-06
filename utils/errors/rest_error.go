package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	ErrType string `json:"err_type"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		ErrType: "err",
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		ErrType: "err",
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		ErrType: "err",
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewSuccessMessage(message string) *RestErr {
	return &RestErr{
		Message: message,
		ErrType: "success",
		Status:  http.StatusOK,
		Error:   "success",
	}
}
