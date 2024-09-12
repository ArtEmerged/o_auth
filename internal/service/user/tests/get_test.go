package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
	"github.com/ArtEmerged/o_auth-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_auth-server/internal/service/user"
)

func TestGetUser(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepo
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx    context.Context
		userID int64
	}
	var (
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		createAt = time.Now()
		updated  = createAt.AddDate(0, 0, 1).Add(time.Hour * 2)

		repositoryErr = fmt.Errorf("repository error")
		userID        = int64(gofakeit.Number(1, 1000))
		response      = &model.UserInfo{
			ID:        userID,
			Name:      gofakeit.FirstName(),
			Email:     gofakeit.Email(),
			CreatedAt: createAt,
			UpdatedAt: &updated,
			Role:      model.RoleUser,
		}
	)

	tests := []struct {
		name          string
		args          args
		want          *model.UserInfo
		wantErr       error
		userRepoMock  userRepoMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success get user",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want:    response,
			wantErr: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				mock.GetUserMock.Expect(ctx, userID).Return(response, nil)
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
			want:    nil,
			wantErr: repositoryErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				mock.GetUserMock.Expect(ctx, userID).Return(nil, repositoryErr)
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
			got, err := service.GetUser(tt.args.ctx, tt.args.userID)
			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
