package user

import (
	"context"
	"fmt"

	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository/user/adapter"
	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// UpdateUser updates an existing user's information.
func (r *userRepo) UpdateUser(ctx context.Context, in *model.UpdateUserRequest) error {
	upUser := adapter.UpdateUserRequestToRepo(in)

	query := `
	UPDATE public.users
	SET name = $1, role = $2, updated_at = $3
	WHERE id = $4 AND status = $5;`

	q := db.Query{
		Name:     "user_repository.UpdateUser",
		QueryRaw: query,
	}

	_, err := r.db.DB().ExecContext(ctx, q, upUser.Name, upUser.Role, upUser.UpdatedAt, upUser.ID, modelRepo.StatusActive)
	if err != nil {
		return fmt.Errorf("updated user:%w", err)
	}

	return nil
}
