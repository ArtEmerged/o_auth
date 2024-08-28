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
	"github.com/ArtEmerged/o_auth-server/internal/service"
	"github.com/ArtEmerged/o_auth-server/internal/service/mocks"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

func TestDeleteUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.DeleteUserRequest
	}

	var (
		mc             = minimock.NewController(t)
		ctx            = context.Background()
		userId         = int64(gofakeit.Number(1, 1000))
		negativeUserId = int64(gofakeit.Number(-1000, 0))
		serviceErr     = fmt.Errorf("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		wantErr         error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success delete user",
			args: args{
				ctx: ctx,
				req: &desc.DeleteUserRequest{
					Id: userId,
				},
			},
			want:    nil,
			wantErr: nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.DeleteUserMock.Expect(ctx, userId).Return(nil)

				return mock
			},
		},
		{
			name: "error negative id",
			args: args{
				ctx: ctx,
				req: &desc.DeleteUserRequest{
					Id: negativeUserId,
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
				req: &desc.DeleteUserRequest{
					Id: userId,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.DeleteUserMock.Expect(ctx, userId).Return(serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userService := tt.userServiceMock(mc)

			api := user.NewImplementation(userService)

			got, err := api.DeleteUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
