package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
)

func (r *userRepo) CreateUser(ctx context.Context, in *def.CreateUserRequest) (int64, error) {
	q := `
	INSERT INTO public.users (name, email, pass_hash, created_at, status, role)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	var id int64

	createdAt := time.Now().UTC()

	err := r.db.QueryRow(ctx, q, in.Name, in.Email, in.PasswordHash, createdAt, def.StatusActive, in.Role).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return -1, fmt.Errorf("user with email %s %w", in.Email, def.ErrAlreadyExists)
			}

		}
		return -1, fmt.Errorf("create user:%w", err)
	}

	return id, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, in *def.UpdateUserRequest) error {
	q := `
	UPDATE public.users
	SET name = $1, role = $2, updated_at = $3
	WHERE id = $4 AND status = $5;`

	updateAt := time.Now().UTC()

	_, err := r.db.Exec(ctx, q, in.Name, in.Role, updateAt, in.ID, def.StatusActive)
	if err != nil {
		return fmt.Errorf("updated user:%w", err)
	}

	return nil
}

func (r *userRepo) DeleteUser(ctx context.Context, id int64) error {
	q := `
	UPDATE public.users 
	SET status = $1
	WHERE id = $2 AND status = $3;`

	_, err := r.db.Exec(ctx, q, def.StatusDeleted, id, def.StatusActive)
	if err != nil {
		return fmt.Errorf("delete user:%w", err)
	}

	return nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*def.UserInfo, error) {
	q := `
	SELECT id, name, email, created_at, updated_at, role
	FROM public.users 
	WHERE id = $1 AND status = $2;
	`

	resp := new(def.UserInfo)

	err := r.db.QueryRow(ctx, q, id, def.StatusActive).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Email,
		&resp.CreatedAt,
		&resp.UpdatedAt,
		&resp.Role,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user %w", def.ErrNotFound)
		}
		return nil, fmt.Errorf("failed get user:%w", err)
	}

	return resp, nil
}
