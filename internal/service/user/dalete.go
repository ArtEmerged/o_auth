package user

import (
	"context"
	"log"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// DeleteUser deletes a user by ID.
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	err = s.cache.Del(ctx, model.UserCacheKey(id))
	if err != nil {
		log.Printf("WARN: failed to delete user in cache: %s\n", err.Error())
	}

	return nil
}
