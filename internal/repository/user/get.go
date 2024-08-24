package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository/user/adapter"
	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// GetUser retrieves a user from the repository by ID and returns the user information.
func (r *userRepo) GetUser(ctx context.Context, id int64) (*model.UserInfo, error) {
	query := `
	SELECT id, name, email, created_at, updated_at, role, status
	FROM public.users
	WHERE id = $1 AND status = $2;`

	q := db.Query{
		Name:     "user_repository.GetUser",
		QueryRaw: query,
	}

	userInfo := new(modelRepo.UserInfo)

	err := r.db.DB().ScanOneContext(ctx, &userInfo, q, id, modelRepo.StatusActive)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user %w", model.ErrNotFound)
		}
		return nil, fmt.Errorf("failed get user:%w", err)
	}

	return adapter.UserInfoToLocal(userInfo), nil
}
