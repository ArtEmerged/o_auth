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
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ArtEmerged/o_auth-server/internal/api/grpc/user"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/service"
	"github.com/ArtEmerged/o_auth-server/internal/service/mocks"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

func TestUpdateUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateUserRequest
	}

	var (
		mc             = minimock.NewController(t)
		ctx            = context.Background()
		userId         = int64(gofakeit.Number(1, 1000))
		negativeUserId = int64(gofakeit.Number(-1000, 0))
		userName       = gofakeit.FirstName()
		userRole       = desc.Role_USER
		serviceErr     = fmt.Errorf("service error")
		serviceRequest = &model.UpdateUserRequest{
			ID:   userId,
			Name: userName,
			Role: model.UserRole(userRole),
		}
		serviceRequestWithoutName = &model.UpdateUserRequest{
			ID:   userId,
			Name: "",
			Role: model.UserRole(userRole),
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		wantErr         error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success update user with name and role",
			args: args{
				ctx: ctx,
				req: &desc.UpdateUserRequest{
					Id:   userId,
					Name: &userName,
					Role: desc.Role_USER,
				},
			},
			want:    nil,
			wantErr: nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.UpdateUserMock.Expect(ctx, serviceRequest).Return(nil)

				return mock
			},
		},
		{
			name: "success update user without name",
			args: args{
				ctx: ctx,
				req: &desc.UpdateUserRequest{
					Id:   userId,
					Name: nil,
					Role: desc.Role_USER,
				},
			},
			want:    nil,
			wantErr: nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.UpdateUserMock.Expect(ctx, serviceRequestWithoutName).Return(nil)

				return mock
			},
		},
		{
			name: "error negative id",
			args: args{
				ctx: ctx,
				req: &desc.UpdateUserRequest{
					Id:   negativeUserId,
					Name: &userName,
					Role: desc.Role_USER,
				},
			},
			want:            nil,
			wantErr:         status.Error(codes.InvalidArgument, "negative id"),
			userServiceMock: func(mc *minimock.Controller) service.UserService { return nil },
		},
		{
			name: "service error internal",
			args: args{
				ctx: ctx,
				req: &desc.UpdateUserRequest{
					Id:   userId,
					Name: &userName,
					Role: desc.Role_USER,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.UpdateUserMock.Expect(ctx, serviceRequest).Return(serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userService := tt.userServiceMock(mc)

			api := user.NewImplementation(userService)

			got, err := api.UpdateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
