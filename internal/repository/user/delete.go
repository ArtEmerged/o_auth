package user

import (
	"context"
	"fmt"

	"github.com/ArtEmerged/library/client/db"

	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// DeleteUser deletes a user from the repository by ID.
func (r *userRepo) DeleteUser(ctx context.Context, id int64) error {
	query := `
	UPDATE public.users 
	SET status = $1
	WHERE id = $2 AND status = $3;`

	q := db.Query{
		Name:     "user_repository.DeleteUser",
		QueryRaw: query,
	}

	_, err := r.db.DB().ExecContext(ctx, q, modelRepo.StatusDeleted, id, modelRepo.StatusActive)
	if err != nil {
		return fmt.Errorf("delete user:%w", err)
	}

	return nil
}
