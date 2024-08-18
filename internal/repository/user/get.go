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
	query := fmt.Sprintf(`
	SELECT %[2]s, %[3]s, %[4]s, %[5]s, %[6]s, %[7]s
	FROM %[1]s
	WHERE %[2]s = $1 AND %[8]s = $2;`,
		tableUsers, // 1

		tableUsersIDColumn,        // 2
		tableUsersNameColumn,      // 3
		tableUsersEmailColumn,     // 4
		tableUsersCreatedAtColumn, // 5
		tableUsersUpdatedAtColumn, // 6
		tableUsersRoleColumn,      // 7
		tableUsersStatusColumn,    // 8
	)

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
