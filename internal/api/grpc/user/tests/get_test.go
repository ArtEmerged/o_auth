package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ArtEmerged/o_auth-server/internal/api/grpc/user"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/service"
	"github.com/ArtEmerged/o_auth-server/internal/service/mocks"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

func TestGetUser(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetUserRequest
	}

	var (
		mc                           = minimock.NewController(t)
		ctx                          = context.Background()
		userId                       = int64(gofakeit.Number(1, 1000))
		negativeUserId               = int64(gofakeit.Number(-1000, 0))
		userName                     = gofakeit.FirstName()
		userEmail                    = gofakeit.Email()
		createdAt                    = time.Now().UTC()
		updatedAt                    = createdAt.AddDate(0, 0, 1)
		userRole                     = desc.Role_USER
		getUserResponseWithCreatedAt = &desc.GetUserResponse{
			UserInfo: &desc.UserInfo{
				Id:        userId,
				Name:      userName,
				Email:     userEmail,
				Role:      userRole,
				Timestamp: &desc.UserInfo_CreatedAt{CreatedAt: timestamppb.New(createdAt)},
			},
		}
		getUserResponseWithUpdatedAt = &desc.GetUserResponse{
			UserInfo: &desc.UserInfo{
				Id:        userId,
				Name:      userName,
				Email:     userEmail,
				Role:      userRole,
				Timestamp: &desc.UserInfo_UpdatedAt{UpdatedAt: timestamppb.New(updatedAt)},
			},
		}

		userInfoWithCreatedAt = &model.UserInfo{
			ID:        userId,
			Name:      userName,
			Email:     userEmail,
			Role:      model.Role(userRole),
			CreatedAt: createdAt,
			UpdatedAt: nil,
		}
		userInfoWithUpdatedAt = &model.UserInfo{
			ID:        userId,
			Name:      userName,
			Email:     userEmail,
			Role:      model.Role(userRole),
			CreatedAt: createdAt,
			UpdatedAt: &updatedAt,
		}
		serviceErr = fmt.Errorf("service error")
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetUserResponse
		wantErr         error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success get user with created at",
			args: args{
				ctx: ctx,
				req: &desc.GetUserRequest{
					Id: userId,
				},
			},
			want:    getUserResponseWithCreatedAt,
			wantErr: nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.GetUserMock.Expect(ctx, userId).Return(userInfoWithCreatedAt, nil)

				return mock
			},
		},
		{
			name: "success get user with updated at",
			args: args{
				ctx: ctx,
				req: &desc.GetUserRequest{
					Id: userId,
				},
			},
			want:    getUserResponseWithUpdatedAt,
			wantErr: nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.GetUserMock.Expect(ctx, userId).Return(userInfoWithUpdatedAt, nil)

				return mock
			},
		},
		{
			name: "error negative id",
			args: args{
				ctx: ctx,
				req: &desc.GetUserRequest{
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
				req: &desc.GetUserRequest{
					Id: userId,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.GetUserMock.Expect(ctx, userId).Return(nil, serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userService := tt.userServiceMock(mc)

			api := user.NewImplementation(userService)

			got, err := api.GetUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
