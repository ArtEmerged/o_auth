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

	query := fmt.Sprintf(`
	UPDATE %[1]s
	SET %[2]s = $1, %[3]s = $2, %[4]s = $3
	WHERE %[5]s = $4 AND %[6]s = $5;`,
		tableUsers, // 1

		tableUsersNameColumn,      // 2
		tableUsersRoleColumn,      // 3
		tableUsersUpdatedAtColumn, // 4
		tableUsersIDColumn,        // 5
		tableUsersStatusColumn,    // 6
	)

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
