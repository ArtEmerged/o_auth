package grpc_user

import (
	"context"

	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ desc.UserV1Server = (*userServer)(nil)

func (s *userServer) CreateUser(ctx context.Context, in *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	if in.GetName() == "" || in.GetEmail() == "" || in.GetPassword() == "" || in.GetPasswordConfirm() == "" || in.GetRole() == desc.Role_UNKNOWN {
		return nil, status.Error(codes.InvalidArgument, "all fields're required")
	}

	id := gofakeit.Number(1, 99999)
	return &desc.CreateUserResponse{Id: int64(id)}, nil
}

func (s *userServer) GetUser(ctx context.Context, in *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	if in.Id < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}

	userInfo := &desc.UserInfo{
		Id:        in.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Timestamp: &desc.UserInfo_CreatedAt{timestamppb.New(gofakeit.Date())},
		Role:      desc.Role_USER,
	}
	return &desc.GetUserResponse{
		UserInfo: userInfo,
	}, nil
}

func (s *userServer) UpdateUser(ctx context.Context, in *desc.UpdateUserRequest) (*desc.UpdateUserResponse, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}
	return nil, nil
}

func (s *userServer) DeleteUser(ctx context.Context, in *desc.DeleteUserRequest) (*desc.DeleteUserResponse, error) {
	if in.GetId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "negative id")
	}
	return nil, nil
}
