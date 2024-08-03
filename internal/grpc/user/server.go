package grpc_user

import (
	"google.golang.org/grpc"

	def "github.com/ArtEmerged/o_auth-server/internal/definitions"
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

var _ desc.UserV1Server = (*userServer)(nil)

type userServer struct {
	desc.UnimplementedUserV1Server
	service def.UserService
}

// Register registers the user service on the gRPC server.
// s - pointer to the gRPC server
// service - the user service interface to be registered
func Register(s *grpc.Server, service def.UserService) {
	desc.RegisterUserV1Server(s, &userServer{service: service})
}
