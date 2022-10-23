package main

import (
	"log"

	todoapp "github.com/Killayt/ToDo-App"
	"github.com/Killayt/ToDo-App/pkg/handler"
	"github.com/Killayt/ToDo-App/pkg/repos"
	"github.com/Killayt/ToDo-App/pkg/service"
)

func main() {

	repos := repos.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("Failed to run http server: %s", err.Error())
	}
}
