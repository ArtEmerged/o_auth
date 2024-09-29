package user

import (
	"context"
	"log"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// GetUser retrieves a user from the repository by ID and returns the user information.
func (s *userService) GetUser(ctx context.Context, id int64) (*model.UserInfo, error) {
	userInfo := new(model.UserInfo)

	err := s.cache.Get(ctx, model.UserCacheKey(id), userInfo)
	if err == nil {
		return userInfo, nil
	}

	log.Printf("WARN: failed to get user in cache: %s\n", err.Error())

	userInfo, err = s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.cache.Set(ctx, model.UserCacheKey(userInfo.ID), userInfo, 0)
	if err != nil {
		log.Printf("WARN: failed to set user in cache: %s\n", err.Error())
	}

	return userInfo, nil
}
