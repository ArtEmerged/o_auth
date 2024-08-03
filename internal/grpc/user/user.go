package grpc_user

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
	"github.com/ArtEmerged/o_auth-server/internal/definitions/adapter"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

// CreateUser handles the gRPC request to create a new user.
// It validates the input, creates the user via the service, and returns the user ID.
func (s *userServer) CreateUser(ctx context.Context, in *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	if in.GetName() == "" || in.GetEmail() == "" || in.GetPassword() == "" || in.GetPasswordConfirm() == "" {
		return nil, status.Error(codes.InvalidArgument, "all fields're required")
	}

	if in.GetPassword() != in.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	id, err := s.service.CreateUser(ctx, adapter.CreateUserRequestToLocal(in))
	if err != nil {
		if errors.Is(err, def.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateUserResponse{Id: id}, nil
}

// GetUser handles the gRPC request to retrieve a user by ID.
// It validates the input, retrieves the user via the service, and returns the user information.
func (s *userServer) GetUser(ctx context.Context, in *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	if in.Id < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	userInfo, err := s.service.GetUser(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, def.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.GetUserResponse{UserInfo: adapter.UserInfoToProto(userInfo)}, nil
}

// UpdateUser handles the gRPC request to update an existing user.
// It validates the input, updates the user via the service, and returns an empty response.
func (s *userServer) UpdateUser(ctx context.Context, in *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	err := s.service.UpdateUser(ctx, adapter.UpdateUserRequestToLocal(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}

// DeleteUser handles the gRPC request to delete a user by ID.
// It validates the input, deletes the user via the service, and returns an empty response.
func (s *userServer) DeleteUser(ctx context.Context, in *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	err := s.service.DeleteUser(ctx, in.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
