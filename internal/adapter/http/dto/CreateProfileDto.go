package dto

type CreateProfileResponse struct {
	UserID string `json:"user_id"`
}

type CreateProfileRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Language        string `json:"language"`
	HasGamification string `json:"has_gamification"`
}
