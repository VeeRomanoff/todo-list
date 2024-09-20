package main

import (
	"github.com/VeeRomanoff/todo-app"
	"github.com/VeeRomanoff/todo-app/internal/handler"
	"github.com/VeeRomanoff/todo-app/internal/repository"
	"github.com/VeeRomanoff/todo-app/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	// Инициализация конфиг файла
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	// Загрузка хуйни из .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
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
		logrus.Fatalf("error initializing db: %s ", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := todo_app.NewServer()
	if err := srv.RunServer(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while starting server: %s", err.Error())
	}
}
