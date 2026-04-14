package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresCfg
	MinioCfg
}

type PostgresCfg struct {
	DatabaseUrl string
}

type MinioCfg struct {
	Endpoint string
	Login    string
	Password string
	Bucket   string
}

func NewConfig() (*Config, error) {
	postgresCfg := PostgresCfg{}
	minioCfg := MinioCfg{}

	if os.Getenv("ENVIRONMENT") != "DEV" {
		if err := godotenv.Load(); err != nil {
			return &Config{}, err
		}
	}

	postgresCfg.DatabaseUrl = os.Getenv("DATABASE_URL")

	minioCfg.Bucket = os.Getenv("AVATAR_BUCKET")
	minioCfg.Endpoint = os.Getenv("MINIO_ENDPOINT")
	minioCfg.Login = os.Getenv("MINIO_LOGIN")
	minioCfg.Password = os.Getenv("MINIO_PASSWORD")

	return &Config{PostgresCfg: postgresCfg, MinioCfg: minioCfg}, nil
}
