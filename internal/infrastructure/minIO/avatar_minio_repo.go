package minIO

import (
	"context"
	"time"

	"github.com/minio/minio-go/v7"
)

type AvatarMinioRepository struct {
	client     *minio.Client
	bucketName string
}

func NewAvatarMinioRepository(client *minio.Client, bucketName string) *AvatarMinioRepository {
	return &AvatarMinioRepository{
		client:     client,
		bucketName: bucketName,
	}
}

func (a AvatarMinioRepository) GetUploadURL(ctx context.Context, objectName string) (string, error) {
	u, err := a.client.PresignedPutObject(
		ctx,
		a.bucketName,
		objectName,
		10*time.Hour,
	)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

func (a AvatarMinioRepository) GetDownloadURL(ctx context.Context, objectName string) (string, error) {
	url, err := a.client.PresignedGetObject(
		ctx,
		a.bucketName,
		objectName,
		24*60*60*time.Second, // 24 часа
		nil,
	)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}
