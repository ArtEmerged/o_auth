package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"

	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository/user/adapter"
	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// CreateUser creates a new user in the repository and returns the user ID.
func (r *userRepo) CreateUser(ctx context.Context, in *model.CreateUserRequest) (int64, error) {
	newUser := adapter.CreateUserRequestToRepo(in)

	query := fmt.Sprintf(`
	INSERT INTO %[1]s (%[2]s, %[3]s, %[4]s, %[5]s, %[6]s, %[7]s)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING %[8]s;`,
		tableUsers, // 1

		tableUsersNameColumn,      // 2
		tableUsersEmailColumn,     // 3
		tableUsersPassHashColumn,  // 4
		tableUsersCreatedAtColumn, // 5
		tableUsersStatusColumn,    // 6
		tableUsersRoleColumn,      // 7
		tableUsersIDColumn,        // 8

	)

	var id int64

	q := db.Query{
		Name:     "user_repository.CreateUser",
		QueryRaw: query,
	}

	err := r.db.DB().QueryRowContext(ctx, q, newUser.Name, newUser.Email, newUser.PasswordHash, newUser.CreatedAt, modelRepo.StatusActive, newUser.Role).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // duplicate key value violates unique constraint
				return -1, fmt.Errorf("user with email %s %w", in.Email, model.ErrAlreadyExists)
			}

		}
		return -1, fmt.Errorf("create user:%w", err)
	}

	return id, nil
}
