package entity

import (
	"identity/internal/domain/value_object"
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	UserId          uuid.UUID
	Username        string
	Email           value_object.EmailValue
	Phone           value_object.PhoneValue
	Language        string
	HasGamification bool
	AtCreated       *time.Time
	AtUpdated       *time.Time
}

func NewProfile(
	userId uuid.UUID,
	username string,
	email value_object.EmailValue,
	phone value_object.PhoneValue,
	language string,
	hasGamification bool,
	atCreated *time.Time,
	atUpdated *time.Time) Profile {

	return Profile{
		UserId:          userId,
		Username:        username,
		Email:           email,
		Phone:           phone,
		Language:        language,
		HasGamification: hasGamification,
		AtCreated:       atCreated,
		AtUpdated:       atUpdated,
	}
}
