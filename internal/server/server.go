package server

import (
	"context"
	"log"

	"github.com/botscubes/user-service/internal/config"
	"github.com/botscubes/user-service/internal/db/pgsql"
	"github.com/botscubes/user-service/internal/db/redis"
	"github.com/botscubes/user-service/internal/usermodel"
	"github.com/botscubes/user-service/pkg/token_storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

// Server for user-service.
type Server struct {
	echo         *echo.Echo
	conf         *config.Config
	tokenStorage token_storage.TokenStorage
	pgpool       *pgxpool.Pool
	userModel    *usermodel.UserModel
}

// Create user-service server.
func NewServer() *Server {
	// TODO: log errors
	var err error
	s := new(Server)

	s.conf, err = config.GetConfig("configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	redis := redis.GetClient(&s.conf.Redis)

	s.tokenStorage = token_storage.NewRedisTokenStorage(ctx, redis)
	s.pgpool, err = pgsql.NewPool(ctx, &s.conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	s.userModel = usermodel.New(s.pgpool)

	s.echo = echo.New()
	s.bindHandlers()

	return s

}

// Run user-service server.
func (s *Server) Run() {
	defer s.CloseConnectons()

	s.echo.Logger.Fatal(s.echo.Start(":1323"))

}

// Close all database connections.
func (s *Server) CloseConnectons() {
	// TODO: if the close fails, register in log and throw an error.

	// Note: if Radis is used in several modules, then you
	// need to close it once through its interface, not token storage.
	s.tokenStorage.Close()

	s.pgpool.Close()

}
