package errors

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Fields map[string]string
}

func NewValidationError() *ValidationError {
	return &ValidationError{make(map[string]string)}
}

func (e *ValidationError) Error() string {
	msgs := make([]string, 0, len(e.Fields))
	for field, msg := range e.Fields {
		msgs = append(msgs, fmt.Sprintf("%s: %s", field, msg))
	}
	return strings.Join(msgs, "; ")
}

func (e *ValidationError) Add(field, message string) {
	e.Fields[field] = message
}

func (e *ValidationError) HasErrors() bool {
	return len(e.Fields) > 0
}
