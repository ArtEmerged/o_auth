package repository

import (
	"context"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// UserRepo defines the methods for user repository operations.
type UserRepo interface {
	// CreateUser creates a new user in the repository and returns the user ID.
	CreateUser(ctx context.Context, user *model.CreateUserRequest) (id int64, err error)
	// UpdateUser updates an existing user's information in the repository.
	UpdateUser(ctx context.Context, user *model.UpdateUserRequest) error
	// DeleteUser deletes a user from the repository by ID.
	DeleteUser(ctx context.Context, id int64) error
	// GetUser retrieves a user from the repository by ID and returns the user information.
	GetUser(ctx context.Context, id int64) (*model.UserInfo, error)
}
