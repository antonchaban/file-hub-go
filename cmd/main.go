package main

import (
	"context"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/handler"
	"github.com/antonchaban/file-hub-go/pkg/repository"
	"github.com/antonchaban/file-hub-go/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title File-Hub API
// @version 1.0
// @description API Server for Files Managing Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatal("error occurred while reading config file: ", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error occurred while loading env variables: ", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatal("error occurred while connecting to db: ", err.Error())
	}

	handlers := handler.NewHandler(service.NewAuthService(repository.NewAuthPostgres(db)),
		service.NewFolderService(repository.NewFolderPostgres(db)),
		service.NewFileService(repository.NewFilePostgres(db), repository.NewFolderPostgres(db)))

	srv := new(fhub.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatal("error occurred while running http server: ", err.Error())
		}
	}()

	logrus.Print("file-hub-go started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("file-hub-go shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
