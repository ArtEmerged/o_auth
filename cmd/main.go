package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ArtEmerged/o_auth-server/config"
	grpc_user "github.com/ArtEmerged/o_auth-server/internal/grpc/user"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
	service "github.com/ArtEmerged/o_auth-server/internal/sevice"
	"github.com/ArtEmerged/o_auth-server/pkg/database"
)

func main() {
	cfg := config.New()
	cfg.Init("")

	db := database.NewPostgres(cfg.GetDbDNS())
	repo := repository.New(db)
	service := service.New(repo, cfg.SaltPassword)

	l, err := net.Listen("tcp", cfg.GetServerAddress())
	if err != nil {
		panic(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	grpc_user.Register(s, service)

	go func() {
		log.Printf("listen on :%s", cfg.ServerPort)
		if err = s.Serve(l); err != nil {
			log.Println(err)
		}
	}()

	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-quit.Done()

	s.GracefulStop()
	log.Println("GRPC server has been shut down")
	db.Close()
	log.Println("Database closed")
}
