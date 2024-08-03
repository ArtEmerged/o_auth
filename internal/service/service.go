package service

import (
	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
)

type userService struct {
	repo def.UserRepo

	salt []byte
}

// New creates a new user service.
func New(repo def.UserRepo, salt string) def.UserService {
	return &userService{repo: repo, salt: []byte(salt)}
}
