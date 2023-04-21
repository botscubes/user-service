package server

import (
	"net/http"

	"github.com/botscubes/user-service/internal/errors"
	"github.com/botscubes/user-service/internal/user"
	"github.com/botscubes/user-service/pkg/service_error"
	"github.com/labstack/echo/v4"
)

// Handlers for server. Handlers are implemented using a closure.
func (s *Server) bindHanlers() {

	s.echo.POST("/signin", func(c echo.Context) error {
		var u *user.User = new(user.User)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, errors.ErrBadRequest)
		}

		_, service_err := user.NewUser(u.Login, u.Password)
		if service_err != errors.NoError {
			return c.JSON(http.StatusOK, service_err)
		}

		return c.JSON(http.StatusOK, service_err)
	})

	s.echo.POST("/signup", func(c echo.Context) error {

		return c.JSON(http.StatusOK, service_error.ServiceError{Code: 1, Message: "test"})
	})

	s.echo.POST("/signout", func(c echo.Context) error {

		return c.JSON(http.StatusOK, service_error.ServiceError{Code: 1, Message: "test"})
	})
}
