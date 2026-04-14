package command

import (
	"context"
	"identity/internal/application/model"
	"identity/internal/domain/repos"

	"github.com/google/uuid"
)

type UpdateProfileCommand struct {
	repo repos.ProfileCommandRepository
}

func NewUpdateProfileCommand(repo repos.ProfileCommandRepository) *UpdateProfileCommand {
	return &UpdateProfileCommand{repo: repo}
}

func (c *UpdateProfileCommand) Execute(ctx context.Context, p model.ProfileModel) (uuid.UUID, error) {
	profileEntity, err := p.ToProfile()
	if err != nil {
		return uuid.Nil, err
	}

	UserId, err := c.repo.Update(ctx, profileEntity)
	if err != nil {
		return uuid.Nil, err
	}

	return UserId, nil
}
