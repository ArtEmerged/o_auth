package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

func TestCreateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request model.CreateUserRequest
		wantErr error
	}{
		{
			name: "Valid request",
			request: model.CreateUserRequest{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Password:        "password123",
				PasswordConfirm: "password123",
			},
			wantErr: nil,
		},
		{
			name: "Missing name",
			request: model.CreateUserRequest{
				Email:           "john.doe@example.com",
				Password:        "password123",
				PasswordConfirm: "password123",
			},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field name is required"),
		},
		{
			name: "Missing email",
			request: model.CreateUserRequest{
				Name:            "John Doe",
				Password:        "password123",
				PasswordConfirm: "password123",
			},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field email is required"),
		},
		{
			name: "Missing password",
			request: model.CreateUserRequest{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				PasswordConfirm: "password123",
			},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field password is required, password and password_confirm don't match"),
		},
		{
			name: "Missing password_confirm",
			request: model.CreateUserRequest{
				Name:     "John Doe",
				Email:    "john.doe@example.com",
				Password: "password123",
			},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field password_confirm is required, password and password_confirm don't match"),
		},
		{
			name: "Missing password and password_confirm",
			request: model.CreateUserRequest{
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field password is required, field password_confirm is required"),
		},
		{
			name: "Password mismatch",
			request: model.CreateUserRequest{
				Name:            "John Doe",
				Email:           "john.doe@example.com",
				Password:        "password123",
				PasswordConfirm: "differentPassword",
			},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "password and password_confirm don't match"),
		},
		{
			name: "Multiple errors",
			request: model.CreateUserRequest{
				Email:           "",
				Password:        "",
				PasswordConfirm: "differentPassword",
			},
			wantErr: fmt.Errorf("%w: %s", model.ErrInvalidArgument, "field name is required, field email is required, field password is required, password and password_confirm don't match"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			require.Equal(t, tt.wantErr, err)
		})
	}
}
