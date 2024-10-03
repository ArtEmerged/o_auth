package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

// DeleteUser handles the gRPC request to delete a user by ID.
// It validates the input, deletes the user via the service, and returns an empty response.
func (s *Implementation) DeleteUser(ctx context.Context, in *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.userService.DeleteUser(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
