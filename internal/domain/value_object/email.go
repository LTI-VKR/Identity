package value_object

import (
	"errors"
	"regexp"
	"strings"
)

type EmailValue string

var (
	errInvalidEmail = errors.New("invalid email address")
	emailRegexp     = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func NewEmailValue(value string) (EmailValue, error) {
	email := EmailValue(strings.TrimSpace(value))
	if !email.isValid() {
		return "", errInvalidEmail
	}
	return email, nil
}

func (value EmailValue) isValid() bool {
	return emailRegexp.MatchString(string(value))
}

func (value EmailValue) String() string {
	return string(value)
}

func (value EmailValue) Equals(other PhoneValue) bool {
	return value.String() == other.String()
}
