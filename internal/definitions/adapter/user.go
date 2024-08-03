package adapter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

// CreateUserRequestToLocal adapts the CreateUserRequest proto to the local definition.
func CreateUserRequestToLocal(in *desc.CreateUserRequest) *def.CreateUserRequest {
	return &def.CreateUserRequest{
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		Role:     def.UserRole(in.Role),
	}
}

// UpdateUserRequestToLocal adapts the UpdateUserRequest proto to the local definition.
func UpdateUserRequestToLocal(in *desc.UpdateUserRequest) *def.UpdateUserRequest {
	return &def.UpdateUserRequest{
		ID:   in.GetId(),
		Name: in.Name,
		Role: def.UserRole(in.Role),
	}
}

// UserInfoToProto adapts the local definition to the proto definition.
func UserInfoToProto(in *def.UserInfo) *desc.UserInfo {
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
