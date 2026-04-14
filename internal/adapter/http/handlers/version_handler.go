package handlers

import (
	"encoding/json"
	"net/http"

	"identity/internal/application/model"
)

type versionResponse struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

// Version godoc
// @Summary Service version
// @Description Текущая версия сервиса
// @Tags system
// @Produce json
// @Success 200 {object} versionResponse
// @Router /version [get]
func Version(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(versionResponse{
		Service: "identity",
		Version: model.Version,
	})
}
