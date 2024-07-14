package grpc_user

import (
	"google.golang.org/grpc"

	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
)

type userServer struct {
	desc.UnimplementedUserV1Server
}

func Register(s *grpc.Server) {
	desc.RegisterUserV1Server(s, &userServer{})
}
