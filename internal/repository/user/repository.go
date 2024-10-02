package user

import (
	"github.com/ArtEmerged/library/client/cache"
	dbClient "github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_auth-server/internal/repository"
)

type userRepo struct {
	db    dbClient.Client
	cache cache.Cache
}

// New creates a new instance of userRepo with the given database connection pool.
// db - database connection pool
func New(db dbClient.Client, cache cache.Cache) repository.UserRepo {
	return &userRepo{
		db:    db,
		cache: cache,
	}
}
