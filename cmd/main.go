package main

import (
	"log"

	"github.com/dnevsky/firstapi"
	"github.com/dnevsky/firstapi/pkg/handler"
	"github.com/dnevsky/firstapi/pkg/repository"
	"github.com/dnevsky/firstapi/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(firstapi.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while start http server: %s", err.Error())
	}

}
