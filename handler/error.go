package handler

import (
	"errors"
	"net/http"
)

type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

func NewNotFoundErr(err error) Error {
	return StatusError{
		Code: http.StatusNotFound,
		Err:  errors.New("Not found"),
	}
}