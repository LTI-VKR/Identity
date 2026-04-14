package postgres

import (
	"context"
	"errors"
	"identity/internal/domain/entity"
	"identity/internal/domain/repos"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileQueryRepository struct {
	pool *pgxpool.Pool
}

func NewProfileQueryRepository(pool *pgxpool.Pool) *ProfileQueryRepository {
	return &ProfileQueryRepository{pool: pool}
}

func (r *ProfileQueryRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (entity.Profile, error) {
	p := entity.Profile{}
	err := r.pool.QueryRow(ctx, `
        select user_id, username, email, phone, language, has_Gamification, created_at, updated_at
        from profiles
        where user_id = $1
    `, userID).Scan(&p.UserId, &p.Username, &p.Email, &p.Phone, &p.Language, &p.HasGamification, &p.AtCreated, &p.AtUpdated)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Profile{}, repos.ErrNotFound
		}
		return entity.Profile{}, err
	}
	return p, nil
}
