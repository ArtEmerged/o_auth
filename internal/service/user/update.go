package user

import (
	"context"
	"errors"

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

	return s.repo.UpdateUser(ctx, in)
}
