package resterrors

import "net/http"

type Causes struct {
	Field string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Message string `json:"message"`
	Err string `json:"error"`
	Code int `json:"code"`
	Causes []Causes `json:"causes"`
}

func (e *Error) ToString() string {
	return e.Message
}

func NewError(message string, err string, code int, causes []Causes) *Error {
	return &Error{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Error {
	return &Error{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Message: message,
		Err: "internal_server_error",
		Code: http.StatusInternalServerError,
	}
}

func NewForbiddenError(message string) *Error {
	return &Error{
		Message: message,
		Err: "forbidden",
		Code: http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Err: "not_found",
		Code: http.StatusNotFound,
	}
}