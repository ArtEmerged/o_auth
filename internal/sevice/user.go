package service

import (
	"context"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
)

func (s *userService) CreateUser(ctx context.Context, in *def.CreateUserRequest) (int64, error) {
	in.PasswordHash = s.hashMd5(in.Password)

	return s.repo.CreateUser(ctx, in)
}

func (s *userService) UpdateUser(ctx context.Context, in *def.UpdateUserRequest) error {
	if in.Name == nil && in.Role == def.ROLE_UNKNOWN {
		return nil
	}

	user, err := s.repo.GetUser(ctx, in.ID)
	if err != nil {
		return err
	}

	if in.Name == nil {
		in.Name = &user.Name
	}

	if in.Role == def.ROLE_UNKNOWN {
		in.Role = user.Role
	}

	return s.repo.UpdateUser(ctx, in)
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *userService) GetUser(ctx context.Context, id int64) (*def.User, error) {
	return s.repo.GetUser(ctx, id)
}
