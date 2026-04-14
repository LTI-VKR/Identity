package model

import (
	"identity/internal/domain/entity"
	domainErrors "identity/internal/domain/errors"
	"identity/internal/domain/value_object"

	"github.com/google/uuid"
)

type ProfileModel struct {
	UserID          uuid.UUID
	Username        string
	Email           string
	Phone           string
	Language        string
	HasGamification bool
}

func (p *ProfileModel) ToProfile() (entity.Profile, error) {
	valErr := domainErrors.NewValidationError()

	email, err := value_object.NewEmailValue(p.Email)
	if err != nil {
		valErr.Add("email", err.Error())
	}

	phone, err := value_object.NewPhoneValue(p.Phone)
	if err != nil {
		valErr.Add("phone", err.Error())
	}

	if valErr.HasErrors() {
		return entity.Profile{}, valErr
	}

	return entity.NewProfile(
		p.UserID,
		p.Username,
		email,
		phone,
		p.Language,
		p.HasGamification,
		nil,
		nil,
	), nil
}

func NewProfileModelFromEntity(profile entity.Profile) ProfileModel {
	return ProfileModel{
		UserID:          profile.UserId,
		Username:        profile.Username,
		Email:           profile.Email.String(),
		Phone:           profile.Phone.String(),
		Language:        profile.Language,
		HasGamification: profile.HasGamification,
	}
}
