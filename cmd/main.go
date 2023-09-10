package main

import (
	"auth"
	log "github.com/sirupsen/logrus"
)

func main() {
	srv := new(auth.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error running http server %s", err.Error())
	}
}
