package model

import "time"

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

// CreateUserRequest represents a create user request.
type CreateUserRequest struct {
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"pass_hash"`
	Role         int32     `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
}

// UserInfo represents a user information.
type UserInfo struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Role      int32      `db:"role"`
}

// UpdateUserRequest represents an update user request.
type UpdateUserRequest struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Role      int32     `db:"role"`
	UpdatedAt time.Time `db:"updated_at"`
}
