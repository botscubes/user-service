package server

import (
	"log"

	"github.com/botscubes/user-service/internal/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	conf *config.Config
}

func NewServer() *Server {
	var err error
	s := new(Server)

	s.conf, err = config.GetConfig("configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	s.echo = echo.New()

	// s.echo.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	s.bindHanlers()

	return s

}

func (s *Server) Run() {
	s.echo.Logger.Fatal(s.echo.Start(":1323"))
}
