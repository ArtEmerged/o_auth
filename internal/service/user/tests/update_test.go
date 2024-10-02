package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ArtEmerged/library/client/cache"
	cacheMock "github.com/ArtEmerged/library/client/cache/mocks"
	"github.com/ArtEmerged/library/client/db"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_auth-server/internal/model"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
	"github.com/ArtEmerged/o_auth-server/internal/repository/mocks"
	"github.com/ArtEmerged/o_auth-server/internal/service/user"
)

func TestUpdateUser(t *testing.T) {
	type userRepoMockFunc func(mc *minimock.Controller) repository.UserRepo
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager
	type cacheMockFunc func(mc *minimock.Controller) cache.Cache

	type args struct {
		ctx context.Context
		req *model.UpdateUserRequest
	}
	var (
		ctx               = context.Background()
		mc                = minimock.NewController(t)
		updateUserRepoErr = fmt.Errorf("get user error")
		getUserRepoErr    = fmt.Errorf("update user error")
		userID            = int64(gofakeit.Number(1, 1000))
		newUserName       = gofakeit.Name()
		oldUserName       = gofakeit.Name()
		oldRole           = model.RoleUser
		newRole           = model.RoleAdmin

		email     = gofakeit.Email()
		createdAt = time.Now().UTC()
		updatedAt = time.Now().UTC()
	)

	tests := []struct {
		name          string
		args          args
		wantErr       error
		userRepoMock  userRepoMockFunc
		cacheMock     cacheMockFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success update user",
			args: args{
				ctx: ctx,
				req: &model.UpdateUserRequest{
					ID:   userID,
					Name: newUserName,
					Role: newRole,
				},
			},
			wantErr: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				getUserResp := &model.UserInfo{
					ID:        userID,
					Name:      oldUserName,
					Email:     email,
					Role:      oldRole,
					CreatedAt: createdAt,
				}
				mock.GetUserMock.Expect(ctx, userID).Return(getUserResp, nil)

				updateUserReq := &model.UpdateUserRequest{
					ID:   userID,
					Name: newUserName,
					Role: newRole,
				}
				mock.UpdateUserMock.Expect(ctx, updateUserReq).Return(updatedAt, nil)
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.Cache {
				mock := cacheMock.NewCacheMock(mc)
				userInfo := &model.UserInfo{
					ID:        userID,
					Name:      newUserName,
					Email:     email,
					Role:      newRole,
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				}
				mock.SetMock.Expect(ctx, model.UserCacheKey(userID), userInfo, 0).Return(nil)

				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "success update user only role",
			args: args{
				ctx: ctx,
				req: &model.UpdateUserRequest{
					ID:   userID,
					Name: "",
					Role: model.RoleAdmin,
				},
			},
			wantErr: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				getUserResp := &model.UserInfo{
					ID:        userID,
					Name:      oldUserName,
					Email:     email,
					Role:      oldRole,
					CreatedAt: createdAt,
				}
				mock.GetUserMock.Expect(ctx, userID).Return(getUserResp, nil)

				updateUserReq := &model.UpdateUserRequest{
					ID:   userID,
					Name: oldUserName,
					Role: newRole,
				}
				mock.UpdateUserMock.Expect(ctx, updateUserReq).Return(updatedAt, nil)
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.Cache {
				mock := cacheMock.NewCacheMock(mc)
				userInfo := &model.UserInfo{
					ID:        userID,
					Name:      oldUserName,
					Email:     email,
					Role:      newRole,
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				}
				mock.SetMock.Expect(ctx, model.UserCacheKey(userID), userInfo, 0).Return(nil)

				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "success update user only name",
			args: args{
				ctx: ctx,
				req: &model.UpdateUserRequest{
					ID:   userID,
					Name: newUserName,
					Role: model.RoleUnknown,
				},
			},
			wantErr: nil,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				getUserResp := &model.UserInfo{
					ID:        userID,
					Name:      oldUserName,
					Email:     email,
					Role:      oldRole,
					CreatedAt: createdAt,
				}
				mock.GetUserMock.Expect(ctx, userID).Return(getUserResp, nil)

				updateUserReq := &model.UpdateUserRequest{
					ID:   userID,
					Name: newUserName,
					Role: oldRole,
				}
				mock.UpdateUserMock.Expect(ctx, updateUserReq).Return(updatedAt, nil)
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.Cache {
				mock := cacheMock.NewCacheMock(mc)
				userInfo := &model.UserInfo{
					ID:        userID,
					Name:      newUserName,
					Email:     email,
					Role:      oldRole,
					CreatedAt: createdAt,
					UpdatedAt: &updatedAt,
				}
				mock.SetMock.Expect(ctx, model.UserCacheKey(userID), userInfo, 0).Return(nil)

				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "error get user repo",
			args: args{
				ctx: ctx,
				req: &model.UpdateUserRequest{
					ID:   userID,
					Name: newUserName,
					Role: model.RoleUser,
				},
			},
			wantErr: getUserRepoErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)

				mock.GetUserMock.Expect(ctx, userID).Return(nil, getUserRepoErr)
				return mock
			},
			cacheMock: func(mc *minimock.Controller) cache.Cache {
				mock := cacheMock.NewCacheMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "error update user repo",
			args: args{
				ctx: ctx,
				req: &model.UpdateUserRequest{
					ID:   userID,
					Name: newUserName,
					Role: model.RoleUser,
				},
			},
			wantErr: updateUserRepoErr,
			userRepoMock: func(mc *minimock.Controller) repository.UserRepo {
				mock := mocks.NewUserRepoMock(mc)
				req := &model.UpdateUserRequest{
					ID:   userID,
					Name: newUserName,
					Role: model.RoleUser,
				}
				mock.GetUserMock.Expect(ctx, userID).Return(&model.UserInfo{
					ID:   userID,
					Name: oldUserName,
					Role: model.RoleUser,
				}, nil)
				mock.UpdateUserMock.Expect(ctx, req).Return(time.Time{}, updateUserRepoErr)
				return mock
			},
			cacheMock:     func(mc *minimock.Controller) cache.Cache { return nil },
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
		{
			name: "without update",
			args: args{
				ctx: ctx,
				req: &model.UpdateUserRequest{
					ID: userID,
				},
			},
			wantErr:       nil,
			userRepoMock:  func(mc *minimock.Controller) repository.UserRepo { return nil },
			cacheMock:     func(mc *minimock.Controller) cache.Cache { return nil },
			txManagerMock: func(mc *minimock.Controller) db.TxManager { return nil },
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userRepo := tt.userRepoMock(mc)
			txManager := tt.txManagerMock(mc)
			cache := tt.cacheMock(mc)

			service := user.New(userRepo, txManager, cache, "")
			err := service.UpdateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.wantErr, err)
		})
	}
}
