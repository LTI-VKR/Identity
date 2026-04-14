package value_object

import (
	"errors"
	"regexp"
	"strings"
)

type PhoneValue string

var (
	errInvalidPhone = errors.New("invalid phone number")
	phoneRegexp     = regexp.MustCompile(`^\+7\d{10}$`)
)

func NewPhoneValue(value string) (PhoneValue, error) {
	phone := PhoneValue(strings.TrimSpace(value))
	if !phone.isValid() {
		return "", errInvalidPhone
	}
	return phone, nil
}

func (value PhoneValue) isValid() bool {
	return phoneRegexp.MatchString(string(value))
}

func (value PhoneValue) String() string {
	return string(value)
}

func (value PhoneValue) Equals(other PhoneValue) bool {
	return value == other
}
