package main

import (
	"github.com/VeeRomanoff/todo-app"
	"github.com/VeeRomanoff/todo-app/pkg/handler"
	"github.com/VeeRomanoff/todo-app/pkg/repository"
	"github.com/VeeRomanoff/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func main() {

	// Инициализация конфиг файла
	if err := initConfig(); err != nil {
		log.Fatal("error initializing config", err)
	}

	// Загрузка хуйни из .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file", err)
	}

	// Создание подключения к базе данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal("error initializing db ", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := todo_app.NewServer()
	if err := srv.RunServer(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error while starting server", err)
	}
}
