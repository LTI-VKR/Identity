package errors

import (
	"errors"
	domainErrors "identity/internal/domain/errors"
	"net/http"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrNotFound        = errors.New("not found")
	ErrConflict        = errors.New("conflict")
)

type MappedError struct {
	Status  int
	Code    string
	Message string
	Fields  map[string]string
}

func NewBasicMapped(status int, code, message string) MappedError {
	return MappedError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

func NewValidatingMapped(status int, code string, Fields map[string]string) MappedError {
	return MappedError{
		Status: status,
		Code:   code,
		Fields: Fields,
	}
}

func Map(err error) MappedError {
	var valErr *domainErrors.ValidationError

	switch {
	case errors.Is(err, ErrInvalidArgument):
		return NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", err.Error())
	case errors.Is(err, ErrNotFound):
		return NewBasicMapped(http.StatusNotFound, "NOT_FOUND", err.Error())
	case errors.Is(err, ErrConflict):
		return NewBasicMapped(http.StatusConflict, "CONFLICT", err.Error())
	case errors.As(err, &valErr):
		return NewValidatingMapped(http.StatusBadRequest, "VALIDATION_FAILED", valErr.Fields)
	default:
		return NewBasicMapped(http.StatusInternalServerError, "INTERNAL", err.Error())
	}
}
