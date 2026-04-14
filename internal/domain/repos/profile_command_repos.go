package repos

import (
	"context"
	"identity/internal/domain/entity"

	"github.com/google/uuid"
)

type ProfileCommandRepository interface {
	Create(ctx context.Context, p entity.Profile) (uuid.UUID, error)
	Update(ctx context.Context, p entity.Profile) (uuid.UUID, error)
}
