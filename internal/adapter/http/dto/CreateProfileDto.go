package dto

type CreateProfileResponse struct {
	UserID string `json:"user_id"`
}

type CreateProfileRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
	Language        string `json:"language" validate:"required"`
	HasGamification bool   `json:"has_gamification" default:"true" validate:"required"`
}
