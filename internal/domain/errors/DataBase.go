package errors

import "errors"

var (
	ErrProfileNotFound = errors.New("профиль не найден")
	ErrProfileConflict = errors.New("конфликт. Такой профиль уже существует")
)
