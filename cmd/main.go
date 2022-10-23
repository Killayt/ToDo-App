package main

import (
	"log"

	todoapp "github.com/Killayt/ToDo-App"
	"github.com/Killayt/ToDo-App/pkg/handler"
	"github.com/Killayt/ToDo-App/pkg/repos"
	"github.com/Killayt/ToDo-App/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("[ERROR] with config file: %s", err.Error())
	}

	repos := repos.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todoapp.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Failed to run http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")

	return viper.ReadInConfig()
}
