package main

import (
	"identity/config"
	api "identity/internal/adapter/http"
	_ "identity/internal/adapter/http/docs"
	"identity/internal/adapter/http/handlers"
	"identity/internal/application/command"
	"identity/internal/application/query"
	"identity/internal/infrastructure/minIO"
	"identity/internal/infrastructure/postgres"
	"net/http"
)

// @title Identity API
// @version 0.1.0
// @description API сервиса профиля
// @BasePath /
func main() {
	port := "8080"

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	postgresPool, err := postgres.NewPool(cfg.DatabaseUrl)
	if err != nil {
		panic("failed to create pool")
	}
	defer postgresPool.Close()

	MinioClient, err := minIO.NewMinioClient(cfg.Endpoint, cfg.Login, cfg.Password)
	if err != nil {
		panic("failed to create minio client")
	}

	// Репозитории
	cmdRepo := postgres.NewProfileCommandRepository(postgresPool)
	qryRepo := postgres.NewProfileQueryRepository(postgresPool)

	minioRepo := minIO.NewAvatarMinioRepository(MinioClient, cfg.Bucket)

	// Команды и Запросы
	createCmd := command.NewCreateProfileCommand(cmdRepo)
	updateCmd := command.NewUpdateProfileCommand(cmdRepo)
	getQry := query.NewGetProfileQuery(qryRepo)
	getAvatarUploadUrlCmd := command.NewGetAvatarQuery(minioRepo)
	getAvatarDownloadUrlCmd := query.NewGetAvatarQuery(minioRepo)

	// Хендлеры
	createHandler := handlers.NewCreateProfileHandler(createCmd)
	getHandler := handlers.NewGetProfileHandler(getQry)
	updateHandler := handlers.NewUpdateProfileHandler(updateCmd, getQry)
	GetAvatarUploadUrlHandler := handlers.NewGetAvatarUploadUrlHandler(getAvatarUploadUrlCmd)
	GetAvatarDownloadUrlHandler := handlers.NewGetAvatarDownloadUrlHandler(getAvatarDownloadUrlCmd)

	// Роутер
	router := api.NewRouter(createHandler, getHandler, updateHandler, GetAvatarDownloadUrlHandler, GetAvatarUploadUrlHandler)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}
