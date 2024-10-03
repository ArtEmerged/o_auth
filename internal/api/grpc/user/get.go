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

// GetUser handles the gRPC request to retrieve a user by ID.
// It validates the input, retrieves the user via the service, and returns the user information.
func (s *Implementation) GetUser(ctx context.Context, in *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userInfo, err := s.userService.GetUser(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.GetUserResponse{UserInfo: adapter.UserInfoToProto(userInfo)}, nil
}
