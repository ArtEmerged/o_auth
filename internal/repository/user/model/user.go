package model

import "time"

// UserStatus represents a user status.
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

type Role string

const (
	// RoleUnknown represents a user role unknown (default).
	RoleUnknown Role = "UNKNOWN"
	// RoleUser represents the user role with normal permissions.
	RoleUser Role = "USER"
	// RoleAdmin represents the admin role with all permissions.
	RoleAdmin Role = "ADMIN"
)

// CreateUserRequest represents a create user request.
type CreateUserRequest struct {
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"pass_hash"`
	Role         Role      `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
}

// UserInfo represents a user information.
type UserInfo struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Role      Role       `db:"role"`
}

// UpdateUserRequest represents an update user request.
type UpdateUserRequest struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Role      Role      `db:"role"`
	UpdatedAt time.Time `db:"updated_at"`
}
