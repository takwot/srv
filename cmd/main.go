package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	todo "github.com/takwot/srv"
	"github.com/takwot/srv/pkg/handler"
	"github.com/takwot/srv/pkg/repository"
	"github.com/takwot/srv/pkg/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while init config", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "",
		Username: "",
		SSLMode:  "",
		Password: "",
		DBName:   "",
		Port:     "",
	})

	if err != nil {
		logrus.Fatalf("Error while initing Database")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal("error occurred while running http server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
