package grpc_user

import (
	desc "github.com/ArtEmerged/o_auth-server/pkg/auth_v1"
	"google.golang.org/grpc"
)

type userServer struct {
	desc.UnimplementedUserV1Server
}

func Register(s *grpc.Server) {
	desc.RegisterUserV1Server(s, &userServer{})
}
