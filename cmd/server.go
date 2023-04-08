package main

import "github.com/botscubes/user-service/internal/server"

func main() {
	s := server.NewServer()

	s.Run()
}
