package server

import (
	"context"
	"log"

	"github.com/botscubes/user-service/internal/config"
	"github.com/botscubes/user-service/internal/db/redis"
	"github.com/botscubes/user-service/pkg/token_storage"
	"github.com/labstack/echo/v4"
)

// Server for user-service.
type Server struct {
	echo         *echo.Echo
	conf         *config.Config
	tokenStorage token_storage.TokenStorage
}

// Create user-service server.
func NewServer() *Server {
	var err error
	s := new(Server)

	s.conf, err = config.GetConfig("configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	redis := redis.GetClient(&s.conf.Redis)
	ctx := context.Background()

	s.tokenStorage = token_storage.NewRedisTokenStorage(redis, &ctx)

	s.echo = echo.New()
	s.bindHanlers()

	return s

}

// Run user-service server.
func (s *Server) Run() {
	s.echo.Logger.Fatal(s.echo.Start(":1323"))

}
