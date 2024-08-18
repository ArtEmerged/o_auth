package user

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ArtEmerged/o_auth-server/internal/adapter"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

// CreateUser handles the gRPC request to create a new user.
// It validates the input, creates the user via the service, and returns the user ID.
func (s *Implementation) CreateUser(ctx context.Context, in *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	if in.GetName() == "" || in.GetEmail() == "" || in.GetPassword() == "" || in.GetPasswordConfirm() == "" {
		return nil, status.Error(codes.InvalidArgument, "all fields're required")
	}

	if in.GetPassword() != in.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	id, err := s.userService.CreateUser(ctx, adapter.CreateUserRequestToLocal(in))
	if err != nil {
		if errors.Is(err, model.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateUserResponse{Id: id}, nil
}
