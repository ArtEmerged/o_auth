package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ArtEmerged/o_auth-server/internal/model"
)

func TestUpdateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request model.UpdateUserRequest
		wantErr error
	}{
		{
			name: "Valid request",
			request: model.UpdateUserRequest{
				Name: "John Doe",
				Role: model.RoleAdmin,
			},
			wantErr: nil,
		},
		{
			name: "Missing name",
			request: model.UpdateUserRequest{
				Role: model.RoleUser,
			},
			wantErr: nil,
		},
		{
			name: "Missing role",
			request: model.UpdateUserRequest{
				Name: "John Doe",
			},
			wantErr: nil,
		},
		{
			name:    "Missing name and role",
			request: model.UpdateUserRequest{},
			wantErr: model.ErrWithoutChanges,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			require.Equal(t, tt.wantErr, err)
		})
	}
}
