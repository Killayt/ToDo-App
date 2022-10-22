package main

import (
	"log"

	todoapp "github.com/Killayt/ToDo-App"
	"github.com/Killayt/ToDo-App/pkg/handler"
	"github.com/Killayt/ToDo-App/pkg/repository"
	"github.com/Killayt/ToDo-App/pkg/service"
)

func main() {
	repos := repository.NewRepository()      //
	services := service.NewService(repos)    // making dependencies to restrictions
	handlers := handler.NewHandler(services) //

	srv := new(todoapp.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to run http server: %s", err.Error())
	}
}
