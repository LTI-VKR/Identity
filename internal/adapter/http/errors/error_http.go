package errors

import "errors"

var (
	ErrInvalidJSON = errors.New("неверный формат JSON")
)
