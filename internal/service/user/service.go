package user

import (
	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
	"github.com/ArtEmerged/o_auth-server/internal/service"
)

type userService struct {
	repo      repository.UserRepo
	txManager db.TxManager

	salt []byte
}

// New creates a new user service.
func New(repo repository.UserRepo, txManager db.TxManager, salt string) service.UserService {
	return &userService{
		repo:      repo,
		txManager: txManager,
		salt:      []byte(salt)}
}
