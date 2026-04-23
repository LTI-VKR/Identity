package handlers

import (
	"encoding/json"
	binder "identity/internal/adapter/http/binding"
	"identity/internal/adapter/http/dto"
	httperr "identity/internal/adapter/http/errors"
	"identity/internal/application/command"
	"identity/internal/application/model"
	"identity/internal/application/query"
	"net/http"

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
// @Failure 400 {object} dto.ErrorResponse "INVALID_ARGUMENT: invalid json body"
// @Failure 409 {object} dto.ErrorResponse "CONFLICT"
// @Failure 422 {object} dto.ErrorValidationResponse
// @Failure 500 {object} dto.ErrorResponse "INTERNAL"
// @Router /profiles [post]
func (h *CreateProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProfileRequest
	err := binder.DecodeAndValidate(r, &req)
	if err != nil {
		httperr.WriteError(w, r, httperr.Map(err))
		return
	}

	profileModel := model.ProfileModel{
		Username:        req.Username,
		Email:           req.Email,
		Phone:           req.Phone,
		Language:        req.Language,
		HasGamification: req.HasGamification,
	}

	userId, err := h.command.Execute(r.Context(), profileModel)
	if err != nil {
		httperr.WriteError(w, r, httperr.Map(err))
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
// @Failure 400 {object} dto.ErrorResponse "INVALID_ARGUMENT: invalid user_id"
// @Failure 404 {object} dto.ErrorResponse "NOT_FOUND"
// @Failure 409 {object} dto.ErrorResponse "CONFLICT"
// @Failure 422 {object} dto.ErrorValidationResponse
// @Failure 500 {object} dto.ErrorResponse "INTERNAL"
// @Router /profiles/{user_id} [get]
func (h *GetProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	rawID := chi.URLParam(r, "user_id")
	userID, err := uuid.Parse(rawID)
	if err != nil {
		mapped := httperr.NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", "invalid user_id")
		httperr.WriteError(w, r, mapped)
		return
	}

	p, err := h.query.Execute(r.Context(), userID)
	if err != nil {
		httperr.WriteError(w, r, httperr.Map(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
// @Success 204 {object} dto.UpdateProfileResponse
// @Failure 400 {object} dto.ErrorValidationResponse
// @Failure 404 {object} dto.ErrorResponse "NOT_FOUND"
// @Failure 409 {object} dto.ErrorResponse "CONFLICT"
// @Failure 422 {object} dto.ErrorResponse "INVALID_ARGUMENT (from domain/app via Map)"
// @Failure 500 {object} dto.ErrorResponse "INTERNAL"
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
		mapped := httperr.NewValidatingMapped(http.StatusBadRequest, "INVALID_ARGUMENT", mappedErrors)
		httperr.WriteError(w, r, mapped)
		return
	}

	current, err := h.get.Execute(r.Context(), userID)
	if err != nil {
		httperr.WriteError(w, r, httperr.Map(err))
		return
	}

	userId, err := h.cmd.Execute(r.Context(), current)
	if err != nil {
		httperr.WriteError(w, r, httperr.Map(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(dto.UpdateProfileResponse{
		UserID: userId.String(),
	})
}
