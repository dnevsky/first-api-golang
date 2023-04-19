package main

import (
	"log"

	"github.com/dnevsky/firstapi"
)

func main() {
	srv := new(firstapi.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error while start http server: %s", err.Error())
	}

}
