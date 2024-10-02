package model

import (
	"fmt"
	"strings"
	"time"
)

const userCacheKey = "users:user:%d"

// UserCacheKey returns key for chat
func UserCacheKey(chatID int64) string {
	return fmt.Sprintf(userCacheKey, chatID)
}

// Role represents a user role.
type Role int32

const (
	// RoleUnknown represents an unknown role.
	RoleUnknown Role = iota
	// RoleUser represents the user role with normal permissions.
	RoleUser
	// RoleAdmin represents the admin role with all permissions.
	RoleAdmin
)

// Status represents a user status.
type Status string

const (
	// StatusUnknown represents a user status unknown (default).
	StatusUnknown Status = "UNKNOWN"
	// StatusActive represents a active user.
	StatusActive Status = "ACTIVE"
	// StatusBlocked represents a blocked user.
	StatusBlocked Status = "BLOCKED"
	// StatusDeleted represents a deleted user.
	StatusDeleted Status = "DELETED"
)

// UpdateUserRequest represents an update user request.
type UpdateUserRequest struct {
	ID   int64
	Name string
	Role Role
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
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Role      Role       `json:"role"`
}

// CreateUserRequest represents a create user request.
type CreateUserRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	PasswordHash    string
	Role            Role
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
