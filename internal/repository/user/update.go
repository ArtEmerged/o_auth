package user

import (
	"context"
	"fmt"
	"time"

	"github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository/user/adapter"
	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// UpdateUser updates an existing user's information.
func (r *userRepo) UpdateUser(ctx context.Context, in *model.UpdateUserRequest) (updateAt time.Time, err error) {
	upUser := adapter.UpdateUserRequestToRepo(in)

	query := `
	UPDATE public.users
	SET name = $1, role = $2, updated_at = $3
	WHERE id = $4 AND status = $5;`

	q := db.Query{
		Name:     "user_repository.UpdateUser",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, upUser.Name, upUser.Role, upUser.UpdatedAt, upUser.ID, modelRepo.StatusActive)
	if err != nil {
		return time.Time{}, fmt.Errorf("updated user:%w", err)
	}

	return upUser.UpdatedAt, nil
}
