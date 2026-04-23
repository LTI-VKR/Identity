package handlers

import (
	"context"
	"encoding/json"
	httperr "identity/internal/adapter/http/errors"
	"identity/internal/application/command"
	"identity/internal/application/query"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type GetAvatarUploadUrlHandler struct {
	command *command.GetAvatarUploadUrlCommand
}

func NewGetAvatarUploadUrlHandler(svc *command.GetAvatarUploadUrlCommand) *GetAvatarUploadUrlHandler {
	return &GetAvatarUploadUrlHandler{command: svc}
}

// GetAvatarUploadUrl godoc
// @Summary Получение ссылки для загрузки аватарки
// @Tags avatar
// @Accept json
// @Produce json
// @Param user_id path string true "user id (uuid)"
// @Success 201 {object} dto.ProfileResponse
// @Router /avatar/{user_id} [post]
func (h *GetAvatarUploadUrlHandler) GetAvatarUploadUrl(w http.ResponseWriter, r *http.Request) {
	userIdString := chi.URLParam(r, "user_id")
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		mapped := httperr.NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", "invalid user_id")
		httperr.WriteError(w, r, mapped)
		return
	}

	url, err := h.command.Execute(context.Background(), userId)
	if err != nil {
		httperr.WriteError(w, r, httperr.Map(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(url)
}

type GetAvatarDownloadUrlHandler struct {
	query *query.GetAvatarDownloadUrlQuery
}

func NewGetAvatarDownloadUrlHandler(svc *query.GetAvatarDownloadUrlQuery) *GetAvatarDownloadUrlHandler {
	return &GetAvatarDownloadUrlHandler{query: svc}
}

// GetAvatarDownloadUrl godoc
// @Summary Получение ссылки для скачивания аватарки
// @Tags avatar
// @Accept json
// @Produce json
// @Param user_id path string true "user id (uuid)"
// @Success 200 {object} string
// @Router /avatar/{user_id} [get]
func (h *GetAvatarDownloadUrlHandler) GetAvatarDownloadUrl(w http.ResponseWriter, r *http.Request) {
	userIdString := chi.URLParam(r, "user_id")
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		mapped := httperr.NewBasicMapped(http.StatusBadRequest, "INVALID_ARGUMENT", "invalid user_id")
		httperr.WriteError(w, r, mapped)
		return
	}

	url, err := h.query.Execute(context.Background(), userId)
	if err != nil {
		httperr.WriteError(w, r, httperr.Map(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(url)
}
