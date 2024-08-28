package user

import (
	"context"
)

// DeleteUser deletes a user by ID.
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}
