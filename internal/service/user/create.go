package user

import (
	"context"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// CreateUser creates a new user and returns the user ID.
func (s *userService) CreateUser(ctx context.Context, in *model.CreateUserRequest) (int64, error) {
	if err := in.Validate(); err != nil {
		return -1, err
	}

	in.PasswordHash = s.hashSha256(in.Password)

	return s.repo.CreateUser(ctx, in)
}
