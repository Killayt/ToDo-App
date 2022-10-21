package main

import (
	"log"

	todoapp "github.com/Killayt/ToDo-App"
	"github.com/Killayt/ToDo-App/pkg/handler"
)

func main() {

	handler := new(handler.Handler)
	srv := new(todoapp.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Failed to run http server: %s", err.Error())
	}
}
