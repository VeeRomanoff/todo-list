package main

import (
	"github.com/VeeRomanoff/todo-app"
	"github.com/VeeRomanoff/todo-app/pkg/handler"
	"github.com/VeeRomanoff/todo-app/pkg/repository"
	"github.com/VeeRomanoff/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error initializing config", err)
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := todo_app.NewServer()
	if err := srv.RunServer(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error while starting server", err)
	}
}
