package main

import (
	todo "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/handler"
	"github.com/antonchaban/file-hub-go/pkg/repository"
	"github.com/antonchaban/file-hub-go/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error occurred while reading config file: ", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running http server: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
