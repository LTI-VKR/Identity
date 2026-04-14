package command

import (
	"context"
	"fmt"
	"identity/internal/domain/repos"
	"log"

	"github.com/google/uuid"
)

type GetAvatarUploadUrlCommand struct {
	repo repos.AvatarObjectRepository
}

func NewGetAvatarQuery(repo repos.AvatarObjectRepository) *GetAvatarUploadUrlCommand {
	return &GetAvatarUploadUrlCommand{repo: repo}
}

func (a *GetAvatarUploadUrlCommand) Execute(ctx context.Context, userID uuid.UUID) (string, error) {
	objectName := fmt.Sprintf("avatars/%s.jpg", userID)

	result, err := a.repo.GetUploadURL(ctx, objectName)
	log.Print(err)
	if err != nil {
		return "", err
	}

	return result, nil
}
