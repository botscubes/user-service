package server

import (
	"context"
	"net/http"
	"time"

	"github.com/botscubes/user-service/internal/user"
	"github.com/botscubes/user-service/pkg/jwt"
	"github.com/botscubes/user-service/pkg/password_hash"
	//"github.com/botscubes/user-service/pkg/service_error"
	"github.com/labstack/echo/v4"
)

type ResponseToken struct {
	Token string `json:"token"`
	//Error *service_error.ServiceError `json:"error"`
}

// Handlers for server. Handlers are implemented using a closure.
func (s *Server) bindHandlers() {

	s.echo.POST("/api/users/signup", func(c echo.Context) error {
		var u *user.User = new(user.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		u, service_err := user.NewUser(u.Login, u.Password)
		if service_err != nil {
			return c.JSON(http.StatusUnprocessableEntity, service_err)
		}

		if exists, err := s.userModel.LoginExists(context.Background(), u.Login); err != nil {

			s.echo.Logger.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		} else if exists {
			return c.JSON(http.StatusUnprocessableEntity, user.ErrLoginExists)
		}

		var err error = nil
		u.Password, err = password_hash.GetPasswordHash(u.Password, s.conf.Server.Salt)
		if err != nil {
			s.echo.Logger.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		err = s.userModel.SaveUser(context.Background(), u)
		if err != nil {
			s.echo.Logger.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusCreated)
	})

	s.echo.POST("/api/users/signin", func(c echo.Context) error {
		var u *user.User = new(user.User)
		if err := c.Bind(u); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		u, service_err := user.NewUser(u.Login, u.Password)
		if service_err != nil {
			return c.JSON(http.StatusUnprocessableEntity, service_err)
		}

		id, password, err := s.userModel.GetIdAndPasswordByLogin(context.Background(), u.Login)
		if err != nil {

			s.echo.Logger.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
		if id == 0 {
			return c.JSON(http.StatusUnprocessableEntity, user.ErrLoginNotExists)
		}

		if !password_hash.CheckPasswordHash(u.Password, password, s.conf.Server.Salt) {
			return c.JSON(http.StatusUnprocessableEntity, user.ErrPasswordIsNotEqual)
		}
		claims := jwt.NewUserClaims(
			id,
			time.Duration(s.conf.Server.TokenLifeTime)*time.Second,
		)
		token, err := jwt.GenerateToken(
			claims,
			s.conf.Server.JWTKey,
		)
		if err != nil {
			s.echo.Logger.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		err = s.tokenStorage.SaveToken(
			context.Background(),
			token,
			claims.GetLifeDuration(),
		)
		if err != nil {

			s.echo.Logger.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusCreated, ResponseToken{token})
	})

	s.echo.DELETE("/api/users/signout", func(c echo.Context) error {
		token := c.Get("token").(string)
		if token == "" {
			return c.NoContent(http.StatusNoContent)
		}
		err := s.tokenStorage.DeleteToken(context.Background(), token)
		if err != nil {
			s.echo.Logger.Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusNoContent)

	}, JWT(s.conf.Server.JWTKey, s.tokenStorage, s.echo.Logger))
}
