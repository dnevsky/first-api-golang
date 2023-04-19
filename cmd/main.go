package main

import (
	"log"

	"github.com/dnevsky/firstapi"
	"github.com/dnevsky/firstapi/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(firstapi.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while start http server: %s", err.Error())
	}

}
