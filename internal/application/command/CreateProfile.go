package command

import (
	"context"
	"identity/internal/application/model"
	"identity/internal/domain/repos"

	"github.com/google/uuid"
)

type CreateProfileCommand struct {
	repo repos.ProfileCommandRepository
}

func NewCreateProfileCommand(repo repos.ProfileCommandRepository) *CreateProfileCommand {
	return &CreateProfileCommand{repo: repo}
}

func (c *CreateProfileCommand) Execute(ctx context.Context, p model.ProfileModel) (uuid.UUID, error) {

	profile, err := p.ToProfile()
	if err != nil {
		return uuid.Nil, err
	}
	profile.UserId = uuid.New()
	result, err := c.repo.Create(ctx, profile)
	if err != nil {
		return uuid.Nil, err
	}

	return result, nil
}
