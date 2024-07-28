package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
	"github.com/jackc/pgx/v4"
)

func (r *userRepo) CreateUser(ctx context.Context, in *def.CreateUserRequest) (int64, error) {
	q := `
	INSERT INTO public.users (name, email, pass_hash, created_at, role)
	VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var id int64

	err := r.db.QueryRow(ctx, q, in.Name, in.Email, in.PasswordHash, time.Now(), in.Role).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("create user:%w", err)
	}

	return id, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, in *def.UpdateUserRequest) error {
	q := `
	UPDATE public.users
	SET name = $1, role = $2, updated_at = $3
	WHERE id = $4 AND deleted_at IS NULL;`

	_, err := r.db.Exec(ctx, q, in.Name, in.Role, time.Now(), in.ID)
	if err != nil {
		return fmt.Errorf("updated user:%w", err)
	}
	return nil
}

func (r *userRepo) DeleteUser(ctx context.Context, id int64) error {
	q := `
	UPDATE public.users 
	SET deleted_at = $1
	WHERE id = $2 AND deleted_at IS NULL;`

	_, err := r.db.Exec(ctx, q, time.Now(), id)
	if err != nil {
		return fmt.Errorf("delete user:%w", err)
	}

	return nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*def.User, error) {
	q := `
	SELECT id, name, email, created_at, updated_at, role
	FROM public.users 
	WHERE id = $1 AND deleted_at IS NULL;
	`

	resp := new(def.User)

	err := r.db.QueryRow(ctx, q, id).Scan(
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
