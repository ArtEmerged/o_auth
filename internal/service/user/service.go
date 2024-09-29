package user

import (
	"github.com/ArtEmerged/library/client/cache"
	"github.com/ArtEmerged/library/client/db"

	"github.com/ArtEmerged/o_auth-server/internal/repository"
	"github.com/ArtEmerged/o_auth-server/internal/service"
)

type userService struct {
	repo      repository.UserRepo
	txManager db.TxManager
	cache     cache.Cache

	salt []byte
}

// New creates a new user service.
func New(repo repository.UserRepo, txManager db.TxManager, cache cache.Cache, salt string) service.UserService {
	return &userService{
		repo:      repo,
		txManager: txManager,
		cache:     cache,
		salt:      []byte(salt)}
}
