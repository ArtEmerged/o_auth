package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_user "github.com/ArtEmerged/o_auth-server/internal/grpc/user"
)

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	grpc_user.Register(s)

	fmt.Println("listen on :50051")

	if err = s.Serve(l); err != nil {
		log.Println(err)
	}
}
