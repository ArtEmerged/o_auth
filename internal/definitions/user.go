package definitions

import (
	"time"

	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type userRole int32

const (
	ROLE_UNKNOWN userRole = iota
	ROLE_USER
	ROLE_ADMIN
)

type UpdateUserRequest struct {
	ID   int64
	Name *string
	Role userRole
}

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
	Role      userRole
}

type CreateUserRequest struct {
	Name         string
	Email        string
	Password     string
	PasswordHash string
	Role         userRole
}

func AdaptedCreateUserRequestToLocal(in *desc.CreateUserRequest) *CreateUserRequest {
	return &CreateUserRequest{
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
		Role:     userRole(in.Role),
	}
}

func AdaptedUpdateUserRequestToLocal(in *desc.UpdateUserRequest) *UpdateUserRequest {
	return &UpdateUserRequest{
		ID:   in.GetId(),
		Name: in.Name,
		Role: userRole(in.Role),
	}
}

func (c *User) ToProto() *desc.UserInfo {
	out := &desc.UserInfo{
		Id:    c.ID,
		Name:  c.Name,
		Email: c.Email,
		Role:  desc.Role(c.Role),
	}

	out.Timestamp = &desc.UserInfo_CreatedAt{CreatedAt: timestamppb.New(c.CreatedAt)}

	if c.UpdatedAt != nil {
		out.Timestamp = &desc.UserInfo_UpdatedAt{UpdatedAt: timestamppb.New(*c.UpdatedAt)}
	}

	return out
}
