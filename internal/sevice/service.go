package service

import (
	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
)

type userService struct {
	repo def.UserRepo

	salt []byte
}

var _ def.UserService = (*userService)(nil)

func New(repo def.UserRepo, salt string) *userService {
	return &userService{repo: repo, salt: []byte(salt)}
}
