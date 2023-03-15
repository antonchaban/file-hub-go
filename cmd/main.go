package main

import (
	todo "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("error occurred while running http server: ", err.Error())
	}
}
