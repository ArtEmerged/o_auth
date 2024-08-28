package app

import (
	"context"
	"log"

	userGRPC "github.com/ArtEmerged/o_auth-server/internal/api/grpc/user"
	"github.com/ArtEmerged/o_auth-server/internal/client/db"
	"github.com/ArtEmerged/o_auth-server/internal/client/db/pg"
	"github.com/ArtEmerged/o_auth-server/internal/client/db/transaction"
	"github.com/ArtEmerged/o_auth-server/internal/closer"
	"github.com/ArtEmerged/o_auth-server/internal/config"
	"github.com/ArtEmerged/o_auth-server/internal/repository"
	userRepo "github.com/ArtEmerged/o_auth-server/internal/repository/user"
	"github.com/ArtEmerged/o_auth-server/internal/service"
	userServ "github.com/ArtEmerged/o_auth-server/internal/service/user"
)

type serviceProvider struct {
	globalConfig *config.Config

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepo

	userService service.UserService

	userImpl *userGRPC.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GlobalConfig initializes and returns the global config.
func (s *serviceProvider) GlobalConfig() *config.Config {
	if s.globalConfig == nil {
		s.globalConfig = config.New()

		err := s.globalConfig.Init("")
		if err != nil {
			log.Fatalf("failed init config: %v", err)
		}
	}

	return s.globalConfig
}

// DBClient initializes and returns the database client.
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GlobalConfig().GetDbDNS())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %v", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager initializes and returns the transaction manager.
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {

	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// UserRepository returns the user repository.
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepo {
	if s.userRepository == nil {
		s.userRepository = userRepo.New(s.DBClient(ctx))
	}

	return s.userRepository
}

// UserService returns the user service.
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userServ.New(s.UserRepository(ctx), s.TxManager(ctx), s.GlobalConfig().GetSalt())
	}

	return s.userService
}

// UserImplementation returns the user gRPC implementation.
func (s *serviceProvider) UserImplementation(ctx context.Context) *userGRPC.Implementation {
	if s.userImpl == nil {
		s.userImpl = userGRPC.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
