package main

import (
	"auth"
	"auth/pkg/handler"
	log "github.com/sirupsen/logrus"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(auth.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server %s", err.Error())
	}
}
