package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"time"

	"strings"

	"github.com/botscubes/user-service/internal/config"
	"github.com/botscubes/user-service/internal/db/pgsql"
	"github.com/botscubes/user-service/internal/db/redis"
	"github.com/botscubes/user-service/internal/usermodel"
	"github.com/botscubes/user-service/pkg/jwt"
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

func JWT(
	JWTKey string,
	tokenStorage token_storage.TokenStorage,
	logger echo.Logger,
) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tmp := c.Request().Header.Get("Authorization")
			token := strings.TrimSpace(strings.TrimPrefix(tmp, "Bearer"))
			exists, err := tokenStorage.CheckToken(context.Background(), token)
			if err != nil {
				logger.Error(err)
				return c.NoContent(http.StatusInternalServerError)
			}
			if !exists {
				return c.NoContent(http.StatusUnauthorized)
			}
			id, err := jwt.GetIdFromToken(token, JWTKey)
			c.Set("user_id", 0)
			c.Set("token", "")
			if err != nil {
				logger.Error(err)
				return c.NoContent(http.StatusInternalServerError)

			} else {
				c.Set("user_id", id)
				c.Set("token", token)
			}
			return next(c)
		}
	}
}

// Create user-service server.
func NewServer() *Server {
	var err error
	s := new(Server)

	s.echo = echo.New()
	s.conf, err = config.GetConfig()
	if err != nil {
		s.echo.Logger.Fatal(err)
	}

	redis := redis.GetClient(&s.conf.Redis)

	s.tokenStorage = token_storage.NewRedisTokenStorage(redis)
	s.pgpool, err = pgsql.NewPool(context.Background(), &s.conf.DB)
	if err != nil {
		s.echo.Logger.Fatal(err)
	}

	s.userModel = usermodel.New(context.Background(), s.pgpool)

	s.bindHandlers()

	//s.echo.Use(JWT(s.conf.Server.JWTKey, s.tokenStorage))
	return s

}

// Run user-service server.
func (s *Server) Run() {
	defer s.CloseConnectons()

	go func() {
		if err := s.echo.Start(":1323"); err != nil && err != http.ErrServerClosed {
			s.echo.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal(err)
	}

	println("Done")
}

// Close all database connections.
func (s *Server) CloseConnectons() {
	// TODO: if the close fails, register in log and throw an error.

	// Note: if Radis is used in several modules, then you
	// need to close it once through its interface, not token storage.
	s.tokenStorage.Close()

	s.pgpool.Close()

}
