package repository

import (
	"context"
	"time"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// UserRepo defines the methods for user repository operations.
type UserRepo interface {
	// CreateUser creates a new user in the repository and returns the use info.
	CreateUser(ctx context.Context, user *model.CreateUserRequest) (*model.UserInfo, error)
	// UpdateUser updates an existing user's information in the repository.
	UpdateUser(ctx context.Context, user *model.UpdateUserRequest) (updateAt time.Time, err error)
	// DeleteUser deletes a user from the repository by ID.
	DeleteUser(ctx context.Context, id int64) error
	// GetUser retrieves a user from the repository by ID and returns the user information.
	GetUser(ctx context.Context, id int64) (*model.UserInfo, error)
}
