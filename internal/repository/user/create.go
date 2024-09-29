package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ArtEmerged/library/client/db"
	"github.com/jackc/pgconn"

	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository/user/adapter"
	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// CreateUser creates a new user in the repository and returns the user ID.
func (r *userRepo) CreateUser(ctx context.Context, in *model.CreateUserRequest) (*model.UserInfo, error) {
	newUser := adapter.CreateUserRequestToRepo(in)

	query := `
	INSERT INTO public.users (name, email, pass_hash, created_at, status, role)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, created_at, updated_at, role;`

	userInfo := new(modelRepo.UserInfo)

	q := db.Query{
		Name:     "user_repository.CreateUser",
		QueryRaw: query,
	}

	err := r.db.DB().ScanOneContext(
		ctx,
		userInfo,
		q,
		newUser.Name,
		newUser.Email,
		newUser.PasswordHash,
		newUser.CreatedAt,
		modelRepo.StatusActive,
		newUser.Role,
	)
	if err != nil {
		log.Printf("ERROR: REPO: %s\n", err.Error())

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // duplicate key value violates unique constraint
				return nil, fmt.Errorf("%w: %s", model.ErrAlreadyExists, pgErr.Message)
			}
		}

		return nil, fmt.Errorf("create user:%w", err)
	}

	return adapter.UserInfoToLocal(userInfo), nil
}
