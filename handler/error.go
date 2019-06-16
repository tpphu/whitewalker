package handler

import (
	"net/http"
)

// Error is a custom interface to handle error
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Error returns the origin error message
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status method return error code
func (se StatusError) Status() int {
	return se.Code
}

// NewNotFoundErr presents for 404 error
func NewNotFoundErr(err error) Error {
	return StatusError{
		Code: http.StatusNotFound,
		Err:  err,
	}
}
