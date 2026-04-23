package errors

import (
	"errors"
	"identity/internal/adapter/http/dto"
	domainErrors "identity/internal/domain/errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorBody interface {
	SetRequestID(string)
}

type MappedError struct {
	Status int
	Body   ErrorBody
}

func NewBasicMapped(status int, code, message string) MappedError {
	return MappedError{
		Status: status,
		Body: &dto.ErrorResponse{
			Code:    code,
			Message: message,
		},
	}
}

func NewValidatingMapped(status int, code string, fields map[string]string) MappedError {
	return MappedError{
		Status: status,
		Body: &dto.ErrorValidationResponse{
			Code:   code,
			Fields: fields,
		},
	}
}

func Map(err error) MappedError {
	var valErr *domainErrors.ValidationError
	var jsonValErr validator.ValidationErrors

	switch {
	case errors.Is(err, domainErrors.ErrProfileNotFound):
		return NewBasicMapped(http.StatusNotFound, "NOT_FOUND", err.Error())
	case errors.Is(err, domainErrors.ErrProfileConflict):
		return NewBasicMapped(http.StatusConflict, "CONFLICT", err.Error())
	case errors.As(err, &valErr):
		return NewValidatingMapped(http.StatusUnprocessableEntity, "VALIDATION_FAILED", valErr.Fields)
	case errors.Is(err, ErrInvalidJSON):
		return NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", "invalid json body")
	case errors.As(err, &jsonValErr):
		fields := make(map[string]string)

		for _, e := range jsonValErr {
			jsonField := e.Field()

			if tag := e.StructField(); tag != "" {
				jsonField = e.Field()
			}

			fields[jsonField] = e.Tag()
		}
		return NewValidatingMapped(http.StatusUnprocessableEntity, "VALIDATION_FAILED", fields)
	default:
		return NewBasicMapped(http.StatusInternalServerError, "INTERNAL", err.Error())
	}
}
