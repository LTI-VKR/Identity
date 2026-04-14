package repos

import (
	"context"
	"identity/internal/domain/entity"

	"github.com/google/uuid"
)

type ProfileQueryRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (entity.Profile, error)
}
