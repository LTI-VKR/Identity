package postgres

import (
	"context"
	"errors"
	"identity/internal/domain/entity"
	"identity/internal/domain/repos"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileCommandRepository struct {
	pool *pgxpool.Pool
}

func NewProfileCommandRepository(pool *pgxpool.Pool) *ProfileCommandRepository {
	return &ProfileCommandRepository{pool: pool}
}

func (r *ProfileCommandRepository) Create(ctx context.Context, p entity.Profile) (uuid.UUID, error) {
	_, err := r.pool.Exec(ctx, `
        insert into profiles (user_id, username, email, phone, language, has_Gamification, created_at, updated_at)
        values ($1, $2, $3, $4, $5, $6, now(), null)
    `, p.UserId, p.Username, p.Email, p.Phone, p.Language, p.HasGamification)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.Nil, repos.ErrConflict
		}
		return uuid.Nil, err
	}
	return p.UserId, nil
}

func (r *ProfileCommandRepository) Update(ctx context.Context, p entity.Profile) (uuid.UUID, error) {
	ct, err := r.pool.Exec(ctx, `
        update profiles
        set username = $2, email = $3, phone = $4, language = $5, has_Gamification = $6, created_at = $7,  updated_at = now()
        where user_id = $1
    `, p.UserId, p.Username, p.Email, p.Phone, p.Language, p.HasGamification, &p.AtCreated)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.Nil, repos.ErrConflict
		}
		return uuid.Nil, err
	}
	if ct.RowsAffected() == 0 {
		return uuid.Nil, repos.ErrNotFound
	}
	return p.UserId, nil
}
