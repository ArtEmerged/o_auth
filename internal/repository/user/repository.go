package user

import (
	dbClient "github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
)

const (
	tableUsers = "public.users"

	tableUsersIDColumn        = "id"
	tableUsersNameColumn      = "name"
	tableUsersEmailColumn     = "email"
	tableUsersPassHashColumn  = "pass_hash"
	tableUsersCreatedAtColumn = "created_at"
	tableUsersUpdatedAtColumn = "updated_at"
	tableUsersStatusColumn    = "status"
	tableUsersRoleColumn      = "role"
)

type userRepo struct {
	db dbClient.Client
}

// New creates a new instance of userRepo with the given database connection pool.
// db - database connection pool
func New(db dbClient.Client) repository.UserRepo {
	return &userRepo{db: db}
}
