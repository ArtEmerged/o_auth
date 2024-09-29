package user

import (
	"context"
	"log"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// CreateUser creates a new user and returns the user ID.
func (s *userService) CreateUser(ctx context.Context, in *model.CreateUserRequest) (int64, error) {
	if err := in.Validate(); err != nil {
		return -1, err
	}

	in.PasswordHash = s.hashSha256(in.Password)

	userInfo, err := s.repo.CreateUser(ctx, in)
	if err != nil {
		return -1, err
	}

	err = s.cache.Set(ctx, model.UserCacheKey(userInfo.ID), userInfo, 0)
	if err != nil {
		log.Printf("WARN: failed to set user in cache: %s\n", err.Error())
	}

	return userInfo.ID, nil
}
