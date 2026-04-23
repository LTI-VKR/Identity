package postgres

import (
	"context"
	"errors"
	"fmt"
	"identity/internal/domain/entity"
	entityErrors "identity/internal/domain/errors"
	"strings"

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
	var emailExists, phoneExists bool

	err := r.pool.QueryRow(ctx, `
        SELECT 
            EXISTS(SELECT 1 FROM profiles WHERE email = $1),
            EXISTS(SELECT 1 FROM profiles WHERE phone = $2)
    `, p.Email, p.Phone).Scan(&emailExists, &phoneExists)
	if err != nil {
		return uuid.Nil, err
	}

	conflicts := make([]string, 0, 1)
	if emailExists {
		conflicts = append(conflicts, "email")
	}
	if phoneExists {
		conflicts = append(conflicts, "phone")
	}
	if len(conflicts) > 0 {
		return uuid.Nil, fmt.Errorf("%w: "+strings.Join(conflicts, ", "), entityErrors.ErrProfileConflict)
	}

	_, errEx := r.pool.Exec(ctx, `
        insert into profiles (user_id, username, email, phone, language, has_Gamification, created_at, updated_at)
        values ($1, $2, $3, $4, $5, $6, now(), null)
    `, p.UserId, p.Username, p.Email, p.Phone, p.Language, p.HasGamification)
	if errEx != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.Nil, errEx
		}
		return uuid.Nil, errEx
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
			return uuid.Nil, entityErrors.ErrProfileConflict
		}
		return uuid.Nil, err
	}
	if ct.RowsAffected() == 0 {
		return uuid.Nil, entityErrors.ErrProfileNotFound
	}
	return p.UserId, nil
}
