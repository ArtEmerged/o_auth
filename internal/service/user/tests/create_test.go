package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
	"github.com/ArtEmerged/o_auth-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_auth-server/internal/service/user"
)

func TestCreateUser(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepo
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.CreateUserRequest
	}
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryErr = fmt.Errorf("repository error")
		userID        = int64(gofakeit.Number(1, 1000))
		userPassword  = gofakeit.Password(true, true, true, true, false, 8)
		req           = model.CreateUserRequest{
			Name:            gofakeit.FirstName(),
			Email:           gofakeit.Email(),
			Password:        userPassword,
			PasswordConfirm: userPassword,
			Role:            model.RoleUser,
		}
	)

	tests := []struct {
		name          string
		args          args
		want          int64
		wantErr       error
		userRepoMock  userRepoMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success create user",
			args: args{
				ctx: ctx,
				req: &req,
			},
			want:    userID,
			wantErr: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				mock.CreateUserMock.Expect(ctx, &req).Return(userID, nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "validation error", // without name
			args: args{
				ctx: ctx,
				req: &model.CreateUserRequest{
					Email:           gofakeit.Email(),
					Password:        userPassword,
					PasswordConfirm: userPassword,
					Role:            model.RoleUser,
				},
			},
			want:          -1,
			wantErr:       fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field name is required"),
			userRepoMock:  func(mc *minimock.Controller) repository.UserRepo { return nil },
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: &req,
			},
			want:    -1,
			wantErr: repositoryErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				mock.CreateUserMock.Expect(ctx, &req).Return(-1, repositoryErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userRepo := tt.userRepoMock(mc)
			txManager := tt.txManagerMock(mc)

			service := user.New(userRepo, txManager, "")
			got, err := service.CreateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
