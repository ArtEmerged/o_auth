package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ArtEmerged/o_auth-server/internal/adapter"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

// UpdateUser handles the gRPC request to update an existing user.
// It validates the input, updates the user via the service, and returns an empty response.
func (s *Implementation) UpdateUser(ctx context.Context, in *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.userService.UpdateUser(ctx, adapter.UpdateUserRequestToLocal(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
