package definitions

import "context"

// UserService defines the methods for user service operations.
type UserService interface {
	// CreateUser creates a new user and returns the user ID.
	CreateUser(ctx context.Context, user *CreateUserRequest) (id int64, err error)
	// UpdateUser updates an existing user's information.
	UpdateUser(ctx context.Context, user *UpdateUserRequest) error
	// DeleteUser deletes a user by ID.
	DeleteUser(ctx context.Context, id int64) error
	// GetUser retrieves a user by ID and returns the user information.
	GetUser(ctx context.Context, id int64) (*UserInfo, error)
}

// UserRepo defines the methods for user repository operations.
type UserRepo interface {
	// CreateUser creates a new user in the repository and returns the user ID.
	CreateUser(ctx context.Context, user *CreateUserRequest) (id int64, err error)
	// UpdateUser updates an existing user's information in the repository.
	UpdateUser(ctx context.Context, user *UpdateUserRequest) error
	// DeleteUser deletes a user from the repository by ID.
	DeleteUser(ctx context.Context, id int64) error
	// GetUser retrieves a user from the repository by ID and returns the user information.
	GetUser(ctx context.Context, id int64) (*UserInfo, error)
}
