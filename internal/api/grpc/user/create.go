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
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := s.userService.CreateUser(ctx, adapter.CreateUserRequestToLocal(in))
	if err != nil {
		if errors.Is(err, model.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		if errors.Is(err, model.ErrInvalidArgument) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateUserResponse{Id: id}, nil
}
