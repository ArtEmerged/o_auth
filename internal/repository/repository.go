package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
)

type userRepo struct {
	db *pgxpool.Pool
}

// New creates a new instance of userRepo with the given database connection pool.
// db - pointer to the PostgreSQL connection pool
func New(db *pgxpool.Pool) def.UserRepo {
	return &userRepo{db: db}
}
