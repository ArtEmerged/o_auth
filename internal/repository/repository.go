package repository

import (
	"github.com/ArtEmerged/o_auth-server/internal/definitions"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

var _ definitions.UserRepo = (*userRepo)(nil)

func New(db *pgxpool.Pool) *userRepo {
	return &userRepo{db: db}
}
