package query

import (
	"context"
	"fmt"
	"identity/internal/domain/repos"

	"github.com/google/uuid"
)

type GetAvatarDownloadUrlQuery struct {
	repo repos.AvatarObjectRepository
}

func NewGetAvatarQuery(repo repos.AvatarObjectRepository) *GetAvatarDownloadUrlQuery {
	return &GetAvatarDownloadUrlQuery{repo: repo}
}

func (a *GetAvatarDownloadUrlQuery) Execute(ctx context.Context, userID uuid.UUID) (string, error) {
	objectName := fmt.Sprintf("avatars/%s.jpg", userID)

	result, err := a.repo.GetDownloadURL(ctx, objectName)
	if err != nil {
		return "", err
	}

	return result, nil
}
