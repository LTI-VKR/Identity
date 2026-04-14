package query

import (
	"context"
	"identity/internal/application/model"
	"identity/internal/domain/repos"

	"github.com/google/uuid"
)

type GetProfileQuery struct {
	repo repos.ProfileQueryRepository
}

func NewGetProfileQuery(repo repos.ProfileQueryRepository) *GetProfileQuery {
	return &GetProfileQuery{repo: repo}
}

func (q *GetProfileQuery) Execute(ctx context.Context, userID uuid.UUID) (model.ProfileModel, error) {
	profile, err := q.repo.GetByUserID(ctx, userID)
	if err != nil {
		return model.ProfileModel{}, err
	}

	result := model.NewProfileModelFromEntity(profile)
	return result, nil
}
