package adapter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ArtEmerged/o_auth-server/internal/model"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

// CreateUserRequestToLocal adapts the CreateUserRequest proto to the local model.
func CreateUserRequestToLocal(in *desc.CreateUserRequest) *model.CreateUserRequest {

	return &model.CreateUserRequest{
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		Role:     model.UserRole(in.Role),
	}
}

// UpdateUserRequestToLocal adapts the UpdateUserRequest proto to the local model.
func UpdateUserRequestToLocal(in *desc.UpdateUserRequest) *model.UpdateUserRequest {
	return &model.UpdateUserRequest{
		ID:   in.GetId(),
		Name: in.GetName(),
		Role: model.UserRole(in.Role),
	}
}

// UserInfoToProto adapts the local model to the proto model.
func UserInfoToProto(in *model.UserInfo) *desc.UserInfo {
	out := &desc.UserInfo{
		Id:    in.ID,
		Name:  in.Name,
		Email: in.Email,
		Role:  desc.Role(in.Role),
	}

	out.Timestamp = &desc.UserInfo_CreatedAt{CreatedAt: timestamppb.New(in.CreatedAt)}

	if in.UpdatedAt != nil {
		out.Timestamp = &desc.UserInfo_UpdatedAt{UpdatedAt: timestamppb.New(*in.UpdatedAt)}
	}

	return out
}
