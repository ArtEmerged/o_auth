package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
	"github.com/ArtEmerged/o_auth-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_auth-server/internal/service/user"
)

func TestDeleteUser(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepo
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx    context.Context
		userID int64
	}
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryErr = fmt.Errorf("repository error")
		userID        = int64(gofakeit.Number(1, 1000))
	)

	tests := []struct {
		name          string
		args          args
		wantErr       error
		userRepoMock  userRepoMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success delete user",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			wantErr: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				mock.DeleteUserMock.Expect(ctx, userID).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "repository error",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			wantErr: repositoryErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				mock.DeleteUserMock.Expect(ctx, userID).Return(repositoryErr)
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
			err := service.DeleteUser(tt.args.ctx, tt.args.userID)
			require.Equal(t, tt.wantErr, err)
		})
	}
}
