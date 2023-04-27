package server

import (
	"context"
	"log"
	"net/http"

	"github.com/botscubes/user-service/internal/errors"
	"github.com/botscubes/user-service/internal/user"
	"github.com/botscubes/user-service/pkg/jwt"
	"github.com/botscubes/user-service/pkg/password_hash"
	"github.com/botscubes/user-service/pkg/service_error"
	"github.com/labstack/echo/v4"
)

type ResponseToken struct {
	Token string                      `json:"token"`
	Error *service_error.ServiceError `json:"error"`
}

// Handlers for server. Handlers are implemented using a closure.
func (s *Server) bindHandlers() {
	// TODO: log errors

	s.echo.POST("/signup", func(c echo.Context) error {
		var u *user.User = new(user.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, errors.ErrBadRequest)
		}

		u, service_err := user.NewUser(u.Login, u.Password)
		if service_err != errors.NoError {
			return c.JSON(http.StatusOK, service_err)
		}

		if exists, err := s.userModel.LoginExists(context.Background(), u.Login); err != nil {
			// TODO: log the error.
			log.Fatal(err) // replace
			return c.JSON(http.StatusInternalServerError, errors.ErrInternalServerError)
		} else if exists {
			return c.JSON(http.StatusOK, errors.ErrLoginExists)
		}

		var err error = nil
		u.Password, err = password_hash.GetPasswordHash(u.Password, s.conf.Server.Salt)
		if err != nil {
			// TODO: log the error.
			log.Fatal(err) // replace
			return c.JSON(http.StatusInternalServerError, errors.ErrInternalServerError)
		}

		err = s.userModel.SaveUser(context.Background(), u)
		if err != nil {
			// TODO: log the error.
			log.Fatal(err) // replace
			return c.JSON(http.StatusInternalServerError, errors.ErrInternalServerError)
		}

		return c.JSON(http.StatusOK, errors.NoError)
	})

	s.echo.POST("/signin", func(c echo.Context) error {
		var u *user.User = new(user.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, errors.ErrBadRequest)
		}

		u, service_err := user.NewUser(u.Login, u.Password)
		if service_err != errors.NoError {
			return c.JSON(http.StatusOK, service_err)
		}

		id, password, err := s.userModel.GetIdAndPasswordByLogin(context.Background(), u.Login)
		if err != nil {
			// TODO: log the error.
			log.Fatal(err) // replace
			return c.JSON(http.StatusInternalServerError, errors.ErrInternalServerError)
		}
		if id == 0 {
			return c.JSON(http.StatusInternalServerError, errors.ErrLoginNotExists)
		}

		if !password_hash.CheckPasswordHash(u.Password, password, s.conf.Server.Salt) {
			return c.JSON(http.StatusOK, errors.ErrPasswordIsNotEqual)
		}

		token, err := jwt.GenerateToken(id, s.conf.Server.JWTKey)
		if err != nil {
			// TODO: log the error.
			log.Fatal(err) // replace
			return c.JSON(http.StatusInternalServerError, errors.ErrInternalServerError)
		}

		err = s.tokenStorage.SaveToken(context.Background(), token, s.conf.Server.TokenLifetime)
		if err != nil {
			// TODO: log the error.
			log.Fatal(err) // replace
			return c.JSON(http.StatusInternalServerError, errors.ErrInternalServerError)
		}

		return c.JSON(http.StatusOK, ResponseToken{token, errors.NoError})
	})

	s.echo.POST("/signout", func(c echo.Context) error {
		token := c.Get("token").(string)
		if token == "" {
			return c.JSON(http.StatusUnauthorized, errors.ErrUnauthorized)
		}
		err := s.tokenStorage.DeleteToken(context.Background(), token)
		if err != nil {
			// TODO: log the error.
			log.Fatal(err) // replace
			return c.JSON(http.StatusInternalServerError, errors.ErrInternalServerError)
		}

		return c.JSON(http.StatusUnauthorized, errors.NoError)

	}, JWT(s.conf.Server.JWTKey, s.tokenStorage))
}
