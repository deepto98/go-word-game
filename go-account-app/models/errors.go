package model

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorType string

const (
	Authorization   ErrorType = "AUTHORIZATION"
	BadRequest      ErrorType = "BADREQUEST"
	Conflict        ErrorType = "CONFLICT"
	Internal        ErrorType = "INTERNAL"
	NotFound        ErrorType = "NOTFOUND"
	PayloadTooLarge ErrorType = "PAYLOADTOOLARGE"
)

//Custom error type to return at api endpoints
type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

//Satisfies error interface of standard go library
func (e *Error) Error() string {
	return e.Message
}

//Maps errors to status codes
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

//Checks runtime type of error and returns http status code
//if the error is in this model
func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

//---------------------------------
// Factory methods to build errors
//---------------------------------

func NewAuthorizationError(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}
func NewBadRequestError(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad Request. Reason : %v", reason),
	}
}

func NewConflictError(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("Resource %v with value %v already exists", name, value),
	}
}

func NewInternalError() *Error {
	return &Error{
		Type:    Internal,
		Message: "Internal server error.",
	}
}

func NewNotFoundError(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("Resource %v with value %v not found", name, value),
	}
}

func NewPayloadTooLargeError(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		Type: PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v",
			maxBodySize, contentLength),
	}
}
