package definitions

import (
	"time"
)

// UserRole represents a user role.
type UserRole int32

const (
	// RoleUnknown represents an unknown role.
	RoleUnknown UserRole = iota
	// RoleUser represents the user role with normal permissions.
	RoleUser
	// RoleAdmin represents the admin role with all permissions.
	RoleAdmin
)

// UserStatus represents a user status.
type UserStatus string

const (
	// StatusUnknown represents a user status unknown (default).
	StatusUnknown UserStatus = "UNKNOWN"
	// StatusActive represents a active user.
	StatusActive UserStatus = "ACTIVE"
	// StatusBlocked represents a blocked user.
	StatusBlocked UserStatus = "BLOCKED"
	// StatusDeleted represents a deleted user.
	StatusDeleted UserStatus = "DELETED"
)

// UpdateUserRequest represents an update user request.
type UpdateUserRequest struct {
	ID   int64
	Name *string
	Role UserRole
}

// Validate - validates the update user request.
func (r *UpdateUserRequest) Validate() error {
	if r.Name == nil && r.Role == RoleUnknown {
		return ErrWithoutChanges
	}

	return nil
}

// UserInfo represents a user information.
type UserInfo struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
	Role      UserRole
}

// CreateUserRequest represents a create user request.
type CreateUserRequest struct {
	Name         string
	Email        string
	Password     string
	PasswordHash string
	Role         UserRole
}
