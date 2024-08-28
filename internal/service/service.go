package service

import (
	"context"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// UserService defines the methods for user service operations.
type UserService interface {
	// CreateUser creates a new user and returns the user ID.
	CreateUser(ctx context.Context, user *model.CreateUserRequest) (id int64, err error)
	// UpdateUser updates an existing user's information.
	UpdateUser(ctx context.Context, user *model.UpdateUserRequest) error
	// DeleteUser deletes a user by ID.
	DeleteUser(ctx context.Context, id int64) error
	// GetUser retrieves a user by ID and returns the user information.
	GetUser(ctx context.Context, id int64) (*model.UserInfo, error)
}
