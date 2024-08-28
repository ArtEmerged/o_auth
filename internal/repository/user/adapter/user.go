package adapter

import (
	"time"

	"github.com/ArtEmerged/o_auth-server/internal/model"
	modelRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user/model"
)

// CreateUserRequestToRepo adapts the CreateUserRequest to repo model.
func CreateUserRequestToRepo(in *model.CreateUserRequest) *modelRepo.CreateUserRequest {
	return &modelRepo.CreateUserRequest{
		Name:         in.Name,
		Email:        in.Email,
		PasswordHash: in.PasswordHash,
		Role:         int32(in.Role),
		CreatedAt:    time.Now().UTC(),
	}
}

// UpdateUserRequestToRepo adapts the UpdateUserRequest to repo model.
func UpdateUserRequestToRepo(in *model.UpdateUserRequest) *modelRepo.UpdateUserRequest {
	return &modelRepo.UpdateUserRequest{
		ID:        in.ID,
		Name:      in.Name,
		Role:      int32(in.Role),
		UpdatedAt: time.Now().UTC(),
	}
}

// UserInfoToLocal adapts the UserInfo to the local model.
func UserInfoToLocal(in *modelRepo.UserInfo) *model.UserInfo {
	out := &model.UserInfo{
		ID:        in.ID,
		Name:      in.Name,
		Email:     in.Email,
		Role:      model.UserRole(in.Role),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
	return out
}
