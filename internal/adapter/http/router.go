package http

import (
	"identity/internal/adapter/http/handlers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewRouter(
	createProfileHandler *handlers.CreateProfileHandler,
	getProfileHandler *handlers.GetProfileHandler,
	updateProfileHandler *handlers.UpdateProfileHandler,
	getAvatarDownloadUrlHandler *handlers.GetAvatarDownloadUrlHandler,
	getAvatarUploadUrlHandler *handlers.GetAvatarUploadUrlHandler,

) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Get("/health", handlers.Health)
	r.Get("/version", handlers.Version)

	r.Post("/profiles", createProfileHandler.CreateProfile)
	r.Get("/profiles/{user_id}", getProfileHandler.GetProfile)
	r.Patch("/profiles/{user_id}", updateProfileHandler.UpdateProfile)

	r.Get("/avatar/{user_id}", getAvatarDownloadUrlHandler.GetAvatarDownloadUrl)
	r.Post("/avatar/{user_id}", getAvatarUploadUrlHandler.GetAvatarUploadUrl)

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))
	return r
}
