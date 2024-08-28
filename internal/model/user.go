package model

import (
	"fmt"
	"strings"
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
	Name string
	Role UserRole
}

// Validate - validates the update user request.
func (r *UpdateUserRequest) Validate() error {
	if r.Name == "" && r.Role == RoleUnknown {
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
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	PasswordHash    string
	Role            UserRole
}

// Validate - validates the create user request.
func (r *CreateUserRequest) Validate() error {
	var errsText []string

	// validate required fields
	if r.Name == "" {
		errsText = append(errsText, "field name is required")
	}
	if r.Email == "" {
		errsText = append(errsText, "field email is required")
	}
	if r.Password == "" {
		errsText = append(errsText, "field password is required")
	}
	if r.PasswordConfirm == "" {
		errsText = append(errsText, "field password_confirm is required")
	}

	// validate password
	if r.Password != r.PasswordConfirm {
		errsText = append(errsText, "password and password_confirm don't match")
	}

	if len(errsText) > 0 {
		return fmt.Errorf("%w: %s", ErrInvalidArgument, strings.Join(errsText, ", "))
	}

	return nil
}
