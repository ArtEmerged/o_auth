package user

import (
	"context"
	"fmt"

	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// DeleteUser deletes a user from the repository by ID.
func (r *userRepo) DeleteUser(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`
	UPDATE %[1]s 
	SET %[2]s = $1
	WHERE %[3]s = $2 AND %[2]s = $3;`,
		tableUsers, // 1

		tableUsersStatusColumn, // 2
		tableUsersIDColumn,     // 3
	)

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
