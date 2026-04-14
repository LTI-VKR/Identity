package repos

import "errors"

var (
	ErrNotFound = errors.New("profile not found")
	ErrConflict = errors.New("profile conflict")
)
