package user

import (
	"github.com/ArtEmerged/o_auth-server/internal/service"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

var _ desc.UserV1Server = (*Implementation)(nil)

// Implementation implements user gRPC interface.
type Implementation struct {
	desc.UnimplementedUserV1Server

	userService service.UserService
}

// NewImplementation registers the user service on the gRPC server.
// s - pointer to the gRPC server
// service - the user service interface to be registered
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{userService: userService}
}
