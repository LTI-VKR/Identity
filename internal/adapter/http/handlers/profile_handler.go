package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"identity/internal/adapter/http/dto"
	httperr "identity/internal/adapter/http/errors"
	"identity/internal/application/command"
	"identity/internal/application/model"
	"identity/internal/application/query"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CreateProfileHandler struct {
	command *command.CreateProfileCommand
}

func NewCreateProfileHandler(svc *command.CreateProfileCommand) *CreateProfileHandler {
	return &CreateProfileHandler{command: svc}
}

type GetProfileHandler struct {
	query *query.GetProfileQuery
}

func NewGetProfileHandler(svc *query.GetProfileQuery) *GetProfileHandler {
	return &GetProfileHandler{query: svc}
}

type UpdateProfileHandler struct {
	cmd *command.UpdateProfileCommand
	get *query.GetProfileQuery
}

func NewUpdateProfileHandler(cmd *command.UpdateProfileCommand, get *query.GetProfileQuery) *UpdateProfileHandler {
	return &UpdateProfileHandler{cmd: cmd, get: get}
}

// CreateProfile godoc
// @Summary Create profile
// @Tags profiles
// @Accept json
// @Produce json
// @Param body body dto.CreateProfileRequest true "create profile"
// @Success 201 {object} dto.CreateProfileResponse
// @Failure 400 {object} errors.ErrorBasicResponse
// @Router /profiles [post]
func (h *CreateProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperr.Write(w, r, httperr.NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", "invalid json body"))
		return
	}
	hasGamificationBool, err := strconv.ParseBool(req.HasGamification)
	if err != nil {
		httperr.Write(w, r, httperr.NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", "invalid json body (hasGamification)"))
		return
	}

	profileModel := model.ProfileModel{
		Username:        req.Username,
		Email:           req.Email,
		Phone:           req.Phone,
		Language:        req.Language,
		HasGamification: hasGamificationBool,
	}

	userId, err := h.command.Execute(r.Context(), profileModel)
	if err != nil {
		httperr.Write(w, r, httperr.Map(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(dto.CreateProfileResponse{
		UserID: userId.String(),
	})
}

// GetProfile godoc
// @Summary Get profile by user id
// @Tags profiles
// @Produce json
// @Param user_id path string true "user id (uuid)"
// @Success 200 {object} dto.ProfileResponse
// @Failure 400 {object} errors.ErrorBasicResponse
// @Router /profiles/{user_id} [get]
func (h *GetProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "user_id")
	userID, err := uuid.Parse(rawID)
	if err != nil {
		httperr.Write(w, r, httperr.NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", "invalid user_id"))
		return
	}

	p, err := h.query.Execute(r.Context(), userID)
	if err != nil {
		httperr.Write(w, r, httperr.Map(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dto.ProfileResponse{
		UserID:   p.UserID.String(),
		Username: p.Username,
		Email:    p.Email,
		Phone:    p.Phone,
		Language: p.Language,
	})
}

// UpdateProfile godoc
// @Summary Patch profile
// @Tags profiles
// @Accept json
// @Produce json
// @Param body body dto.UpdateProfileRequest true "update profile"
// @Param user_id path string true "user id (uuid)"
// @Success 200 {object} dto.UpdateProfileResponse
// @Failure 400 {object} errors.ErrorValidationResponse
// @Router /profiles/{user_id} [patch]
func (h *UpdateProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "user_id")
	userID, errId := uuid.Parse(rawID)

	var req dto.UpdateProfileRequest
	errBody := json.NewDecoder(r.Body).Decode(&req)

	mappedErrors := make(map[string]string)
	if errId != nil {
		mappedErrors["user_id"] = "invalid user_id in route"
	}
	if errBody != nil {
		mappedErrors["body"] = "invalid json body"

	}
	if len(mappedErrors) > 0 {
		httperr.Write(w, r, httperr.NewValidatingMapped(http.StatusBadRequest, "INVALID_ARGUMENT", mappedErrors))
		return
	}

	current, err := h.get.Execute(r.Context(), userID)
	if err != nil {
		httperr.Write(w, r, httperr.Map(err))
		return
	}

	userId, err := h.cmd.Execute(r.Context(), current)
	if err != nil {
		httperr.Write(w, r, httperr.Map(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dto.UpdateProfileResponse{
		UserID: userId.String(),
	})
}
