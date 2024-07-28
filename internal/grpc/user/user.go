package grpc_user

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

func (s *userServer) CreateUser(ctx context.Context, in *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	if in.GetName() == "" || in.GetEmail() == "" || in.GetPassword() == "" || in.GetPasswordConfirm() == "" {
		return nil, status.Error(codes.InvalidArgument, "all fields're required")
	}

	if in.GetPassword() != in.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	id, err := s.service.CreateUser(ctx, def.AdaptedCreateUserRequestToLocal(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.CreateUserResponse{Id: id}, nil
}

func (s *userServer) GetUser(ctx context.Context, in *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	if in.Id < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	user, err := s.service.GetUser(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, def.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &desc.GetUserResponse{UserInfo: user.ToProto()}, nil
}
func (s *userServer) UpdateUser(ctx context.Context, in *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	err := s.service.UpdateUser(ctx, def.AdaptedUpdateUserRequestToLocal(in))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}

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
