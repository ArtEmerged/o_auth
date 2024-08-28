package user

import (
	"context"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

// GetUser retrieves a user from the repository by ID and returns the user information.
func (s *userService) GetUser(ctx context.Context, id int64) (*model.UserInfo, error) {
	return s.repo.GetUser(ctx, id)
}
