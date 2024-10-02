package user

import (
	"context"
	"errors"
	"log"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// UpdateUser updates an existing user in the repository.
func (s *userService) UpdateUser(ctx context.Context, in *model.UpdateUserRequest) error {
	err := in.Validate()
	if err != nil {
		if errors.Is(err, model.ErrWithoutChanges) {
			return nil
		}

		return err
	}

	user, err := s.repo.GetUser(ctx, in.ID)
	if err != nil {
		return err
	}

	if in.Name == "" {
		in.Name = user.Name
	}

	if in.Role == model.RoleUnknown {
		in.Role = user.Role
	}

	updatedAt, err := s.repo.UpdateUser(ctx, in)
	if err != nil {
		return err
	}

	userInfo := &model.UserInfo{
		ID:        in.ID,
		Name:      in.Name,
		Email:     user.Email,
		Role:      in.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: &updatedAt,
	}

	err = s.cache.Set(ctx, model.UserCacheKey(in.ID), userInfo, 0)
	if err != nil {
		log.Printf("WARN: failed to set user in cache: %s\n", err.Error())
	}

	return nil
}
