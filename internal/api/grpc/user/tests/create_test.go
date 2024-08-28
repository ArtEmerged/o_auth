package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ArtEmerged/o_auth-server/internal/api/grpc/user"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/service"
	"github.com/ArtEmerged/o_auth-server/internal/service/mocks"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

func TestCreateUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateUserRequest
	}

	var (
		mc             = minimock.NewController(t)
		ctx            = context.Background()
		userId         = int64(gofakeit.Number(1, 1000))
		userName       = gofakeit.FirstName()
		userEmail      = gofakeit.Email()
		userPassword   = gofakeit.Password(true, true, true, true, false, 8)
		userRole       = desc.Role_USER
		serviceErr     = fmt.Errorf("service error")
		serviceRequest = &model.CreateUserRequest{
			Name:            userName,
			Email:           userEmail,
			Password:        userPassword,
			PasswordConfirm: userPassword,
			Role:            model.UserRole(userRole),
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateUserResponse
		wantErr         error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success create user",
			args: args{
				ctx: ctx,
				req: &desc.CreateUserRequest{
					Name:            userName,
					Email:           userEmail,
					Password:        userPassword,
					PasswordConfirm: userPassword,
					Role:            desc.Role_USER,
				},
			},
			want: &desc.CreateUserResponse{
				Id: userId,
			},
			wantErr: nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.CreateUserMock.Expect(ctx, serviceRequest).Return(userId, nil)

				return mock
			},
		},
		{
			name: "service error ErrAlreadyExists",
			args: args{
				ctx: ctx,
				req: &desc.CreateUserRequest{
					Name:            userName,
					Email:           userEmail,
					Password:        userPassword,
					PasswordConfirm: userPassword,
					Role:            desc.Role_USER,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.AlreadyExists, model.ErrAlreadyExists.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.CreateUserMock.Expect(ctx, serviceRequest).Return(userId, model.ErrAlreadyExists)

				return mock
			},
		},
		{
			name: "service error ErrInvalidArgument",
			args: args{
				ctx: ctx,
				req: &desc.CreateUserRequest{
					Name:            userName,
					Email:           userEmail,
					Password:        userPassword,
					PasswordConfirm: userPassword,
					Role:            desc.Role_USER,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, model.ErrInvalidArgument.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.CreateUserMock.Expect(ctx, serviceRequest).Return(userId, model.ErrInvalidArgument)

				return mock
			},
		},
		{
			name: "service error internal",
			args: args{
				ctx: ctx,
				req: &desc.CreateUserRequest{
					Name:            userName,
					Email:           userEmail,
					Password:        userPassword,
					PasswordConfirm: userPassword,
					Role:            desc.Role_USER,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.CreateUserMock.Expect(ctx, serviceRequest).Return(userId, serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userService := tt.userServiceMock(mc)

			api := user.NewImplementation(userService)

			got, err := api.CreateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
